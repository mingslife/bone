package bone

import (
	"github.com/go-kit/kit/transport/http"
)

var (
	NewServer         = http.NewServer
	EncodeJSONRequest = http.EncodeJSONRequest
	EncodeXMLRequest  = http.EncodeXMLRequest
)
