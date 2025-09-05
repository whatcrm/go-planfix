package models

import "time"

type CustomField struct {
	ID              int64             `json:"id"`
	Name            string            `json:"name"`
	Names           map[string]string `json:"names,omitempty"`
	Type            int64             `json:"type"`
	ObjectType      int64             `json:"objectType"`
	GroupId         int64             `json:"groupId,omitempty"`
	DirectoryId     int64             `json:"directoryId,omitempty"`
	DirectoryFields []int64           `json:"directoryFields,omitempty"`
	EnumValues      []string          `json:"enumValues,omitempty"`
	MainValue       string            `json:"mainValue,omitempty"`
}

type CustomFieldGroup struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	EntityType  string    `json:"entityType,omitempty"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

type CustomFieldValueResponse struct {
	Field *CustomField `json:"field"`
	Value interface{}  `json:"value"`
}

type CustomFieldValueRequest struct {
	Field *BaseEntity `json:"field"`
	Value interface{} `json:"value"`
}

type CustomFieldSet struct {
	Name string `json:"name"`
}
