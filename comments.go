package goplanfix

import (
	"context"
	"fmt"
	"github.com/whatcrm/go-planfix/models"
	"github.com/whatcrm/go-planfix/utils"
	"net/http"
	"net/url"
)

func (c *Client) GetComment(ctx context.Context, id string, fields string, sourceId string) (*models.CommentGetResponse, error) {
	requestURL := c.APIBase + fmt.Sprintf(utils.CommentByIDEndpoint, id)

	queryParams := url.Values{}
	if fields != "" {
		queryParams.Set("fields", fields)
	}
	if sourceId != "" {
		queryParams.Set("sourceId", sourceId)
	}

	if len(queryParams) > 0 {
		requestURL += "?" + queryParams.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	var response models.CommentGetResponse
	err = c.SendWithAccessToken(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
func (c *Client) DeleteComment(ctx context.Context, id string) error {
	requestURL := c.APIBase + fmt.Sprintf(utils.CommentByIDEndpoint, id)

	req, err := http.NewRequestWithContext(ctx, "DELETE", requestURL, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}
