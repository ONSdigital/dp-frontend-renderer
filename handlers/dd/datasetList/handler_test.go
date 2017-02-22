package datasetList

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-frontend-models/model"
	"github.com/ONSdigital/dp-frontend-models/model/dd/datasetList"
	"github.com/ONSdigital/dp-frontend-renderer/render"
	"github.com/ONSdigital/dp-frontend-renderer/render/rendertest"
	zebedeeModel "github.com/ONSdigital/go-ns/zebedee/model"
	. "github.com/smartystreets/goconvey/convey"
)

type mockZebedeeClient struct{}

func (m *mockZebedeeClient) GetData(uri, requestID string) ([]byte, string, error) {
	return []byte(`{"serviceMessage": "Foo bar"}`), "", nil
}
func (m *mockZebedeeClient) GetTaxonomy(uri string, depth int, requestID string) ([]zebedeeModel.ContentNode, error) {
	return nil, nil
}
func (m *mockZebedeeClient) GetParents(uri, requestID string) ([]zebedeeModel.ContentNode, error) {
	return nil, nil
}

func TestHandler(t *testing.T) {

	f := &rendertest.FakeRenderer{}
	render.Renderer = f

	render.ZebedeeClient = &mockZebedeeClient{}

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
		So(f.Binding, ShouldHaveSameTypeAs, &datasetList.DatasetList{})
		p := f.Binding.(*datasetList.DatasetList)
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

	Convey("Dataset titles are rendered", t, func() {
		recorder := httptest.NewRecorder()
		rdr := bytes.NewReader([]byte(`{"datasets":{"items":[{"id":"ID1","title":"A Test Dataset"}]}}`))
		request, err := http.NewRequest("POST", "/", rdr)
		So(err, ShouldBeNil)
		Handler(recorder, request)
		So(recorder.Code, ShouldEqual, 200)
		So(f.Binding, ShouldHaveSameTypeAs, &datasetList.DatasetList{})
		p := f.Binding.(*datasetList.DatasetList)
		So(p.Datasets.Items, ShouldHaveLength, 1)
		So(p.Datasets.Items[0].Title, ShouldEqual, "A Test Dataset")
	})
}
