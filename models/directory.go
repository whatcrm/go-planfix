package models

type Directory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsActive    bool   `json:"isActive"`
}

type DirectoryGetResponse struct {
	Result    string    `json:"result"`
	Directory Directory `json:"directory"`
}

type DirectoryRequest struct {
	CommonListParams
}

type DirectoryListResponse struct {
	Result      string      `json:"result"`
	Directories []Directory `json:"directories"`
}

type DirectoryEntry struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

type DirectoryEntryGetResponse struct {
	Result string         `json:"result"`
	Entry  DirectoryEntry `json:"entry"`
}

type DirectoryEntryRequest struct {
	Value string `json:"value"`
}

type DirectoryEntryCreateResponse struct {
	Result string `json:"result"`
	ID     int    `json:"id"`
}

type DirectoryEntryListResponse struct {
	Result  string           `json:"result"`
	Entries []DirectoryEntry `json:"entries"`
}

type DirectoryGroupsResponse struct {
	Result string          `json:"result"`
	Groups []GroupResponse `json:"groups"`
}

type DirectoryFiltersResponse struct {
	Result  string    `json:"result"`
	Filters []Filters `json:"filters"`
}

type Filters struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Owner struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"owner"`
}
