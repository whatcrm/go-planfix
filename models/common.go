package models

type TimePoint struct {
	Date               string `json:"date,omitempty"`
	Time               string `json:"time,omitempty"`
	Datetime           string `json:"datetime,omitempty"`
	DateTimeUtcSeconds string `json:"dateTimeUtcSeconds,omitempty"`
}

type Response struct {
	Result string `json:"result"`
	ID     int    `json:"id,omitempty"`
}

type ResponseError struct {
	Result string `json:"result"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}

type PaginatedResponse struct {
	Status    string      `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	Page      int         `json:"page,omitempty"`
	PageSize  int         `json:"pageSize,omitempty"`
	Total     int         `json:"total,omitempty"`
	TotalPage int         `json:"totalPage,omitempty"`
}

type BaseEntity struct {
	ID   int64  `json:"id"`
	Name string `json:"name,omitempty"`
}

type ComplexContactFilter struct {
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

// Common reusable small types
type PersonResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PeopleResponse struct {
	Users  []PersonResponse `json:"users,omitempty"`
	Groups []GroupResponse  `json:"groups,omitempty"`
}

type GroupResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Common list params, intended to be embedded in request structs
type CommonListParams struct {
	Offset   int    `json:"offset,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Fields   string `json:"fields,omitempty"`
	SourceId string `json:"sourceId,omitempty"`
}

// Generic complex filter (without subfilter). Use specialized types when subfilter is required.
type ComplexFilter struct {
	Type     int         `json:"type"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
	Field    int         `json:"field,omitempty"`
}
