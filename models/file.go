package models

type File struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Size        int64  `json:"size,omitempty"`
	DownloadURL string `json:"downloadUrl,omitempty"`
}

type FileRequest struct {
	ID int `json:"id"`
}

type FileGetResponse struct {
	Result string `json:"result"`
	File   File   `json:"file"`
}

type FileUploadRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type FileUploadResponse struct {
	Result string `json:"result"`
	ID     int    `json:"id"`
}

type FileUpdateRequest struct {
	Name string `json:"name"`
}

type FileUpdateResponse struct {
	Result string `json:"result"`
}
