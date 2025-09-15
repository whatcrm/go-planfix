package models

import "time"

type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status,omitempty"`
	ManagerID   int       `json:"managerId,omitempty"`
	StartDate   time.Time `json:"startDate,omitempty"`
	EndDate     time.Time `json:"endDate,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

type ProjectListRequest struct {
	CommonListParams
	Filters []ComplexProjectFilter `json:"filters,omitempty"`
}

type ComplexProjectFilter struct {
	Type     int         `json:"type"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
	Field    int         `json:"field,omitempty"`
}

// ProjectResponse представляет ответ с проектом
type ProjectResponse struct {
	ID                 int                        `json:"id,omitempty"`
	SourceObjectId     string                     `json:"sourceObjectId,omitempty"`
	SourceDataVersion  string                     `json:"sourceDataVersion,omitempty"`
	Name               string                     `json:"name,omitempty"`
	Description        string                     `json:"description,omitempty"`
	Status             *BaseEntity                `json:"status,omitempty"`
	Owner              *PersonResponse            `json:"owner,omitempty"`
	Parent             *BaseEntity                `json:"parent,omitempty"`
	Template           *BaseEntity                `json:"template,omitempty"`
	Group              *GroupResponse             `json:"group,omitempty"`
	Counterparty       *PersonResponse            `json:"counterparty,omitempty"`
	StartDate          *TimePoint                 `json:"startDate,omitempty"`
	EndDate            *TimePoint                 `json:"endDate,omitempty"`
	HiddenForEmployees bool                       `json:"hiddenForEmployees,omitempty"`
	HiddenForClients   bool                       `json:"hiddenForClients,omitempty"`
	Overdue            bool                       `json:"overdue,omitempty"`
	IsCloseToDeadline  bool                       `json:"isCloseToDeadline,omitempty"`
	HasEndDate         bool                       `json:"hasEndDate,omitempty"`
	IsDeleted          bool                       `json:"isDeleted,omitempty"`
	Assignees          *PeopleResponse            `json:"assignees,omitempty"`
	Participants       *PeopleResponse            `json:"participants,omitempty"`
	Auditors           *PeopleResponse            `json:"auditors,omitempty"`
	ClientManagers     *PeopleResponse            `json:"clientManagers,omitempty"`
	CustomFieldData    []CustomFieldValueResponse `json:"customFieldData,omitempty"`
}

// ProjectGetResponse представляет ответ на получение проекта
type ProjectGetResponse struct {
	Result  string          `json:"result"`
	Project ProjectResponse `json:"project"`
}

// ProjectListResponse представляет ответ со списком проектов
type ProjectListResponse struct {
	Result   string            `json:"result"`
	Projects []ProjectResponse `json:"projects"`
}
