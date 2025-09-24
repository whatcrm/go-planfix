package goplanfix

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/whatcrm/go-planfix/models"
	"github.com/whatcrm/go-planfix/utils"
)

// ChatSendNewMessage отправляет новое сообщение в чат из стороннего чата в планфикс
func (c *Client) ChatSendNewMessage(ctx context.Context, req *models.ChatMessageRequest) error {
	req.Cmd = "newMessage"

	// Подготавливаем данные для отправки
	data := url.Values{}
	data.Set("cmd", req.Cmd)
	data.Set("providerId", req.ProviderID)
	data.Set("chatId", req.ChatID)
	data.Set("planfix_token", req.PlanfixToken)
	data.Set("message", req.Message)
	data.Set("contactId", req.ContactID)
	data.Set("contactName", req.ContactName)

	if req.Channel != "" {
		data.Set("channel", req.Channel)
	}
	if req.Title != "" {
		data.Set("title", req.Title)
	}
	if req.ContactLastName != "" {
		data.Set("contactLastName", req.ContactLastName)
	}
	if req.ContactIco != "" {
		data.Set("contactIco", req.ContactIco)
	}
	if req.ContactEmail != "" {
		data.Set("contactEmail", req.ContactEmail)
	}
	if req.ContactPhone != "" {
		data.Set("contactPhone", req.ContactPhone)
	}
	if req.ContactData != "" {
		data.Set("contactData", req.ContactData)
	}
	if req.IsEcho {
		data.Set("isEcho", "true")
	}
	if req.UserEmail != "" {
		data.Set("userEmail", req.UserEmail)
	}

	// Добавляем вложения
	for name, url := range req.Attachments {
		data.Add("attachments[name]", name)
		data.Add("attachments[url]", url)
	}

	// Добавляем дополнительные данные задачи
	for field, value := range req.TaskData {
		data.Set("data_"+field, value)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.APIBase+utils.ChatAPIEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	err = c.SendWithAccessToken(httpReq, nil)
	return err
}

// ChatGetTask получает номер задачи по chatId
func (c *Client) ChatGetTask(ctx context.Context, req *models.ChatGetTaskRequest) (*models.ChatTaskResponse, error) {
	req.Cmd = "getTask"

	data := url.Values{}
	data.Set("cmd", req.Cmd)
	data.Set("providerId", req.ProviderID)
	data.Set("planfix_token", req.PlanfixToken)
	data.Set("chatId", req.ChatID)

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.APIBase+utils.ChatAPIEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	var response models.ChatTaskResponse
	err = c.SendWithAccessToken(httpReq, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ChatGetContact получает номер контакта по contactId
func (c *Client) ChatGetContact(ctx context.Context, req *models.ChatGetContactRequest) (*models.ChatContactResponse, error) {
	req.Cmd = "getContact"

	data := url.Values{}
	data.Set("cmd", req.Cmd)
	data.Set("providerId", req.ProviderID)
	data.Set("planfix_token", req.PlanfixToken)
	data.Set("contactId", req.ContactID)

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.APIBase+utils.ChatAPIEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	var response models.ChatContactResponse
	err = c.SendWithAccessToken(httpReq, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ChatUpdateContact обновляет данные контакта
func (c *Client) ChatUpdateContact(ctx context.Context, req *models.ChatUpdateContactRequest) error {
	req.Cmd = "updateContact"

	data := url.Values{}
	data.Set("cmd", req.Cmd)
	data.Set("providerId", req.ProviderID)
	data.Set("planfix_token", req.PlanfixToken)
	data.Set("contactId", req.ContactID)
	data.Set("contactName", req.ContactName)

	if req.ContactLastName != "" {
		data.Set("contactLastName", req.ContactLastName)
	}
	if req.ContactIco != "" {
		data.Set("contactIco", req.ContactIco)
	}
	if req.ContactEmail != "" {
		data.Set("contactEmail", req.ContactEmail)
	}
	if req.ContactPhone != "" {
		data.Set("contactPhone", req.ContactPhone)
	}
	if req.ContactData != "" {
		data.Set("contactData", req.ContactData)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.APIBase+utils.ChatAPIEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	err = c.SendWithAccessToken(httpReq, nil)
	return err
}

// ChatMessageStatus обновляет статус сообщения
func (c *Client) ChatMessageStatus(ctx context.Context, req *models.ChatMessageStatusRequest) error {
	req.Cmd = "messageStatus"

	data := url.Values{}
	data.Set("cmd", req.Cmd)
	data.Set("providerId", req.ProviderID)
	data.Set("planfix_token", req.PlanfixToken)
	data.Set("messageId", req.MessageID)
	data.Set("messageStatus", req.MessageStatus)

	if req.MessageStatusText != "" {
		data.Set("messageStatusText", req.MessageStatusText)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.APIBase+utils.ChatAPIEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	err = c.SendWithAccessToken(httpReq, nil)
	return err
}

// ChatSendMessageToExternal отправляет сообщение во внешний чат
// Этот метод используется для отправки сообщений из ПланФикса в сторонний чат
func (c *Client) ChatSendMessageToExternal(ctx context.Context, externalURL string, req *models.ChatNewMessageToExternalRequest) (*models.ChatNewMessageResponse, error) {
	req.Cmd = "newMessage"

	data := url.Values{}
	data.Set("cmd", req.Cmd)
	data.Set("providerId", req.ProviderID)
	data.Set("token", req.Token)
	data.Set("message", req.Message)
	data.Set("messageId", req.MessageID)
	data.Set("userName", req.UserName)
	data.Set("userLastName", req.UserLastName)
	data.Set("userIco", req.UserIco)
	data.Set("taskEmail", req.TaskEmail)

	if req.ChatID != "" {
		data.Set("chatId", req.ChatID)
	}
	if req.ContactPhone != "" {
		data.Set("contactPhone", req.ContactPhone)
	}
	if req.Channel != "" {
		data.Set("channel", req.Channel)
	}

	// Добавляем вложения
	for name, url := range req.Attachments {
		data.Add("attachments[name]", name)
		data.Add("attachments[url]", url)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", externalURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	var response models.ChatNewMessageResponse
	err = c.SendWithAccessToken(httpReq, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ChatSendFirstMessageToExternal отправляет первое сообщение во внешний чат по телефону
func (c *Client) ChatSendFirstMessageToExternal(ctx context.Context, externalURL string, req *models.ChatNewMessageToExternalRequest) (*models.ChatNewMessageResponse, error) {
	req.Cmd = "newMessage"

	data := url.Values{}
	data.Set("cmd", req.Cmd)
	data.Set("providerId", req.ProviderID)
	data.Set("contactPhone", req.ContactPhone)
	data.Set("token", req.Token)
	data.Set("message", req.Message)
	data.Set("messageId", req.MessageID)
	data.Set("userName", req.UserName)
	data.Set("userLastName", req.UserLastName)
	data.Set("userIco", req.UserIco)
	data.Set("taskEmail", req.TaskEmail)

	if req.Channel != "" {
		data.Set("channel", req.Channel)
	}

	// Добавляем вложения
	for name, url := range req.Attachments {
		data.Add("attachments[name]", name)
		data.Add("attachments[url]", url)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", externalURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	var response models.ChatNewMessageResponse
	err = c.SendWithAccessToken(httpReq, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
