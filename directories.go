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
	"strings"
)

func (c *Client) GetDirectoryByID(ctx context.Context, id int, fields string) (*models.DirectoryGetResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.DirectoryByIDEndpoint, id)

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

	var response models.DirectoryGetResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetListDirectories(ctx context.Context, request *models.DirectoryRequest) (*models.DirectoryListResponse, error) {
	requestURL := c.APIBase + utils.DirectoryListEndpoint

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DirectoryListResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
func (c *Client) AddDirectoryEntry(ctx context.Context, directoryID int, request *models.DirectoryEntryRequest) (*models.DirectoryEntryCreateResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.DirectoryEntryEndpoint, directoryID)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DirectoryEntryCreateResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
func (c *Client) GetDirectoryEntry(ctx context.Context, directoryID int, key int, fields string) (*models.DirectoryEntryGetResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.DirectoryEntryByKeyEndpoint, directoryID, key)

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

	var response models.DirectoryEntryGetResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
func (c *Client) UpdateDirectoryEntry(ctx context.Context, directoryID int, key int, request *models.DirectoryEntryRequest) error {
	requestURL := c.APIBase + fmt.Sprintf(utils.DirectoryEntryByKeyEndpoint, directoryID, key)

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
func (c *Client) DeleteDirectoryEntry(ctx context.Context, directoryID int, key int) error {
	requestURL := c.APIBase + fmt.Sprintf(utils.DirectoryEntryByKeyEndpoint, directoryID, key)

	req, err := http.NewRequestWithContext(ctx, "DELETE", requestURL, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}
func (c *Client) GetListDirectoryEntries(ctx context.Context, directoryID int, request *models.DirectoryRequest) (*models.DirectoryEntryListResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.DirectoryEntriesListEndpoint, directoryID)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.DirectoryEntryListResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetListDirectoryGroups(ctx context.Context, fields ...string) (*models.DirectoryGroupsResponse, error) {
	requestURL := c.APIBase + utils.DirectoryGroupsEndpoint

	if len(fields) > 0 {
		queryParams := url.Values{}
		queryParams.Set("fields", strings.Join(fields, ", "))
		requestURL += "?" + queryParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.DirectoryGroupsResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetListDirectoryEntriesFilters(ctx context.Context, directoryID string) (*models.DirectoryFiltersResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.DirectoryFiltersEndpoint, directoryID)

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.DirectoryFiltersResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
