package dtos

type MetaData struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Meta MetaData    `json:"meta,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
