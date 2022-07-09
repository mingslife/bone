package bone

import (
	"net/http"
	"testing"
)

func TestRunApplication(t *testing.T) {
	options := DefaultApplicationOptions()
	application := NewApplication(options)
	go application.Run()

	_, err := http.Get("http://127.0.0.1:8080/")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
