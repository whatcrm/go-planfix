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

func (c *Client) GetReportByID(ctx context.Context, id int, fields string) (*models.ReportGetResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ReportByIDEndpoint, id)

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

	var response models.ReportGetResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetReports(ctx context.Context, request *models.ReportListRequest) (*models.ReportListResponse, error) {
	requestURL := c.APIBase + utils.ReportListEndpoint

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ReportListResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
func (c *Client) GetReportSaves(ctx context.Context, id int) (*models.ReportSaveListResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ReportSaveListEndpoint, id)

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.ReportSaveListResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
func (c *Client) GetReportSaveData(ctx context.Context, id int, saveID int, request *models.ReportSaveDataRequest) (*models.ReportSaveDataResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ReportSaveDataEndpoint, id, saveID)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ReportSaveDataResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
func (c *Client) GenerateReport(ctx context.Context, id int, request *models.ReportGenerateRequest) (*models.ReportGenerateResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ReportGenerateEndpoint, id)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.ReportGenerateResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
func (c *Client) GetReportStatus(ctx context.Context, requestID string) (*models.ReportStatusResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.ReportStatusEndpoint, requestID)

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.ReportStatusResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
