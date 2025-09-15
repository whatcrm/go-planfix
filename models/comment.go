package models

import "time"

type Comment struct {
	ID         int       `json:"id"`
	Text       string    `json:"text"`
	AuthorID   int       `json:"authorId"`
	EntityID   int       `json:"entityId"`
	EntityType string    `json:"entityType"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

type CommentListRequest struct {
	Offset      int    `json:"offset,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	Fields      string `json:"fields,omitempty"`
	TypeList    string `json:"typeList,omitempty"`
	ResultOrder []struct {
		Field     string `json:"field"`
		Direction string `json:"direction"`
	} `json:"resultOrder,omitempty"`
}

type CommentRequest struct {
	SourceId          string           `json:"sourceId,omitempty"`
	SourceObjectId    string           `json:"sourceObjectId,omitempty"`
	SourceDataVersion string           `json:"sourceDataVersion,omitempty"`
	DateTime          *TimePoint       `json:"dateTime,omitempty"`
	Description       string           `json:"description,omitempty"`
	Owner             *PersonRequest   `json:"owner,omitempty"`
	IsPinned          bool             `json:"isPinned,omitempty"`
	IsHidden          bool             `json:"isHidden,omitempty"`
	Recipients        *NotifiedRequest `json:"recipients,omitempty"`
	Files             []FileRequest    `json:"files,omitempty"`
}

type PersonRequest struct {
	ID string `json:"id"`
}

type NotifiedRequest struct {
	Users  []PersonRequest `json:"users,omitempty"`
	Groups []GroupRequest  `json:"groups,omitempty"`
	Roles  []string        `json:"roles,omitempty"`
}

type GroupRequest struct {
	ID int `json:"id"`
}

type CommentResponse struct {
	ID                   int64                `json:"id"`
	SourceObjectId       string               `json:"sourceObjectId,omitempty"`
	SourceDataVersion    string               `json:"sourceDataVersion,omitempty"`
	DateTime             *TimePoint           `json:"dateTime,omitempty"`
	Type                 string               `json:"type,omitempty"`
	FromType             string               `json:"fromType,omitempty"`
	Description          string               `json:"description,omitempty"`
	Task                 *BaseEntity          `json:"task,omitempty"`
	Project              *BaseEntity          `json:"project,omitempty"`
	Contact              *PersonResponse      `json:"contact,omitempty"`
	Owner                *PersonResponse      `json:"owner,omitempty"`
	IsDeleted            bool                 `json:"isDeleted,omitempty"`
	IsPinned             bool                 `json:"isPinned,omitempty"`
	IsHidden             bool                 `json:"isHidden,omitempty"`
	IsNotRead            bool                 `json:"isNotRead,omitempty"`
	Recipients           *PeopleResponse      `json:"recipients,omitempty"`
	Reminders            []Reminder           `json:"reminders,omitempty"`
	DataTags             []CommentDataTag     `json:"dataTags,omitempty"`
	Files                []File               `json:"files,omitempty"`
	ChangeTaskStartDate  *ChangeDate          `json:"changeTaskStartDate,omitempty"`
	ChangeTaskExpectDate *ChangeDate          `json:"changeTaskExpectDate,omitempty"`
	ChangeStatus         *CommentChangeStatus `json:"changeStatus,omitempty"`
}

type Reminder struct {
	// TODO: Add reminder fields based on swagger schema
}

type CommentDataTag struct {
	DataTag struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"dataTag"`
	Key int `json:"key"`
}

type ChangeDate struct {
	// TODO: Add change date fields based on swagger schema
}

type CommentChangeStatus struct {
	NewStatusId int `json:"newStatusId"`
	OldStatusId int `json:"oldStatusId"`
}

type CommentGetResponse struct {
	Result  string          `json:"result"`
	Comment CommentResponse `json:"comment"`
}

type CommentListResponse struct {
	Result   string            `json:"result"`
	Comments []CommentResponse `json:"comments"`
}

type ContactCommentListRequest struct {
	CommonListParams
	TypeList    string `json:"typeList,omitempty"`
	ResultOrder []struct {
		Field     string `json:"field"`
		Direction string `json:"direction"`
	} `json:"resultOrder,omitempty"`
}
