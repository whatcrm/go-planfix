package goplanfix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-planfix/models"
	"go-planfix/utils"
	"net/http"
)

func (c *Client) GetObject(ctx context.Context, objectID int) (*models.Object, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ObjectByIDEndpoint, objectID)

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.Response
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	// TODO: Parse object data from response
	object := &models.Object{}
	return object, nil
}

func (c *Client) GetObjects(ctx context.Context, request *models.ObjectListRequest) (*models.Response, error) {
	requestURL := c.APIBase + utils.ObjectListEndpoint

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

func (c *Client) GetObjectTaskStatuses(ctx context.Context, objectID int, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ObjectStatusesEndpoint, objectID)

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
