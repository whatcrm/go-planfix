package models

import "time"

type DataTag struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type,omitempty"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

type DataTagEntry struct {
	Key        string    `json:"key"`
	Value      string    `json:"value"`
	DataTagID  int       `json:"dataTagId"`
	EntityID   int       `json:"entityId"`
	EntityType string    `json:"entityType"`
	CommentID  int       `json:"commentId,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

type DataTagListRequest struct {
	Offset   int    `json:"offset,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Fields   string `json:"fields,omitempty"`
}

type DataTagEntryListRequest struct {
	Offset    int                    `json:"offset,omitempty"`
	PageSize  int                    `json:"pageSize,omitempty"`
	Fields    string                 `json:"fields,omitempty"`
	TaskId    int                    `json:"taskId,omitempty"`
	ContactId string                 `json:"contactId,omitempty"`
	Filters   []ComplexDataTagFilter `json:"filters,omitempty"`
}

type DataTagEntryRequest struct {
	DataTag *BaseEntity `json:"dataTag"`
	Items   []struct {
		CustomFieldData []CustomFieldValueRequest `json:"customFieldData"`
	} `json:"items"`
}

type ComplexDataTagFilter struct {
	Type     int         `json:"type"`
	Field    int         `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type DataTagResponse struct {
	ID int `json:"id"`
	//Name  string         `json:"name"`
	//Group *GroupResponse `json:"group,omitempty"`
	Items     []CustomField `json:"items,omitempty"`
	CommentID int           `json:"commentId,omitempty"`
}

type DataTagGetResponse struct {
	Result  string          `json:"result"`
	DataTag DataTagResponse `json:"dataTag"`
}

type DataTagListResponse struct {
	Result   string            `json:"result"`
	DataTags []DataTagResponse `json:"dataTags"`
}

type DataTagEntryResponse struct {
	DataTag         DataTagResponse            `json:"dataTag"`
	Key             int                        `json:"key"`
	CommentId       int                        `json:"commentId,omitempty"`
	Task            *BaseEntity                `json:"task,omitempty"`
	Contact         *BaseEntity                `json:"contact,omitempty"`
	CustomFieldData []CustomFieldValueResponse `json:"customFieldData,omitempty"`
}

type DataTagEntryGetResponse struct {
	Result string               `json:"result"`
	Entry  DataTagEntryResponse `json:"entry"`
}

type DataTagEntryListResponse struct {
	Result         string                 `json:"result"`
	DataTagEntries []DataTagEntryResponse `json:"dataTagEntries"`
}

type DataTagEntryUpdateRequest struct {
	CustomFieldData []CustomFieldValueRequest `json:"customFieldData"`
}

type DataTagEntryCommentResponse struct {
	Result string `json:"result"`
	Keys   []int  `json:"keys"`
}
