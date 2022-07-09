package bone

import (
	"testing"
)

type serviceForComponentTest struct{}

var _ Service = (*serviceForComponentTest)(nil)

type moduleForComponentTest struct {
	Service *serviceForComponentTest `inject:""`
}

func (*moduleForComponentTest) Name() string {
	return "module.test"
}

func (*moduleForComponentTest) Init() error {
	return nil
}

func (*moduleForComponentTest) Register() error {
	return nil
}

func (*moduleForComponentTest) Unregister() error {
	return nil
}

var _ Module = (*moduleForComponentTest)(nil)

func TestUseComponent(t *testing.T) {
	options := DefaultApplicationOptions()
	application := NewApplication(options)
	m := new(moduleForComponentTest)
	application.Use(m)
	if m.Service == nil {
		t.Fail()
	}
}
