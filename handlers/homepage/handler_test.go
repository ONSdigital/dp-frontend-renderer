package homepage

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-models/model/homepage"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/go-ns/log"
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

type fakeReader struct {
}

func (*fakeReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("Error from reader")
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
		rdr := bytes.NewReader([]byte(`{"service_message": "Foo bar"}`))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		request.Header.Set("Accept-Language", "en")
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 200)
		So(f.binding, ShouldHaveSameTypeAs, &homepage.Page{})
		p := f.binding.(*homepage.Page)
		log.Debug("page", log.Data{"p": p})
		So(p.ServiceMessage, ShouldEqual, "Foo bar")
		So(p.PatternLibraryAssetsPath, ShouldEqual, config.PatternLibraryAssetsPath)
	})

	Convey("SparklineData dates are copied to HeadlineFigure", t, func() {
		recorder := httptest.NewRecorder()
		rdr := bytes.NewReader([]byte(`{"data": {"headlineFigures": [{"sparklineData": [{"name": "foo"}, {"name": "bar"}, {"name": "baz"}]}]}}`))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		request.Header.Set("Accept-Language", "en")
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 200)
		So(f.binding, ShouldHaveSameTypeAs, &homepage.Page{})
		p := f.binding.(*homepage.Page)
		So(p.Data.HeadlineFigures[0].StartDate, ShouldEqual, "foo")
		So(p.Data.HeadlineFigures[0].EndDate, ShouldEqual, "baz")
	})

	Convey("SparklineData dates are skipped if sparklineData is empty", t, func() {
		recorder := httptest.NewRecorder()
		rdr := bytes.NewReader([]byte(`{"data": {"headlineFigures": [{"sparklineData": []}]}}`))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		request.Header.Set("Accept-Language", "en")
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 200)
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
		So(f.binding, ShouldHaveSameTypeAs, model.ErrorResponse{})
		p := f.binding.(model.ErrorResponse)
		So(p.Error, ShouldEqual, "Error from HTML")
		f.errorOnHTML = false
	})

	Convey("Handler returns language of English when neither 'en' or 'cy' is set from request header", t, func() {
		recorder := httptest.NewRecorder()
		rdr := bytes.NewReader([]byte(`{}`))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		request.Header.Set("Accept-Language", "foo")
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 200)
		So(f.binding, ShouldHaveSameTypeAs, &homepage.Page{})
		p := f.binding.(*homepage.Page)
		So(p.Language, ShouldEqual, "en")
	})

	Convey("Handler returns 400 status code when io reader has error", t, func() {
		recorder := httptest.NewRecorder()
		rdr := &fakeReader{}
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 400)
		So(f.binding, ShouldHaveSameTypeAs, model.ErrorResponse{})
		p := f.binding.(model.ErrorResponse)
		So(p.Error, ShouldEqual, "Error from reader")
	})
}
