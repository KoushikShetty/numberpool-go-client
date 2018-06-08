package numberpool

import (
	"encoding/json"
	"net/http"

	perror "github.com/plivo/Commonlib/go/error"
)

// Resource - Numberpool resource
type Resource struct {
	ID                     string        `json:"id"`
	Status                 string        `json:"status"`
	Name                   string        `json:"name"`
	SubAccount             string        `json:"sub_account"`
	TotalCount             string        `json:"total_number_count"`
	RequestURL             string        `json:"resource_url"`
	PhoneNumberResourceURL string        `json:"phone_numbers_resource_uri"`
	Composition            []Composition `json:"composition"`
}

// Composition represents how the pool is composed of based on country, area, type etc..
type Composition struct {
	Count    int       `json:"number_count"`
	Criteria *Criteria `json:"criteria"`
}

// NewComposition creates a new composition object
func NewComposition(count int, criteria *Criteria) *Composition {
	return &Composition{count, criteria}
}

// Criteria - criteria object
type Criteria struct {
	CountryISO string   `json:"country_iso"`
	Type       string   `json:"type"`
	Pattern    []string `json:"pattern"`
}

// NewCiteria - ctor
func NewCiteria(countryISO, numType string, pattern []string) *Criteria {
	return &Criteria{
		CountryISO: countryISO,
		Type:       numType,
		Pattern:    pattern,
	}
}

func (cr *Resource) load(data *http.Response) perror.PlivoError {
	defer data.Body.Close()
	switch data.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(data.Body).Decode(cr); err != nil {
			poolErr := *PoolResourceLoadError
			return poolErr.SetDescription(err.Error()).SetInternalData(err)
		}
	default:
		res := *PoolResourceLoadError
		return res.SetDescription("Invalid status code from service")

	}
	return nil
}

// ListResource - Holds the list of numberpool resources
type ListResource struct {
	Pools []Resource `json:"pools"`
	Meta  meta       `json:"meta"`
}

// meta - represents the meta information for the lsit resource
type meta struct {
	Limit    string `json:"limit"`
	Offset   string `json:"offset"`
	Total    int    `json:"total_count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// Load - loads the list resource
func (cr *ListResource) load(data *http.Response) perror.PlivoError {
	defer data.Body.Close()
	switch data.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(data.Body).Decode(cr); err != nil {
			poolErr := *PoolListResourceLoadError
			return poolErr.SetDescription(err.Error()).SetInternalData(err)
		}
	default:
		res := *PoolListResourceLoadError
		return res.SetDescription("Invalid status code from service")

	}
	return nil
}
