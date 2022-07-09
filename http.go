package bone

import (
	"github.com/go-kit/kit/transport/http"
)

var (
	// NewServer is shortcut of go-kit's http.NewServer.
	// NewServer constructs a new server, which implements http.Handler and wraps
	// the provided endpoint.
	NewServer = http.NewServer
	// EncodeJSONResquest is shortcut of go-kit's http.EncodeJSONRequest.
	// EncodeJSONRequest is an EncodeRequestFunc that serializes the request as a
	// JSON object to the Request body. Many JSON-over-HTTP services can use it as
	// a sensible default. If the request implements Headerer, the provided headers
	// will be applied to the request.
	EncodeJSONRequest = http.EncodeJSONRequest
	// EncodeXMLRequest is shortcut of go-kit's http.EncodeXMLRequest.
	// EncodeXMLRequest is an EncodeRequestFunc that serializes the request as a
	// XML object to the Request body. If the request implements Headerer,
	// the provided headers will be applied to the request.
	EncodeXMLRequest = http.EncodeXMLRequest
	// EncodeJSONResponse is shortcut of go-kit's http.EncodeJSONResponse.
	// EncodeJSONResponse is a EncodeResponseFunc that serializes the response as a
	// JSON object to the ResponseWriter. Many JSON-over-HTTP services can use it as
	// a sensible default. If the response implements Headerer, the provided headers
	// will be applied to the response. If the response implements StatusCoder, the
	// provided StatusCode will be used instead of 200.
	EncodeJSONResponse = http.EncodeJSONResponse
)
