package dto

import (
	"github.com/google/uuid"
)

type ModifyStokToolRequest struct {
	Mode   string    `json:"mode" default:"IN"`
	SlocID int64     `json:"sloc_id"`
	ToolID uuid.UUID `json:"tool_id"`
	Qty    int32     `json:"qty"`
}

type FindStockToolByEntity struct {
	StorageLocation int64     `json:"storage_location"`
	ToolID          uuid.UUID `json:"tool_id"`
}
