package dto

import "time"

// FileUploadRequest merepresentasikan metadata form-data upload.
// File binary-nya sendiri diambil langsung dari *gin.Context di handler
// (c.Request.FormFile), bukan lewat struct ini — konsisten dengan pola
// binding project untuk file upload multipart.
type FileUploadRequest struct {
	// Folder menentukan sub-direktori logis (contoh: "tools", "avatars",
	// "work-orders"). Dipakai Service untuk membangun path fisik:
	// {folder}/{yyyy}/{mm}/{uuid}{ext}
	Folder string `form:"folder" binding:"required"`
}

// FileResponse adalah representasi publik dari UploadFile.
// Sengaja TIDAK menyertakan Path/DiskName fisik — caller module lain
// hanya perlu tahu UUID dan URL untuk akses file, sesuai requirement
// "module lain hanya menyimpan file_id, URL dibentuk oleh File Service".
type FileResponse struct {
	ID           int64     `json:"id"`
	UUID         string    `json:"uuid"`
	OriginalName string    `json:"original_name"`
	MimeType     string    `json:"mime_type"`
	Extension    string    `json:"extension"`
	Size         int64     `json:"size"`
	Checksum     string    `json:"checksum"`
	URL          string    `json:"url"`
	IsArchived   bool      `json:"is_archived"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// FileReplaceRequest dipakai endpoint replace — menerima UUID file lama
// yang akan digantikan, plus Folder untuk file baru (biasanya sama
// dengan folder file lama, tapi dibuka untuk fleksibilitas).
type FileReplaceRequest struct {
	Folder string `form:"folder" binding:"required"`
}
