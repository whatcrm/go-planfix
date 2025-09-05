package models

import "time"

type Object struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type,omitempty"`
	Status    string    `json:"status,omitempty"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type ObjectListRequest struct {
	Offset   int `json:"offset,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
}

type ObjectResponse struct {
	ID                int                        `json:"id,omitempty"`
	SourceObjectId    string                     `json:"sourceObjectId,omitempty"`
	SourceDataVersion string                     `json:"sourceDataVersion,omitempty"`
	Name              string                     `json:"name,omitempty"`
	Description       string                     `json:"description,omitempty"`
	Status            *BaseEntity                `json:"status,omitempty"`
	Owner             *PersonResponse            `json:"owner,omitempty"`
	Project           *BaseEntity                `json:"project,omitempty"`
	Group             *GroupResponse             `json:"group,omitempty"`
	StartDate         *TimePoint                 `json:"startDate,omitempty"`
	EndDate           *TimePoint                 `json:"endDate,omitempty"`
	IsDeleted         bool                       `json:"isDeleted,omitempty"`
	CustomFieldData   []CustomFieldValueResponse `json:"customFieldData,omitempty"`
}

type ObjectGetResponse struct {
	Result string         `json:"result"`
	Object ObjectResponse `json:"object"`
}

type ObjectListResponse struct {
	Result  string           `json:"result"`
	Objects []ObjectResponse `json:"objects"`
}
