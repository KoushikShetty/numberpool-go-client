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

func (cr *Resource) Load(data *http.Response) perror.PlivoError {
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

type ListResource struct {
	Pools []Resource `json:"pools"`
	Meta  `json:"meta"`
}

type Meta struct {
	Limit    string `json:"limit"`
	offset   string `json:"offset"`
	Total    int    `json:"total_count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

func (cr *ListResource) Load(data *http.Response) perror.PlivoError {
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
