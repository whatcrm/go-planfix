package goplanfix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"go-planfix/models"
	"go-planfix/utils"
)

func (c *Client) GetContactByID(ctx context.Context, contactID int, fields ...string) (*models.ContactGetResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ContactByIDEndpoint, contactID)

	if len(fields) > 0 {
		queryParams := url.Values{}
		queryParams.Set("fields", strings.Join(fields, ", "))
		requestURL += "?" + queryParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.ContactGetResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateContact(ctx context.Context, contactID string, contact *models.Contact) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ContactByIDEndpoint, contactID)

	jsonBody, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetContactOrCompanyList(ctx context.Context, fields []string, pageSize int) (*models.ContactListResponse, error) {
	requestURL := c.APIBase + utils.ContactListEndpoint

	request := &models.ContactListRequest{
		Offset:   0,
		PageSize: pageSize,
		Fields:   strings.Join(fields, ","),
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ContactListResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateContactOrCompany(ctx context.Context, contact *models.Contact) (*models.Response, error) {
	requestURL := c.APIBase + utils.ContactEndpoint
	jsonBody, err := json.Marshal(contact)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var contactResponse models.Response
	err = c.SendWithAccessToken(req, &contactResponse)
	if err != nil {
		fmt.Printf("Error creating contact: %v\n", err)
		return nil, err
	}

	return &contactResponse, nil
}

func (c *Client) ImportContacts(ctx context.Context, contacts []*models.Contact) (*models.Response, error) {
	requestURL := c.APIBase + utils.ContactImportEndpoint

	jsonBody, err := json.Marshal(contacts)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetContactComments(ctx context.Context, contactID string, request *models.ContactCommentListRequest) (*models.CommentListResponse, error) {
	requestURL := c.APIBase + utils.ContactCommentEndpoint + contactID + "/list"

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.CommentListResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) AddCommentContact(ctx context.Context, contactID string, comment *models.Comment) (*models.Response, error) {
	requestRL := c.APIBase + fmt.Sprintf(utils.ContactCommentEndpoint, contactID)

	jsonBody, err := json.Marshal(comment)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestRL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) AddDataTagEntryWithANewComment(ctx context.Context, contactID string, tag *models.DataTag) (*models.DataTagResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ContactDataTagEndpoint, contactID)

	jsonBody, err := json.Marshal(tag)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DataTagResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) AddDataTagEntryWithExistingComment(ctx context.Context, contactID string, commentID int, tag *models.DataTag) (*models.DataTagResponse, error) {
	requestURL := c.APIBase + utils.ContactDataTagEndpoint + contactID + "/" + strconv.Itoa(commentID)

	jsonBody, err := json.Marshal(tag)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DataTagResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) UpdateComment(ctx context.Context, contactID string, commentID int, comment *models.Comment) error {
	requestURL := c.APIBase + utils.ContactCommentEndpoint + "/" + contactID + "/" + strconv.Itoa(commentID)

	jsonBody, err := json.Marshal(comment)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

func (c *Client) GetContactFiles(ctx context.Context, contactID int) (*models.PaginatedResponse, error) {
	endpoint := utils.ContactEndpoint + "/" + strconv.Itoa(contactID) + "/files"
	requestURL := c.APIBase + endpoint

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.PaginatedResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetContactTemplates(ctx context.Context) (*models.PaginatedResponse, error) {
	requestURL := c.APIBase + utils.ContactTemplatesEndpoint

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.PaginatedResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetContactGroups(ctx context.Context) (*models.PaginatedResponse, error) {
	requestURL := c.APIBase + utils.ContactGroupsEndpoint

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.PaginatedResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetContactFilters(ctx context.Context) (*models.PaginatedResponse, error) {
	requestURL := c.APIBase + utils.ContactFiltersEndpoint

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.PaginatedResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
