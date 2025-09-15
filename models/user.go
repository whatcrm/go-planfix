package models

// UserResponse представляет ответ с пользователем
type UserResponse struct {
	ID                int                        `json:"id,omitempty"`
	SourceObjectId    string                     `json:"sourceObjectId,omitempty"`
	SourceDataVersion string                     `json:"sourceDataVersion,omitempty"`
	Name              string                     `json:"name,omitempty"`
	Email             string                     `json:"email,omitempty"`
	Position          *BaseEntity                `json:"position,omitempty"`
	Group             *GroupResponse             `json:"group,omitempty"`
	IsActive          bool                       `json:"isActive,omitempty"`
	IsDeleted         bool                       `json:"isDeleted,omitempty"`
	CustomFieldData   []CustomFieldValueResponse `json:"customFieldData,omitempty"`
}

// UserGetResponse представляет ответ на получение пользователя
type UserGetResponse struct {
	Result string       `json:"result"`
	User   UserResponse `json:"user"`
}

// UserListResponse представляет ответ со списком пользователей
type UserListResponse struct {
	Result string         `json:"result"`
	Users  []UserResponse `json:"users"`
}

type User struct {
	Name         string `json:"name"`
	Midname      string `json:"midname"`
	Lastname     string `json:"lastname"`
	Gender       string `json:"gender"`
	LanguageCode string `json:"languageCode"`
	BirthDate    struct {
		Date     string `json:"date"`
		Time     string `json:"time"`
		Datetime string `json:"datetime"`
	} `json:"birthDate"`
	Groups []struct {
		Id int `json:"id"`
	} `json:"groups"`
	Role            string        `json:"role"`
	Login           string        `json:"login"`
	Password        string        `json:"password"`
	Email           string        `json:"email"`
	SecondaryEmails []interface{} `json:"secondaryEmails"`
	Status          string        `json:"status"`
	Phones          []struct {
		Number       string `json:"number"`
		MaskedNumber string `json:"maskedNumber"`
		Type         int    `json:"type"`
	} `json:"phones"`
	Position struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"position"`
}

type UserListRequest struct {
	CommonListParams
	OnlyActive bool                `json:"onlyActive,omitempty"`
	PrefixedId bool                `json:"prefixedId,omitempty"`
	Filters    []ComplexUserFilter `json:"filters,omitempty"`
}

type ComplexUserFilter struct {
	Type     int         `json:"type"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
	Field    int         `json:"field,omitempty"`
}
