package dto

import "time"

type StorageLocationRequest struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	OfficeID int64  `json:"office_id" binding:"required"`
}

type StorageLocationResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	OfficeID  int64     `json:"office_id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StorageLocationChangeRequest struct {
	ID int64 `json:"id"`
}
