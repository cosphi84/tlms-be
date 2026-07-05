package services

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"
	"tlms/internal/auth"
	"tlms/internal/dto"
	"tlms/internal/models"
	"tlms/internal/repositories"
	"tlms/internal/storage"
	"tlms/internal/validators"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrFileMetaNotFound = errors.New("file: metadata not found")
)

type FileService interface {
	Upload(ctx context.Context, folder string, fh *multipart.FileHeader) (*dto.FileResponse, error)
	GetMetadata(uuid string) (*dto.FileResponse, error)
	OpenForDownload(ctx context.Context, uuid string) (file interface{ Read([]byte) (int, error) }, meta *models.UploadFile, err error)
	Delete(ctx context.Context, uuid string) error

	// ReplaceWithTx dipanggil DI DALAM transaction milik caller module.
	// Alur (sesuai requirement): upload file baru → caller update FK →
	// caller commit → baru file lama dihapus fisik & soft-deleted.
	// Karena itu method ini TIDAK menghapus file lama sendiri — ia
	// mengembalikan info file lama supaged caller memutuskan kapan
	// FinalizeReplace dipanggil (setelah commit sukses).
	ReplaceWithTx(ctx context.Context, tx *gorm.DB, folder string, fh *multipart.FileHeader, oldFileUUID string) (newFile *dto.FileResponse, err error)

	// FinalizeReplace membersihkan file lama SETELAH caller berhasil commit.
	// Dipanggil terpisah, di luar transaction DB (karena ini I/O disk,
	// bukan operasi DB) — mencegah file fisik lama terhapus duluan
	// padahal transaction caller ternyata rollback.
	FinalizeReplace(ctx context.Context, oldFileUUID string) error
}

type fileService struct {
	repo      repositories.FileRepository
	storage   storage.Storage
	validator validators.FileValidator
}

func NewFileService(repo repositories.FileRepository, strg storage.Storage, validator validators.FileValidator) FileService {
	return &fileService{repo, strg, validator}
}

// buildPath menghasilkan relative path fisik: {folder}/{yyyy}/{mm}/{uuid}{ext}
// Business rule struktur folder tanggal ada di sini, BUKAN di Storage layer.
func buildPath(folder, fileUUID, ext string) string {
	now := time.Now()
	return fmt.Sprintf("%s/%04d/%02d/%s%s", folder, now.Year(), int(now.Month()), fileUUID, ext)
}

func toFileResponse(m *models.UploadFile) *dto.FileResponse {
	return &dto.FileResponse{
		ID:           m.ID,
		UUID:         m.UUID,
		OriginalName: m.OriginalName,
		MimeType:     m.MimeType,
		Extension:    m.Extension,
		Size:         m.Size,
		Checksum:     m.Checksum,
		URL:          fmt.Sprintf("/files/%s", m.UUID),
		IsArchived:   m.IsArchived,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func (s *fileService) Upload(ctx context.Context, folder string, fh *multipart.FileHeader) (*dto.FileResponse, error) {
	usr, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, errors.New("invalid claims")
	}

	ext, err := s.validator.ValidateHeader(fh)
	if err != nil {
		return nil, err
	}

	f, err := fh.Open()
	if err != nil {
		return nil, fmt.Errorf("file: failed to open uploaded file: %w", err)
	}
	defer f.Close()

	if err := s.validator.ValidateContent(f, ext); err != nil {
		return nil, err
	}

	checksum, size, err := s.validator.ComputeChecksum(f)
	if err != nil {
		return nil, err
	}
	if _, err := f.Seek(0, 0); err != nil {
		return nil, fmt.Errorf("file: failed to reset file position: %w", err)
	}

	fileUUID := uuid.NewString()
	relPath := buildPath(folder, fileUUID, ext)

	if _, err := s.storage.Save(ctx, relPath, f); err != nil {
		return nil, fmt.Errorf("file: failed to persist file: %w", err)
	}

	record := &models.UploadFile{
		UUID:         fileUUID,
		DiskName:     fileUUID + ext,
		OriginalName: fh.Filename,
		MimeType:     fh.Header.Get("Content-Type"),
		Extension:    ext,
		Size:         size,
		Checksum:     checksum,
		Path:         relPath,
		Storage:      "local",
		CreatedAt:    time.Now(),
		CreatedBy:    &usr.UserID,
	}

	if err := s.repo.Create(record); err != nil {
		// Rollback fisik: metadata gagal tersimpan, file yang sudah
		// ditulis ke disk harus dibersihkan agar tidak orphan.
		_ = s.storage.Delete(ctx, relPath)
		return nil, fmt.Errorf("file: failed to save metadata: %w", err)
	}

	return toFileResponse(record), nil
}

func (s *fileService) GetMetadata(fileUUID string) (*dto.FileResponse, error) {
	record, err := s.repo.FindByUUID(fileUUID)
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, ErrFileMetaNotFound
	}
	return toFileResponse(record), nil
}

func (s *fileService) OpenForDownload(ctx context.Context, fileUUID string) (interface{ Read([]byte) (int, error) }, *models.UploadFile, error) {
	record, err := s.repo.FindByUUID(fileUUID)
	if err != nil {
		return nil, nil, err
	}
	if record == nil {
		return nil, nil, ErrFileMetaNotFound
	}

	path := record.Path
	if record.IsArchived && record.ArchivedPath != nil {
		path = *record.ArchivedPath
	}

	reader, err := s.storage.Open(ctx, path)
	if err != nil {
		return nil, nil, err
	}

	return reader, record, nil
}

func (s *fileService) Delete(ctx context.Context, fileUUID string) error {
	usr, err := auth.GetClaims(ctx)
	if err != nil {
		return errors.New("invalid claims")
	}

	record, err := s.repo.FindByUUID(fileUUID)
	if err != nil {
		return err
	}
	if record == nil {
		return ErrFileMetaNotFound
	}

	// Sesuai requirement: soft delete metadata dulu, penghapusan fisik
	// dilakukan belakangan oleh retention job — bukan di sini.
	return s.repo.SoftDelete(record.ID, &usr.UserID)
}

func (s *fileService) ReplaceWithTx(ctx context.Context, tx *gorm.DB, folder string, fh *multipart.FileHeader, oldFileUUID string) (*dto.FileResponse, error) {
	usr, err := auth.GetClaims(ctx)
	if err != nil {
		return nil, errors.New("invalid claims")
	}

	oldFile, err := s.repo.FindByUUID(oldFileUUID)
	if err != nil {
		return nil, err
	}
	if oldFile == nil {
		return nil, ErrFileMetaNotFound
	}

	ext, err := s.validator.ValidateHeader(fh)
	if err != nil {
		return nil, err
	}

	f, err := fh.Open()
	if err != nil {
		return nil, fmt.Errorf("file: failed to open uploaded file: %w", err)
	}
	defer f.Close()

	if err := s.validator.ValidateContent(f, ext); err != nil {
		return nil, err
	}

	checksum, size, err := s.validator.ComputeChecksum(f)
	if err != nil {
		return nil, err
	}
	if _, err := f.Seek(0, 0); err != nil {
		return nil, fmt.Errorf("file: failed to reset file position: %w", err)
	}

	newUUID := uuid.NewString()
	relPath := buildPath(folder, newUUID, ext)

	if _, err := s.storage.Save(ctx, relPath, f); err != nil {
		return nil, fmt.Errorf("file: failed to persist replacement file: %w", err)
	}

	record := &models.UploadFile{
		UUID:         newUUID,
		DiskName:     newUUID + ext,
		OriginalName: fh.Filename,
		MimeType:     fh.Header.Get("Content-Type"),
		Extension:    ext,
		Size:         size,
		Checksum:     checksum,
		Path:         relPath,
		Storage:      "local",
		CreatedAt:    time.Now(),
		CreatedBy:    &usr.UserID,
	}

	// Insert metadata baru DI DALAM transaction caller — jika caller
	// rollback (mis. update FK gagal), record ini ikut hilang otomatis.
	txRepo := s.repo.WithTx(tx)
	if err := txRepo.Create(record); err != nil {
		_ = s.storage.Delete(ctx, relPath)
		return nil, fmt.Errorf("file: failed to save replacement metadata: %w", err)
	}

	return toFileResponse(record), nil
}

func (s *fileService) FinalizeReplace(ctx context.Context, oldFileUUID string) error {
	oldFile, err := s.repo.FindByUUID(oldFileUUID)
	if err != nil {
		return err
	}
	if oldFile == nil {
		// Sudah tidak ada — anggap sukses (idempotent), mungkin sudah
		// pernah di-finalize sebelumnya (retry scenario).
		return nil
	}

	if err := s.storage.Delete(ctx, oldFile.Path); err != nil {
		return fmt.Errorf("file: failed to delete old physical file: %w", err)
	}

	return s.repo.SoftDelete(oldFile.ID, oldFile.CreatedBy)
}
