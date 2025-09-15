package goplanfix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/whatcrm/go-planfix/models"
	"github.com/whatcrm/go-planfix/utils"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) GetTaskByID(ctx context.Context, taskID int, fields ...string) (*models.Task, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.TaskByIDEndpoint, taskID)

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

	// TODO: Parse task data from response
	task := &models.Task{}
	return task, nil
}

func (c *Client) UpdateTask(ctx context.Context, taskID int, task *models.Task) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.TaskByIDEndpoint, taskID)

	jsonBody, err := json.Marshal(task)
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

func (c *Client) GetListTasks(ctx context.Context, request *models.TaskListRequest) (*models.PaginatedResponse, error) {
	requestURL := c.APIBase + utils.TaskListEndpoint

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

func (c *Client) CreateTask(ctx context.Context, task *models.Task) (*models.Response, error) {
	requestURL := c.APIBase + utils.TaskEndpoint

	jsonBody, err := json.Marshal(task)
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

func (c *Client) GetTaskComments(ctx context.Context, taskID int, request *models.CommentListRequest) (*models.TaskComments, error) {
	requestURL := c.APIBase + utils.TaskCommentsEndpoint + strconv.Itoa(taskID) + "/list"

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.TaskComments
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) AddCommentTask(ctx context.Context, taskID int, comment *models.CommentRequest) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.TaskCommentsEndpoint, taskID)

	jsonBody, err := json.Marshal(comment)
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

func (c *Client) UpdateCommentTask(ctx context.Context, taskID int, commentID int, comment *models.CommentRequest) error {
	requestURL := c.APIBase + utils.TaskCommentsEndpoint + strconv.Itoa(taskID) + "/" + strconv.Itoa(commentID)

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

func (c *Client) AddDataTagNewComment(ctx context.Context, taskID int, comment *models.DataTagEntryRequest) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.TaskDataTagEndpoint, taskID)

	jsonBody, err := json.Marshal(comment)
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

func (c *Client) AddDataTagExistingComment(ctx context.Context, taskID int, commentID int, comment *models.DataTagEntryRequest) (*models.DataTagEntryCommentResponse, error) {
	requestURL := c.APIBase + utils.TaskDataTagEndpoint + strconv.Itoa(taskID) + "/" + strconv.Itoa(commentID)

	jsonBody, err := json.Marshal(comment)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DataTagEntryCommentResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetTaskFiles(ctx context.Context, taskID int, onlyFromDescription bool) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.TaskFilesEndpoint, taskID)

	if onlyFromDescription {
		queryParams := url.Values{}
		queryParams.Set("onlyFromDescription", "true")
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

func (c *Client) GetTaskTemplates(ctx context.Context, offset int, pageSize int, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.TaskTemplatesEndpoint

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

func (c *Client) GetRecurringTasks(ctx context.Context, offset int, pageSize int, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.TaskRecurringEndpoint

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

func (c *Client) GetTaskFilters(ctx context.Context) (*models.Response, error) {
	requestURL := c.APIBase + utils.TaskFiltersEndpoint

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, nil)
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

func (c *Client) GetTaskChecklist(ctx context.Context, taskID int, request *models.ChecklistRequest) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.TaskChecklistEndpoint, taskID)

	jsonBody, err := json.Marshal(request)
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
