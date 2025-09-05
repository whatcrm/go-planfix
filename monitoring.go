package goplanfix

import (
	"context"
	"go-planfix/models"
	"go-planfix/utils"
	"net/http"
)

// Ping checks the availability of the REST service
func (c *Client) Ping(ctx context.Context) (*models.Response, error) {
	requestURL := c.APIBase + utils.PingEndpoint

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
