package bone

import (
	"context"
	"net/http"
	"testing"
)

type endpointForRouterTest struct{}

func (*endpointForRouterTest) Test(ctx context.Context, req any) (rsp any, err error) {
	return struct{}{}, nil
}

var _ Endpoint = (*endpointForRouterTest)(nil)

type transportForRouterTest struct {
	Router   *Router                `inject:"application.router"`
	Endpoint *endpointForRouterTest `inject:""`
}

func (t *transportForRouterTest) Register() error {
	t.Router.Methods(http.MethodGet).Path("/").Handler(NewServer(t.Endpoint.Test, nil, EncodeJSONResponse))
	return nil
}

var _ Transport = (*transportForRouterTest)(nil)

type moduleForRouterTest struct {
	Transport *transportForRouterTest `inject:""`
}

func (*moduleForRouterTest) Name() string {
	return "module.test"
}

func (*moduleForRouterTest) Init() error {
	return nil
}

func (*moduleForRouterTest) Register() error {
	return nil
}

func (*moduleForRouterTest) Unregister() error {
	return nil
}

var _ Module = (*moduleForRouterTest)(nil)

func TestRegisterRouter(t *testing.T) {
	options := DefaultApplicationOptions()
	application := NewApplication(options)
	application.Use(new(moduleForRouterTest))
	go application.Run()

	rsp, err := http.Get("http://127.0.0.1:8080/")
	if err != nil || rsp.StatusCode == http.StatusOK {
		t.Log(err)
		t.Fail()
	}
}
