package goplanfix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-planfix/models"
	"go-planfix/utils"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) GetProjectByID(ctx context.Context, projectID int, fields ...string) (*models.Project, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ProjectByIDEndpoint, projectID)

	if len(fields) > 0 {
		queryParams := url.Values{}
		queryParams.Set("fields", fields[0])
		requestURL += "?" + queryParams.Encode()
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

	// TODO: Parse project data from response
	project := &models.Project{}
	return project, nil
}

func (c *Client) UpdateProject(ctx context.Context, projectID int, project *models.Project) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ProjectByIDEndpoint, projectID)

	jsonBody, err := json.Marshal(project)
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

func (c *Client) GetListsProjects(ctx context.Context, request *models.ProjectListRequest) (*models.PaginatedResponse, error) {
	requestURL := c.APIBase + utils.ProjectListEndpoint

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
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

func (c *Client) CreateProject(ctx context.Context, project *models.Project) (*models.Response, error) {
	requestURL := c.APIBase + utils.ProjectEndpoint

	jsonBody, err := json.Marshal(project)
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

func (c *Client) GetProjectFiles(ctx context.Context, projectID int, offset int, pageSize int) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ProjectFilesEndpoint, projectID)

	queryParams := url.Values{}
	if offset > 0 {
		queryParams.Set("offset", strconv.Itoa(offset))
	}
	if pageSize > 0 {
		queryParams.Set("pageSize", strconv.Itoa(pageSize))
	}
	if len(queryParams) > 0 {
		requestURL += "?" + queryParams.Encode()
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

func (c *Client) GetProjectTemplates(ctx context.Context, offset int, pageSize int, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.ProjectTemplatesEndpoint

	queryParams := url.Values{}
	if offset > 0 {
		queryParams.Set("offset", strconv.Itoa(offset))
	}
	if pageSize > 0 {
		queryParams.Set("pageSize", strconv.Itoa(pageSize))
	}
	if len(fields) > 0 {
		queryParams.Set("fields", fmt.Sprintf("%s", fields[0]))
	}
	if len(queryParams) > 0 {
		requestURL += "?" + queryParams.Encode()
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

func (c *Client) GetProjectGroups(ctx context.Context, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.ProjectGroupsEndpoint

	if len(fields) > 0 {
		queryParams := url.Values{}
		queryParams.Set("fields", fields[0])
		requestURL += "?" + queryParams.Encode()
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
