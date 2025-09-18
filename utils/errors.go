package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIError struct {
	Result   string `json:"result"`
	Code     int    `json:"code"`
	ErrorMsg string `json:"error"`
}

func (e *APIError) Error() string {
	description := GetErrorDescription(e.Code)
	if description != "" {
		if e.ErrorMsg != "" {
			return fmt.Sprintf("Planfix API Error %d: %s - %s", e.Code, description, e.ErrorMsg)
		}
		return fmt.Sprintf("Planfix API Error %d: %s", e.Code, description)
	}

	if e.ErrorMsg != "" {
		return fmt.Sprintf("API Error %d: %s", e.Code, e.ErrorMsg)
	}
	return fmt.Sprintf("API Error %d: %s", e.Code, e.Result)
}

func (e *APIError) GetDescription() string {
	return GetErrorDescription(e.Code)
}

func GetErrorDescription(code int) string {
	errorMessages := map[int]string{
		0:  "Unknown error",
		1:  "Invalid token",
		2:  "The token is not active",
		5:  "Access to this method for scope is forbidden",
		6:  "Access to this method/object for the user is forbidden",
		10: "The account is blocked for spam",
		11: "Account is frozen",
		16: "User is inactive",
		20: "Subscription is not paid",
		21: "API usage is not available for a free account",
		22: "The daily limit of requests to the REST API is spent",
		23: "The account limit for the number of contacts has been reached",
		30: "Invalid JSON format",
		40: "One of the mandatory method parameters is missing",
		41: "Invalid parameter value",

		// Task errors
		1000: "Task does not exist",
		1001: "Error adding task",
		1002: "Error updating task",

		// Contact errors
		2000: "Contact does not exist",
		2001: "Error adding contact",
		2002: "Error updating data",
		2003: "When attempting to grant access to Planfix, the contact does not have an email set",
		2004: "Invalid email format",
		2005: "The email specified is already taken",
		2006: "Contact's full name is not filled in",

		// Project errors
		3000: "Project does not exist",
		3001: "Error adding project",
		3002: "Error updating project",

		// Employee errors
		4000: "Employee does not exist",
		4001: "Error adding employee",
		4002: "Error updating employee data",
		4003: "The specified email is not unique",
		4004: "Invalid username format",
		4005: "The specified username is not unique",
		4006: "Employee's full name is not filled in",

		// Comment errors
		5000: "Comment does not exist",
		5001: "Error adding comment",
		5002: "Error updating comment data",

		// Data tags and data tag records errors
		6000: "Data tag does not exist",
		6010: "Data tag record does not exist",
		6011: "Error adding data tag entry",
		6012: "Error updating data tag entry",
		6013: "Error deleting data tag entry",

		// File errors
		7000: "File does not exist",
		7001: "Error adding file",
		7002: "Uploaded file size exceeded within the subscription",
		7003: "Error deleting file",

		// Directory and directory entries errors
		8000: "Directory does not exist",
		8010: "Directory entry does not exist",
		8011: "Error adding directory entry",
		8012: "Error updating directory entry",
		8013: "Error deleting directory entry",

		9000: "Report does not exist",
	}

	if message, exists := errorMessages[code]; exists {
		return message
	}
	return ""
}

func ParseAPIError(resp *http.Response) error {
	var apiErr APIError

	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	return &apiErr
}
