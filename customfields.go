package goplanfix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/whatcrm/go-planfix/models"
	"github.com/whatcrm/go-planfix/utils"
	"net/http"
)

func (c *Client) GetCustomFieldTypes(ctx context.Context) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldTypesEndpoint

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) GetCustomFieldsForDataTag(ctx context.Context, dataTagID int, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.CustomFieldDataTagEndpoint, dataTagID)

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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
func (c *Client) GetCustomFieldsForDirectory(ctx context.Context, directoryID int, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.CustomFieldDirectoryEndpoint, directoryID)

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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
func (c *Client) GetMainCustomFields(ctx context.Context, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldMainEndpoint

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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
func (c *Client) CreateMainCustomField(ctx context.Context, field *models.CustomField) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldMainEndpoint

	jsonBody, err := json.Marshal(field)
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
func (c *Client) UpdateMainCustomField(ctx context.Context, fieldID int, field *models.CustomField) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.CustomFieldMainByIDEndpoint, fieldID)

	jsonBody, err := json.Marshal(field)
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
func (c *Client) GetTaskCustomFields(ctx context.Context, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldTaskEndpoint

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) GetCustomFieldsForTaskByID(ctx context.Context, taskID int, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.CustomFieldTaskByIDEndpoint, taskID)

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) CreateTaskCustomField(ctx context.Context, field *models.CustomField) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldTaskEndpoint

	jsonBody, err := json.Marshal(field)
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

func (c *Client) GetTaskCustomFieldSets(ctx context.Context) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldTaskGroupEndpoint

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) CreateTaskCustomFieldSet(ctx context.Context, fieldSet *models.CustomFieldSet) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldTaskGroupEndpoint

	jsonBody, err := json.Marshal(fieldSet)
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

func (c *Client) GetContactCustomFields(ctx context.Context, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldContactEndpoint

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) GetCustomFieldsForContact(ctx context.Context, contactID string, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.CustomFieldContactByIDEndpoint, contactID)

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) CreateContactCustomField(ctx context.Context, field *models.CustomField) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldContactEndpoint

	jsonBody, err := json.Marshal(field)
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

func (c *Client) GetContactCustomFieldSets(ctx context.Context) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldContactGroupEndpoint

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) CreateContactCustomFieldSet(ctx context.Context, fieldSet *models.CustomFieldSet) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldContactGroupEndpoint

	jsonBody, err := json.Marshal(fieldSet)
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

func (c *Client) GetProjectCustomFields(ctx context.Context, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldProjectEndpoint

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) GetCustomFieldsForProject(ctx context.Context, projectID int, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.CustomFieldProjectByIDEndpoint, projectID)

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) CreateProjectCustomField(ctx context.Context, field *models.CustomField) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldProjectEndpoint

	jsonBody, err := json.Marshal(field)
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

func (c *Client) GetProjectCustomFieldSets(ctx context.Context) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldProjectGroupEndpoint

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) CreateProjectCustomFieldSet(ctx context.Context, fieldSet *models.CustomFieldSet) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldProjectGroupEndpoint

	jsonBody, err := json.Marshal(fieldSet)
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

func (c *Client) GetUserCustomFields(ctx context.Context, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldUserEndpoint

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) GetCustomFieldsForUser(ctx context.Context, userID string, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.CustomFieldUserByIDEndpoint, userID)

	if len(fields) > 0 {
		queryParams := map[string]string{
			"fields": fields[0],
		}
		requestURL += "?"
		for k, v := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", k, v)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) CreateUserCustomField(ctx context.Context, field *models.CustomField) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldUserEndpoint

	jsonBody, err := json.Marshal(field)
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

func (c *Client) GetUserCustomFieldSets(ctx context.Context) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldUserGroupEndpoint

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
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

func (c *Client) CreateUserCustomFieldSet(ctx context.Context, fieldSet *models.CustomFieldSet) (*models.Response, error) {
	requestURL := c.APIBase + utils.CustomFieldUserGroupEndpoint

	jsonBody, err := json.Marshal(fieldSet)
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
