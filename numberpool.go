package numberpool

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	perror "github.com/plivo/Commonlib/go/error"
)

// NumberpoolClient - Client that will be used to interface with the Numberpool service
type Numberpool interface {
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, perror.PlivoError)
	Get(ctx context.Context, id string) (*Resource, perror.PlivoError)
	GetAll(ctx context.Context, subAccount string, limit, offset int) (*NumberpoolClient, perror.PlivoError)
}

// New - Constructs a NumberPool client
func New(URL string) (Numberpool, perror.PlivoError) {
	if URL == "" {
		return nil, perror.ErrBadRequestInvalidParameter
	}
	_, err := url.Parse(URL)
	if err != nil {
		return nil, perror.ErrBadRequestInvalidParameter.SetDescription("")
	}
	return &NumberpoolClient{URL}, nil
}

type NumberpoolClient struct {
	url     string
	timeout int
}

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

// Create - Initiates a numberpool creation process, returns the ID of the pool that will be created
func (nc *NumberpoolClient) Create(req *CreateRequest) (*CreateResponse, perror.PlivoError) {

	data, err := req.Marshal()
	if err != nil {
		return nil, PoolCreateRequestMarshalError
	}

	resp, err := nc.sendRequest(nc.url, "POST", data)
	if err != nil {
		poolError := *PoolCreationError
		return nil, poolError.SetDescription(err.Error()).SetInternalData(err)
	}

	response := &CreateResponse{}
	if err = response.Load(resp); err != nil {
		poolErr := *PoolCreateResponseLoadError
		return nil, poolErr.SetDescription(err.Error()).SetInternalData(err)
	}

	return response, nil
}

// Get - Initiates a numberpool creation process, returns the ID of the pool that will be created
func (nc *NumberpoolClient) Create(req *CreateRequest) (*CreateResponse, perror.PlivoError) {

	data, err := req.Marshal()
	if err != nil {
		return nil, PoolCreateRequestMarshalError
	}

	resp, err := nc.sendRequest(nc.url, "POST", data)
	if err != nil {
		poolError := *PoolCreationError
		return nil, poolError.SetDescription(err.Error()).SetInternalData(err)
	}

	response := &CreateResponse{}
	if err = response.Load(resp); err != nil {
		poolErr := *PoolCreateResponseLoadError
		return nil, poolErr.SetDescription(err.Error()).SetInternalData(err)
	}

	return response, nil
}
