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
)

func (c *Client) GetUserByID(ctx context.Context, userID string, fields ...string) (*models.User, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.UserByIDEndpoint, userID)

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

	// TODO: Parse user data from response
	user := &models.User{}
	return user, nil
}

func (c *Client) UpdateUser(ctx context.Context, userID string, user *models.User) (*models.Response, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.UserByIDEndpoint, userID)

	jsonBody, err := json.Marshal(user)
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

func (c *Client) GetListUsers(ctx context.Context, request *models.UserListRequest) (*models.PaginatedResponse, error) {
	requestURL := c.APIBase + utils.UserListEndpoint

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

func (c *Client) CreateUser(ctx context.Context, user *models.User) (*models.Response, error) {
	requestURL := c.APIBase + utils.UserEndpoint

	jsonBody, err := json.Marshal(user)
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
func (c *Client) GetUserPositions(ctx context.Context, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.UserPositionsEndpoint

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

func (c *Client) GetListUserGroups(ctx context.Context, fields ...string) (*models.Response, error) {
	requestURL := c.APIBase + utils.UserGroupsEndpoint

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
