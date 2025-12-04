package models

// TaskResponse представляет ответ с задачей
type TaskResponse struct {
	ID                 int                        `json:"id,omitempty"`
	SourceObjectId     string                     `json:"sourceObjectId,omitempty"`
	SourceDataVersion  string                     `json:"sourceDataVersion,omitempty"`
	Name               string                     `json:"name,omitempty"`
	Description        string                     `json:"description,omitempty"`
	Status             *TaskStatus                `json:"status,omitempty"`
	Owner              *PersonResponse            `json:"owner,omitempty"`
	Project            *BaseEntity                `json:"project,omitempty"`
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

// TaskGetResponse представляет ответ на получение задачи
type TaskGetResponse struct {
	Result string       `json:"result"`
	Task   TaskResponse `json:"task"`
}

// TaskListResponse представляет ответ со списком задач
type TaskListResponse struct {
	Result string         `json:"result"`
	Tasks  []TaskResponse `json:"tasks"`
}

// TaskStatus представляет статус задачи
type TaskStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TaskComments struct {
	Result   string            `json:"result"`
	Comments []CommentResponse `json:"comments"`
}

type Task struct {
	ID          int            `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	Status      string         `json:"status,omitempty"`
	Priority    string         `json:"priority,omitempty"`
	ProjectID   int            `json:"projectId,omitempty"`
	AssigneeID  int            `json:"assigneeId,omitempty"`
	Assignees   *TaskAssignees `json:"assignees,omitempty"`
	DueDate     string         `json:"dueDate,omitempty"`
	CreatedAt   string         `json:"createdAt,omitempty"`
	UpdatedAt   string         `json:"updatedAt,omitempty"`
}

type TaskAssignees struct {
	Users []TaskAssignee `json:"users,omitempty"`
}

type TaskAssignee struct {
	ID string `json:"id"`
}

type TaskListRequest struct {
	CommonListParams
	FilterId string              `json:"filterId,omitempty"`
	Filters  []ComplexTaskFilter `json:"filters,omitempty"`
}

type ComplexTaskFilter struct {
	Type      int         `json:"type"`
	Operator  string      `json:"operator"`
	Value     interface{} `json:"value"`
	Field     int         `json:"field,omitempty"`
	Subfilter *struct {
		DataTagId int `json:"dataTagId"`
		Filter    struct {
			Type  int `json:"type"`
			Field int `json:"field"`
		} `json:"filter"`
	} `json:"subfilter,omitempty"`
}

type ChecklistRequest struct {
	CommonListParams
}
