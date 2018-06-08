package numberpool

import (
	perror "github.com/plivo/Commonlib/go/error"
)

var (
	PoolCreateRequestMarshalError = perror.New(1, "Error while marshalling request", nil)
	PoolCreationError             = perror.New(2, "Error creating pool", nil)
	PoolCreateResponseError       = perror.New(3, "Error response from service", nil)
	PoolCreateResponseLoadError   = perror.New(4, "Error while loading response", nil)
	PoolResourceFetchError        = perror.New(5, "Error while fetching numberpool resource", nil)
	PoolResourceLoadError         = perror.New(6, "Error while loading numberpool resource", nil)
	PoolListResourceFetchError    = perror.New(7, "Error while fetching numberpool list resource", nil)
	PoolListResourceLoadError     = perror.New(8, "Error while loading numberpool listresource", nil)
	PoolDeleteError               = perror.New(9, "Error while deleting numberpool resource", nil)
)
