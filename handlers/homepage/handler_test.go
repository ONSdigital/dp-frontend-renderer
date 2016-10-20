package homepage

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	. "github.com/smartystreets/goconvey/convey"
	unrolled "github.com/unrolled/render"
)

type fakeRenderer struct {
	binding     interface{}
	errorOnHTML bool
}

func (f *fakeRenderer) HTML(w io.Writer, status int, name string, binding interface{}, htmlOpt ...unrolled.HTMLOptions) error {
	if f.errorOnHTML {
		return errors.New("Error from HTML")
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

func TestHandler(t *testing.T) {
	f := &fakeRenderer{}
	render.Renderer = f

	config.PatternLibraryAssetsPath = "foobar.com"

	Convey("Handler returns 400 status code response when request body is empty", t, func() {
		recorder := httptest.NewRecorder()
		rdr := bytes.NewReader([]byte(``))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		request.Header.Set("Accept-Language", "en")
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 400)
	})

	Convey("Handler returns matching data from request page model", t, func() {
		recorder := httptest.NewRecorder()
		rdr := bytes.NewReader([]byte(`{"serviceMessage": "Foo bar"}`))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		request.Header.Set("Accept-Language", "en")
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 200)
		So(f.binding, ShouldHaveSameTypeAs, Page{})
		p := f.binding.(Page)
		So(p.ServiceMessage, ShouldEqual, "Foo bar")
		So(p.PatternLibraryAssetsPath, ShouldEqual, config.PatternLibraryAssetsPath)
	})

	Convey("Handler returns 500 status code when HTML render returns an error", t, func() {
		f.errorOnHTML = true
		recorder := httptest.NewRecorder()
		rdr := bytes.NewReader([]byte(`{}`))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		request.Header.Set("Accept-Language", "cy")
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 500)
		So(f.binding, ShouldHaveSameTypeAs, ErrorResponse{})
		p := f.binding.(ErrorResponse)
		So(p.Error, ShouldEqual, "Error from HTML")
	})

	// Convey("Handler returns language of English when neither 'en' or 'cy' is set from request header", t, func() {
	// 	fmt.Println("ErrorOnHTML equals ", f.errorOnHTML)
	// 	recorder := httptest.NewRecorder()
	// 	rdr := bytes.NewReader([]byte(`{}`))
	// 	request, err := http.NewRequest("POST", "/", rdr)
	// 	So(err, ShouldBeNil)
	// 	request.Header.Set("Accept-Language", "foo")
	// 	Handler(recorder, request)
	// 	So(recorder.Code, ShouldEqual, 200)
	// 	So(f.binding, ShouldHaveSameTypeAs, ErrorResponse{})
	// 	p := f.binding.(ErrorResponse)
	// 	So(p.Error, ShouldEqual, "Error from HTML")
	// })
}
