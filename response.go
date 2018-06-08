package numberpool

import (
	"encoding/json"
	"net/http"

	perror "github.com/plivo/Commonlib/go/error"
)

// CreateResponse is the respose to the
type CreateResponse struct {
	ID,
	Status string
}

func (cr *CreateResponse) Load(data *http.Response) perror.PlivoError {
	defer data.Body.Close()
	switch data.StatusCode {
	case http.StatusProcessing:
		fallthrough
	case http.StatusOK:
		if err := json.NewDecoder(data.Body).Decode(cr); err != nil {
			poolErr := *PoolCreateResponseLoadError
			return poolErr.SetDescription(err.Error()).SetInternalData(err)
		}
	default:
		res := *PoolCreateResponseError
		return res.SetDescription("Invalid status code from service")

	}
	return nil
}
