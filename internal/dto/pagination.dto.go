package dto

type PaginationRequest struct {
	Page     int    `form:"page"`
	Limit    int    `form:"limit"`
	Search   string `form:"search"`
	SortedBy string `form:"sorted_by"`
	SortDir  string `form:"sort_dir"`
}

type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
}
