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
)

func (c *Client) GetDataTag(ctx context.Context, id int, fields string) (*models.DataTagGetResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.DataTagByIDEndpoint, id)

	queryParams := url.Values{}
	if fields != "" {
		queryParams.Set("fields", fields)
	}

	if len(queryParams) > 0 {
		requestURL += "?" + queryParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.DataTagGetResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetDataTags(ctx context.Context, request *models.DataTagListRequest) (*models.DataTagListResponse, error) {
	requestURL := c.APIBase + utils.DataTagListEndpoint

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DataTagListResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetDataTagEntry(ctx context.Context, key int, fields string) (*models.DataTagEntryGetResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.DataTagEntryByKeyEndpoint, key)

	queryParams := url.Values{}
	if fields != "" {
		queryParams.Set("fields", fields)
	}

	if len(queryParams) > 0 {
		requestURL += "?" + queryParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.DataTagEntryGetResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateDataTagEntry(ctx context.Context, key int, request *models.DataTagEntryUpdateRequest, silent bool) error {
	requestURL := c.APIBase + fmt.Sprintf(utils.DataTagEntryByKeyEndpoint, key)

	queryParams := url.Values{}
	if silent {
		queryParams.Set("silent", "true")
	}

	if len(queryParams) > 0 {
		requestURL += "?" + queryParams.Encode()
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

func (c *Client) DeleteDataTagEntry(ctx context.Context, key int) error {
	requestURL := c.APIBase + fmt.Sprintf(utils.DataTagEntryByKeyEndpoint, key)

	req, err := http.NewRequestWithContext(ctx, "DELETE", requestURL, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

func (c *Client) GetListDataTagEntries(ctx context.Context, id int, request *models.DataTagEntryListRequest) (*models.DataTagEntryListResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.DataTagEntriesListEndpoint, id)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DataTagEntryListResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
