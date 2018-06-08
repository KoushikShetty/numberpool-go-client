package numberpool

import (
	"encoding/json"
	"net/http"
)

// CreateResponse is the respose to the
type CreateResponse struct {
	ID,
	Status string
}

func (cr *CreateResponse) Load(data *http.Response) error {
	defer data.Body.Close()
	return json.NewDecoder(data.Body).Decode(cr)
}
