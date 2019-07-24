package render

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

var testMarkdown = []string{
	"- First bullet point\n\n- Second bullet point\n\n\tSecond line of second bullet point",
	"1. First bullet point\n\n2. Second bullet point",
	"##Heading2",
	"### Heading 3",
	"#### Heading 4 \n\n Some text, _some italic text_, **some strong text**",
	"[Test link](https://www.test.com)",
	"### Heading 3 \n\n some text [Anchor link](#)",
}

// the new line at the end of each test case is needed
var testHTMLForMarkdown = []string{
	"<ul>\n<li><p>First bullet point</p></li>\n\n<li><p>Second bullet point</p>\n\n<p>Second line of second bullet point</p></li>\n</ul>\n",
	"<ol>\n<li><p>First bullet point</p></li>\n\n<li><p>Second bullet point</p></li>\n</ol>\n",
	"<h2>Heading2</h2>\n",
	"<h3>Heading 3</h3>\n",
	"<h4>Heading 4</h4>\n\n<p>Some text, <em>some italic text</em>, <strong>some strong text</strong></p>\n",
	"<p><a href=\"https://www.test.com\">Test link</a></p>\n",
	"<h3>Heading 3</h3>\n\n<p>some text <a href=\"#\">Anchor link</a></p>\n",
}

func TestMarkdown(t *testing.T) {
	Convey("markdown should return expected HTML", t, func() {
		for i := 0; i < len(testMarkdown); i++ {
			So(Markdown(testMarkdown[i]), ShouldEqual, testHTMLForMarkdown[i])
		}
	})
}

// TestLocalise ensures that the correct strings are returned given a range of different arguments
func TestLocalise(t *testing.T) {
	Convey("English singular is returned", t, func() {
		So(Localise("Foo", "en", 1), ShouldEqual, "One Foo (English)")
	})
	Convey("English plural is returned for more than two", t, func() {
		So(Localise("Foo", "en", 4), ShouldEqual, "More than two Foo (English)")
	})
	Convey("Welsh singular is returned", t, func() {
		So(Localise("Foo", "cy", 1), ShouldEqual, "One Foo (Welsh)")
	})
	Convey("Welsh plural for more than two is returned", t, func() {
		So(Localise("Foo", "cy", 4), ShouldEqual, "More than two Foo (Welsh)")
	})
}
