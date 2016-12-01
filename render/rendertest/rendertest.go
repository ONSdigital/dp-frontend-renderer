package rendertest

import (
	"errors"
	"io"
	"net/http"

	unrolled "github.com/unrolled/render"
)

type FakeRenderer struct {
	Binding     interface{}
	ErrorOnHTML bool
}

func (f *FakeRenderer) HTML(w io.Writer, status int, name string, binding interface{}, htmlOpt ...unrolled.HTMLOptions) error {
	if f.ErrorOnHTML {
		return errors.New("Error from HTML")
	}
	w.(http.ResponseWriter).WriteHeader(status)
	f.Binding = binding
	return nil
}
func (f *FakeRenderer) JSON(w io.Writer, status int, v interface{}) error {
	w.(http.ResponseWriter).WriteHeader(status)
	f.Binding = v
	return nil
}

type FakeReader struct {
}

func (*FakeReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("Error from reader")
}
