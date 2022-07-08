package bone

import "context"

// Model is a special component, which should be used for the model's
// definition.
type Model interface {
}

// Repo is a special component, which should be used for the repository's
// definition.
type Repo interface {
}

// Service is a special component, which should be used for the service's
// definition.
type Service interface {
}

// Endpoint is a special component, which should be used for the endpoint's
// definition.
type Endpoint interface {
}

// Transport is a special component, which should be used for the transport's
// definition.
type Transport interface {
	// Register should be used to registering endpoints.
	Register() error
}

// Module is a special component, which should be used for the module's
// definition.
type Module interface {
	Component
}

// Handler defines how to handle a http request, should have decoding part and
// proccessing part at the same time.
type Handler func(ctx context.Context, req any) (rsp any, err error)

// Controller is a special component, which should be used for the controller's
// definition.
type Controller interface {
	// Register should be used to registering handlers.
	Register() error
}
