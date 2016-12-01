package homepage

import (
	"testing"
	"net/http/httptest"
	"bytes"
	"net/http"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/dp-frontend-renderer/render/rendertest"
	"github.com/ONSdigital/dp-frontend-models/model/dd/homepage"
)

func TestHandler(t *testing.T) {

	f := &rendertest.FakeRenderer{}
	render.Renderer = f

	Convey("Handler returns 400 status code response when request body is empty", t, func() {
		recorder := httptest.NewRecorder()
		rdr := bytes.NewReader([]byte(``))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		request.Header.Set("Accept-Language", "en")
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 400)
	})

	Convey("Handler returns 500 status code when HTML render returns an error", t, func() {
		f.ErrorOnHTML = true
		recorder := httptest.NewRecorder()
		rdr := bytes.NewReader([]byte(`{}`))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		request.Header.Set("Accept-Language", "cy")
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 500)
		So(f.Binding, ShouldHaveSameTypeAs, model.ErrorResponse{})
		p := f.Binding.(model.ErrorResponse)
		So(p.Error, ShouldEqual, "Error from HTML")
		f.ErrorOnHTML = false
	})

	Convey("Handler returns language of English when neither 'en' or 'cy' is set from request header", t, func() {
		recorder := httptest.NewRecorder()
		rdr := bytes.NewReader([]byte(`{}`))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		request.Header.Set("Accept-Language", "foo")
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 200)
		So(f.Binding, ShouldHaveSameTypeAs, &homepage.Homepage{})
		p := f.Binding.(*homepage.Homepage)
		So(p.Language, ShouldEqual, "en")
	})

	Convey("Handler returns 400 status code when io reader has error", t, func() {
		recorder := httptest.NewRecorder()
		rdr := &rendertest.FakeReader{}
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 400)
		So(f.Binding, ShouldHaveSameTypeAs, model.ErrorResponse{})
		p := f.Binding.(model.ErrorResponse)
		So(p.Error, ShouldEqual, "Error from reader")
	})
}
