package models

type ToolsCategory string

const (
	ToolsPrimary    ToolsCategory = "primary"
	ToolsSecondary  ToolsCategory = "secondary"
	ToolsAdditional ToolsCategory = "additional"
	ToolsSpecial    ToolsCategory = "special"
)

type StockToolsReferenceType string

const (
	RefInitialStock StockToolsReferenceType = "initial_stock"
	RefReplacement  StockToolsReferenceType = "replacement"
	RefProcurement  StockToolsReferenceType = "procurement"
	RefAdjustment   StockToolsReferenceType = "adjustment"
)

type ToolsUsagePeriodUnit string

const (
	ToolsUsagePeriodYearly  ToolsUsagePeriodUnit = "Y"
	ToolsUsagePeriodMonthly ToolsUsagePeriodUnit = "M"
	ToolsUsagePeriodWeekly  ToolsUsagePeriodUnit = "W"
	ToolsUsagePeriodDialy   ToolsUsagePeriodUnit = "D"
)

type ModifyStockToolsType string

const (
	ModifyStockToolsTypeIncoming ModifyStockToolsType = "IN"
	ModifyStockToolsTypeOutgoing ModifyStockToolsType = "OUT"
)
