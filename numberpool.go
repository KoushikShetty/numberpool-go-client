package numberpool

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	perror "github.com/plivo/Commonlib/go/error"
)

// Paths for the endpoints
const (
	PathPoolCreation = "%s/number_pool"
	PathPoolUpdate   = "%s/number_pool/%s"
	PathPoolDelete   = "%s/number_pool/%s"
	PathGetPool      = "%s/number_pool/%s"
	PathGetAllPool   = "%s/number_pool"
)

// NumberpoolClient - Client that will be used to interface with the Numberpool service
type Numberpool interface {
	Create(req *CreateRequest) (*CreateResponse, perror.PlivoError)
	Get(id string) (*Resource, perror.PlivoError)
	GetAll(subAccount string, limit, offset int) (*ListResource, perror.PlivoError)
	Delete(id string) perror.PlivoError
}

// NumberpoolClient - Client for the numberpool resource
type NumberpoolClient struct {
	host    string
	timeout int
	authID  string
	baseURL string
}

// New - Constructs a NumberPool client
func New(authID, host string, timeout int) (Numberpool, perror.PlivoError) {
	if host == "" {
		return nil, perror.ErrBadRequestInvalidParameter
	}

	if authID == "" {
		return nil, perror.ErrBadRequestInvalidParameter
	}

	// defaulting to 1000ms
	if timeout < 1 {
		timeout = 1000
	}

	_, err := url.Parse(host)
	if err != nil {
		return nil, perror.ErrBadRequestInvalidParameter.SetDescription("")
	}
	return &NumberpoolClient{
		host:    host,
		timeout: timeout,
		authID:  authID,
		baseURL: fmt.Sprintf("%s/v1/account/%s", host, authID),
	}, nil
}

// Create - Initiates a numberpool creation process, returns the ID of the pool that will be created
func (nc *NumberpoolClient) Create(req *CreateRequest) (*CreateResponse, perror.PlivoError) {
	data, err := req.Marshal()
	if err != nil {
		return nil, PoolCreateRequestMarshalError
	}

	path := fmt.Sprintf(PathPoolCreation, nc.baseURL)
	resp, err := nc.sendRequest(path, "POST", data)
	if err != nil {
		poolError := *PoolCreationError
		return nil, poolError.SetDescription(err.Error()).SetInternalData(err)
	}

	response := &CreateResponse{}
	return response, response.load(resp)
}

// Get - fetch the numberpool associated with the specified id
func (nc *NumberpoolClient) Get(id string) (*Resource, perror.PlivoError) {
	if id == "" {
		return nil, perror.ErrBadRequestInvalidParameter
	}

	path := fmt.Sprintf(PathGetPool, nc.host, id)
	resp, err := nc.sendRequest(path, "GET", nil)
	if err != nil {
		poolError := *PoolResourceFetchError
		return nil, poolError.SetDescription(err.Error()).SetInternalData(err)
	}

	response := &Resource{}
	return response, response.load(resp)
}

// GetAll - fetchs all the numberpool associated with the account_id
func (nc *NumberpoolClient) GetAll(subAccount string, limit, offset int) (*ListResource, perror.PlivoError) {
	path := fmt.Sprintf(PathGetAllPool, nc.host)
	resp, err := nc.sendRequest(path, "GET", nil)
	if err != nil {
		poolError := *PoolListResourceFetchError
		return nil, poolError.SetDescription(err.Error()).SetInternalData(err)
	}

	response := &ListResource{}
	return response, response.load(resp)
}

// GetAll - fetchs all the numberpool associated with the account_id
func (nc *NumberpoolClient) Delete(id string) perror.PlivoError {
	if id == "" {
		return perror.ErrBadRequestInvalidParameter
	}

	path := fmt.Sprintf(PathPoolDelete, nc.host, id)
	resp, err := nc.sendRequest(path, "DELETE", nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		poolError := *PoolDeleteError
		return poolError.SetDescription(err.Error()).SetInternalData(err)
	}

	return nil
}

// Internal function to send the request to the client
func (nc *NumberpoolClient) sendRequest(url, method string, data io.Reader) (*http.Response, error) {
	// Create the context for the request
	c, cancelFn := context.WithTimeout(context.Background(), time.Duration(nc.timeout)*time.Millisecond)
	defer cancelFn()

	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	req.WithContext(c)
	client := http.Client{}
	return client.Do(req)
}
