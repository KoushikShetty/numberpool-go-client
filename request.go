package numberpool

import (
	"bytes"
	"encoding/json"
	"io"
)

// CreateRequest represents the Request object that the numberpool works with
type CreateRequest struct {
	Name          string         `json:"name"`
	ApplicationID string         `json:"application_id"`
	Callback      *Callback      `json:"callback"`
	Composition   []*Composition `json:"composition"`
}

// Marshal - serialzes the object
func (c *CreateRequest) Marshal() (io.Reader, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

// NewCreateRequest - creates the request for numberpool creation
func NewCreateRequest(name, appID string, cb *Callback, composition []*Composition) *CreateRequest {
	return &CreateRequest{
		Name:          name,
		ApplicationID: appID,
		Callback:      cb,
		Composition:   composition,
	}
}

// Callback is the uri that is used to send the process completion information.
type Callback struct {
	URL    string `json:"url"`
	Method string `json:"method"`
}

// NewCallback - ctor
func NewCallback(url, method string) *Callback {
	return &Callback{
		url, method,
	}
}
