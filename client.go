package goplanfix

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"go-planfix/utils"
)

type Client struct {
	Client         *http.Client
	APIBase        string
	Token          string
	tokenExpiresAt time.Time
}

func NewClient(token string) *Client {
	return &Client{
		Client:  &http.Client{Timeout: utils.DefaultTimeout},
		APIBase: utils.BaseURL,
		Token:   token,
	}
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
	)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err = c.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return utils.ParseAPIError(resp)
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err := io.Copy(w, resp.Body)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) SendWithAccessToken(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
	)

	if req.Header.Get("Content-Type") == "" && req.Method != "DELETE" {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err = c.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return utils.ParseAPIError(resp)
	}

	if v == nil {
		return nil
	}

	if w, ok := v.(io.Writer); ok {
		_, err = io.Copy(w, resp.Body)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
