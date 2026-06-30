package dto

type CreateOfficeRequest struct {
	ParentID int64  `json:"parent_id" binding:"required"`
	Code     string `json:"code" binding:"required,max=10"`
	Name     string `json:"name" binding:"required,max=100"`
	Type     string `json:"type" binding:"required,max=100"`
}

type UpdateOfficeRequest struct {
	Code string `json:"code" binding:"required,max=10"`
	Name string `json:"name" binding:"required,max=100"`
	Type string `json:"type" binding:"required,max=100"`
}

type OfficeResponse struct {
	ID            int64            `json:"id"`
	ParentID      *int64           `json:"parent_id,omitempty"`
	Code          string           `json:"code"`
	Name          string           `json:"name"`
	Type          string           `json:"type"`
	Depth         int              `json:"depth"`
	ChildrenCount int              `json:"children_count"`
	CreatedAt     string           `json:"created_at"`
	CreatedBy     *int64           `json:"created_by,omitempty"`
	Children      []OfficeResponse `json:"children,omitempty"`
}

type DeleteOfficeRequest struct {
	OfficeID int64 `json:"office_id" binding:"required"`
}

type OfficeOptionResponse struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
}
