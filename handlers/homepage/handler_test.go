package homepage

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-frontend-models/model"
	homepage "github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/go-ns/render"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	unrolled "github.com/unrolled/render"
)

type fakeRenderer struct {
	binding     interface{}
	errorOnHTML bool
}

func (f *fakeRenderer) HTML(w io.Writer, status int, name string, binding interface{}, htmlOpt ...unrolled.HTMLOptions) error {
	if f.errorOnHTML {
		return errors.New("error from HTML")
	}
	w.(http.ResponseWriter).WriteHeader(status)
	f.binding = binding
	return nil
}
func (f *fakeRenderer) JSON(w io.Writer, status int, v interface{}) error {
	w.(http.ResponseWriter).WriteHeader(status)
	f.binding = v
	return nil
}

type fakeReader struct {
}

func (*fakeReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("error from reader")
}

// doTestRequest helper function that creates a router and mocks requests
func doTestRequest(target string, req *http.Request, handlerFunc http.HandlerFunc, w *httptest.ResponseRecorder) *httptest.ResponseRecorder {
	if w == nil {
		w = httptest.NewRecorder()
	}
	router := mux.NewRouter()
	router.HandleFunc(target, handlerFunc)
	router.ServeHTTP(w, req)
	return w
}

func TestHandler(t *testing.T) {
	f := &fakeRenderer{}
	render.Renderer = f
	cfg := config.Config{
		BindAddr:                 ":20010",
		Debug:                    false,
		SiteDomain:               "ons.gov.uk",
		SupportedLanguages:       [2]string{"en", "cy"},
		PatternLibraryAssetsPath: "foobar.com",
	}

	Convey("Handler returns 400 status code response when request body is empty", t, func() {
		rdr := bytes.NewReader([]byte(``))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		w := doTestRequest("/", request, Handler(cfg), nil)
		So(w.Code, ShouldEqual, 400)
	})

	Convey("Handler returns matching data from request page model", t, func() {
		rdr := bytes.NewReader([]byte(`{"service_message": "Foo bar"}`))
		request := httptest.NewRequest("POST", "/", rdr)
		w := doTestRequest("/", request, Handler(cfg), nil)
		So(w.Code, ShouldEqual, 200)
		So(f.binding, ShouldHaveSameTypeAs, &homepage.Page{})
		p := f.binding.(*homepage.Page)
		So(p.ServiceMessage, ShouldEqual, "Foo bar")
		So(p.PatternLibraryAssetsPath, ShouldEqual, cfg.PatternLibraryAssetsPath)
	})

	Convey("Handler returns 500 status code when HTML render returns an error", t, func() {
		f.errorOnHTML = true
		rdr := bytes.NewReader([]byte(`{}`))
		request := httptest.NewRequest("POST", "/", rdr)
		request.Header.Set("Accept-Language", "cy")
		w := doTestRequest("/", request, Handler(cfg), nil)
		So(w.Code, ShouldEqual, 500)
		So(f.binding, ShouldHaveSameTypeAs, model.ErrorResponse{})
		p := f.binding.(model.ErrorResponse)
		So(p.Error, ShouldEqual, "error from HTML")
		f.errorOnHTML = false
	})

	Convey("Handler returns 400 status code when io reader has error", t, func() {
		rdr := &fakeReader{}
		request := httptest.NewRequest("POST", "/", rdr)
		w := doTestRequest("/", request, Handler(cfg), nil)
		So(w.Code, ShouldEqual, 400)
		So(f.binding, ShouldHaveSameTypeAs, model.ErrorResponse{})
		p := f.binding.(model.ErrorResponse)
		So(p.Error, ShouldEqual, "error from reader")
	})
}
