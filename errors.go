package numberpool

import (
	perror "github.com/plivo/Commonlib/go/error"
)

var (
	PoolCreateRequestMarshalError = perror.New(0, "Error while marshalling request", nil)
	PoolCreationError             = p.New(1, "Error creating pool", nil)
	PoolCreateResponseLoadError   = perror.New(2, "Error while loading response", nil)
)
