package render

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLegacyDatasetDownloadURI(t *testing.T) {
	Convey("should generated expected legacy dataset download URI", t, func() {
		So(LegacyDataSetDownloadURI("/legacy/dataset/page", "test.csv"), ShouldEqual, "/file?uri=/legacy/dataset/page/test.csv")
	})
}

func TestSubtract(t *testing.T) {
	Convey("substract should return expected value", t, func() {
		So(Subtract(100, 1), ShouldEqual, 99)
	})
}

func TestHumanSize(t *testing.T) {
	Convey("humanSize should return the expected value for 1500 bytes", t, func() {
		res, err := HumanSize("1500")
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "1.5 KB")
	})

	Convey("humanSize should return the expected value for empty input", t, func() {
		res, err := HumanSize("")
		So(err, ShouldBeNil)
		So(res, ShouldBeEmpty)
	})

	Convey("humanSize should return error for non numeric input", t, func() {
		res, err := HumanSize("green eggs and ham")
		So(err, ShouldNotBeNil)
		So(res, ShouldBeEmpty)
	})
}
