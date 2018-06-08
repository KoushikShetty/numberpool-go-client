package numberpool

import (
	"bytes"
	"encoding/json"
	"io"
)

// CreateRequest represents the Request object that the numberpool works with
type CreateRequest struct {
	Name          string         `json:"name"`
	ApplicationID *string        `json:"application_id"`
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

// Callback is the uri that is used to send the process completion information.
type Callback struct {
	URL    string `json:"url"`
	Method string `json:"method"`
}

func NewCallback(url, method string) *Callback {
	return &Callback{
		url, method,
	}
}

// Composition represents how the pool is composed of based on country, area, type etc..
type Composition struct {
	Count    int      `json:"number_count"`
	Criteria Criteria `json:"criteria"`
}

// NewComposition creates a new composition object
func NewComposition(count int, criteria Criteria) *Composition {
	return &Composition{count, criteria}
}

type Criteria struct {
	CountryISO string   `json:"country_iso"`
	Type       string   `json:"type"`
	Pattern    []string `json:"pattern"`
}
