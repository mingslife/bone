package bone

import (
	"context"
	"encoding/json"
	"net/http"
)

var Http = new(httputil)

type httputil struct{}

func (*httputil) EncodeJSON(ctx context.Context, w http.ResponseWriter, rsp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(rsp)
}
