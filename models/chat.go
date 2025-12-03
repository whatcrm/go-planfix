package models

import "encoding/json"

type ChatAttachment struct {
	Name string `json:"name" form:"name"`
	URL  string `json:"url" form:"url"`
}

type ChatMessageRequest struct {
	Cmd             string            `json:"cmd" form:"cmd"`                   // тип операции: newMessage
	ProviderID      string            `json:"providerId" form:"providerId"`     // идентификатор сторонней системы
	Channel         string            `json:"channel,omitempty" form:"channel"` // дополнительный идентификатор канала
	ChatID          string            `json:"chatId" form:"chatId"`
	PlanfixToken    string            `json:"planfix_token" form:"planfix_token"`
	Message         string            `json:"message" form:"message"`
	Title           string            `json:"title,omitempty" form:"title"`
	ContactID       string            `json:"contactId" form:"contactId"`
	ContactName     string            `json:"contactName" form:"contactName"`
	ContactLastName string            `json:"contactLastName,omitempty" form:"contactLastName"`
	ContactIco      string            `json:"contactIco,omitempty" form:"contactIco"`
	ContactEmail    string            `json:"contactEmail,omitempty" form:"contactEmail"`
	ContactPhone    string            `json:"contactPhone,omitempty" form:"contactPhone"`
	ContactData     string            `json:"contactData,omitempty" form:"contactData"`
	Attachments     []ChatAttachment  `json:"attachments,omitempty" form:"attachments"`
	IsEcho          bool              `json:"isEcho,omitempty" form:"isEcho"`
	UserEmail       string            `json:"userEmail,omitempty" form:"userEmail"`
	TaskData        map[string]string `json:"-" form:"-"`
}

type ChatGetTaskRequest struct {
	Cmd          string `json:"cmd" form:"cmd"` // тип операции: getTask
	ProviderID   string `json:"providerId" form:"providerId"`
	PlanfixToken string `json:"planfix_token" form:"planfix_token"`
	ChatID       string `json:"chatId" form:"chatId"`
}

type ChatGetContactRequest struct {
	Cmd          string `json:"cmd" form:"cmd"` // тип операции: getContact
	ProviderID   string `json:"providerId" form:"providerId"`
	PlanfixToken string `json:"planfix_token" form:"planfix_token"`
	ContactID    string `json:"contactId" form:"contactId"`
}

type ChatUpdateContactRequest struct {
	Cmd             string `json:"cmd" form:"cmd"` // тип операции: updateContact
	ProviderID      string `json:"providerId" form:"providerId"`
	PlanfixToken    string `json:"planfix_token" form:"planfix_token"`
	ContactID       string `json:"contactId" form:"contactId"`
	ContactName     string `json:"contactName" form:"contactName"`
	ContactLastName string `json:"contactLastName,omitempty" form:"contactLastName"`
	ContactIco      string `json:"contactIco,omitempty" form:"contactIco"`
	ContactEmail    string `json:"contactEmail,omitempty" form:"contactEmail"`
	ContactPhone    string `json:"contactPhone,omitempty" form:"contactPhone"`
	ContactData     string `json:"contactData,omitempty" form:"contactData"`
}

type ChatMessageStatusRequest struct {
	Cmd               string `json:"cmd" form:"cmd"` // тип операции: messageStatus
	ProviderID        string `json:"providerId" form:"providerId"`
	PlanfixToken      string `json:"planfix_token" form:"planfix_token"`
	MessageID         string `json:"messageId" form:"messageId"`
	MessageStatus     string `json:"messageStatus" form:"messageStatus"`
	MessageStatusText string `json:"messageStatusText,omitempty" form:"messageStatusText"`
}

// ChatNewMessageToExternalRequest представляет запрос на отправку сообщения во внешний чат
type ChatNewMessageToExternalRequest struct {
	Cmd          string           `json:"cmd" form:"cmd"` // тип операции: newMessage
	ProviderID   string           `json:"providerId" form:"providerId"`
	ChatID       string           `json:"chatId,omitempty" form:"chatId"`
	ContactPhone string           `json:"contactPhone,omitempty" form:"contactPhone"`
	Channel      string           `json:"channel,omitempty" form:"channel"`
	Token        string           `json:"token" form:"token"`
	Message      string           `json:"message" form:"message"`
	MessageID    string           `json:"messageId" form:"messageId"`
	UserName     string           `json:"userName" form:"userName"`
	UserLastName string           `json:"userLastName" form:"userLastName"`
	UserIco      string           `json:"userIco" form:"userIco"`
	TaskEmail    string           `json:"taskEmail" form:"taskEmail"`
	Attachments  []ChatAttachment `json:"attachments,omitempty" form:"attachments"`
}

type ChatResponse struct {
	Data  map[string]interface{} `json:"data,omitempty"`
	Error string                 `json:"error,omitempty"`
}

type ChatTaskData struct {
	Number         int  `json:"number"`
	StatusIsActive bool `json:"statusIsActive,omitempty"`
}

type ChatTaskResponse struct {
	Data  json.RawMessage `json:"data"`
	Error string          `json:"error,omitempty"`
}

type ChatContactResponse struct {
	Data struct {
		Number int `json:"number"`
	} `json:"data"`
	Error string `json:"error,omitempty"`
}

type ChatNewMessageResponse struct {
	ChatID    string `json:"chatId"`
	ContactID string `json:"contactId"`
	Error     string `json:"error,omitempty"`
}
