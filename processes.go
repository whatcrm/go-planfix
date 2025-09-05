package goplanfix

import (
	"context"
	"fmt"
	"go-planfix/models"
	"go-planfix/utils"
	"net/http"
)

func (c *Client) GetTaskProcesses(ctx context.Context, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.ProcessTaskEndpoint

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

func (c *Client) GetTaskStatusesForProcess(ctx context.Context, processID int, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ProcessTaskStatusesEndpoint, processID)

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

func (c *Client) GetContactProcesses(ctx context.Context, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.ProcessContactEndpoint

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
