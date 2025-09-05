package goplanfix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-planfix/models"
	"go-planfix/utils"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func (c *Client) GetFileByID(ctx context.Context, id int, fields string) (*models.FileGetResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.FileByIDEndpoint, id)

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

	var response models.FileGetResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateFile(ctx context.Context, id int, request *models.FileUpdateRequest) (*models.FileUpdateResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.FileByIDEndpoint, id)

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.FileUpdateResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DeleteFile(ctx context.Context, id int) error {
	requestURL := c.APIBase + fmt.Sprintf(utils.FileByIDEndpoint, id)

	req, err := http.NewRequestWithContext(ctx, "DELETE", requestURL, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

func (c *Client) DownloadFile(ctx context.Context, id int) (io.ReadCloser, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.FileDownloadEndpoint, id)

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		resp.Body.Close()
		return nil, utils.ParseAPIError(resp)
	}

	return resp.Body, nil
}

func (c *Client) AttachFileToTask(ctx context.Context, fileID int, taskID int) error {
	requestURL := c.APIBase + fmt.Sprintf(utils.FileAttachTaskEndpoint, fileID)

	queryParams := url.Values{}
	queryParams.Set("id", fmt.Sprintf("%d", taskID))
	requestURL += "?" + queryParams.Encode()

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

func (c *Client) AttachFileToContact(ctx context.Context, fileID int, contactID int) error {
	requestURL := c.APIBase + fmt.Sprintf(utils.FileAttachContactEndpoint, fileID)

	queryParams := url.Values{}
	queryParams.Set("id", fmt.Sprintf("%d", contactID))
	requestURL += "?" + queryParams.Encode()

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

func (c *Client) AttachFileToProject(ctx context.Context, fileID int, projectID int) error {
	requestURL := c.APIBase + fmt.Sprintf(utils.FileAttachProjectEndpoint, fileID)

	queryParams := url.Values{}
	queryParams.Set("id", fmt.Sprintf("%d", projectID))
	requestURL += "?" + queryParams.Encode()

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

func (c *Client) UploadFile(ctx context.Context, filePath string) (*models.FileUploadResponse, error) {
	requestURL := c.APIBase + utils.FileUploadEndpoint

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		return nil, err
	}

	w.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, &b)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	var response models.FileUploadResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UploadFileByLink(ctx context.Context, request *models.FileUploadRequest) (*models.FileUploadResponse, error) {
	requestURL := c.APIBase + utils.FileFromURLEndpoint

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	var response models.FileUploadResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
