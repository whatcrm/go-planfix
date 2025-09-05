package models

type ReportResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsActive    bool   `json:"isActive"`
}

type ReportGetResponse struct {
	Result string         `json:"result"`
	Report ReportResponse `json:"report"`
}

type ReportListRequest struct {
	Offset   int    `json:"offset,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Fields   string `json:"fields,omitempty"`
}

type ReportListResponse struct {
	Result  string           `json:"result"`
	Reports []ReportResponse `json:"reports"`
}

type ReportGenerateRequest struct {
	// TODO: Add fields based on swagger schema
}

type ReportGenerateResponse struct {
	Result    string `json:"result"`
	RequestID string `json:"requestId"`
}

type ReportStatusResponse struct {
	Result  string `json:"result"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type ReportSave struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

type ReportSaveListResponse struct {
	Result string       `json:"result"`
	Saves  []ReportSave `json:"saves"`
}

type ReportSaveDataRequest struct {
	Offset   int    `json:"offset,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Fields   string `json:"fields,omitempty"`
}

type ReportSaveDataItem struct {
	// TODO: Add fields based on swagger schema
}

type ReportSaveDataRow struct {
	// TODO: Add fields based on swagger schema
}

type ReportSaveDataResponse struct {
	Result string              `json:"result"`
	Data   []ReportSaveDataRow `json:"data"`
}
