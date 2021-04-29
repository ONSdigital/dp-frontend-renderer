package render

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLegacyDatasetDownloadURI(t *testing.T) {
	Convey("should generated expected legacy dataset download URI", t, func() {
		So(LegacyDataSetDownloadURI("/legacy/dataset/page", "test.csv"), ShouldEqual, "/file?uri=/legacy/dataset/page/test.csv")
	})
	Convey("no extra parameter given should still generate uri", t, func() {
		So(LegacyDataSetDownloadURI("/legacy/dataset/page"), ShouldEqual, "/file?uri=/legacy/dataset/page")
	})
}

func TestAdd(t *testing.T) {
	Convey("add should return expected value", t, func() {
		So(Add(99, 1), ShouldEqual, 100)
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
	Convey("English plural is returned for more than one", t, func() {
		So(Localise("Foo", "en", 4), ShouldEqual, "Two or more Foos (English)")
	})
	Convey("Welsh nil value sentence returned", t, func() {
		So(Localise("Foo", "cy", 0), ShouldEqual, "No Foos (Welsh)")
	})
	Convey("Welsh singular is returned", t, func() {
		So(Localise("Foo", "cy", 1), ShouldEqual, "One Foo (Welsh)")
	})
	Convey("Welsh plural for two is returned", t, func() {
		So(Localise("Foo", "cy", 2), ShouldEqual, "Two Foos (Welsh)")
	})
	Convey("Welsh plural for few(3) is returned", t, func() {
		So(Localise("Foo", "cy", 3), ShouldEqual, "Three Foos (Welsh)")
	})
	Convey("Welsh plural for other (4) is returned", t, func() {
		So(Localise("Foo", "cy", 4), ShouldEqual, "Four, five or more than six but not six Foos (Welsh)")
	})
	Convey("Welsh plural for other (5) is returned", t, func() {
		So(Localise("Foo", "cy", 5), ShouldEqual, "Four, five or more than six but not six Foos (Welsh)")
	})
	Convey("Welsh plural for many (6) is returned", t, func() {
		So(Localise("Foo", "cy", 6), ShouldEqual, "Six Foos (Welsh)")
	})
	Convey("Welsh plural for many (7) is returned", t, func() {
		So(Localise("Foo", "cy", 7), ShouldEqual, "Four, five or more than six but not six Foos (Welsh)")
	})
}

//TestDomainSetLang ensures the returned URL is set accurately
func TestDomainSetLang(t *testing.T) {
	Convey("English domain requested", t, func() {
		So(DomainSetLang("www.ons.gov.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("ons.gov.uk", "", "en"), ShouldEqual, "https://www.ons.gov.uk")
		So(DomainSetLang("https://www.ons.gov.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("https://ons.gov.uk", "", "en"), ShouldEqual, "https://www.ons.gov.uk")
		So(DomainSetLang("www.cy.ons.gov.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("cy.ons.gov.uk", "", "en"), ShouldEqual, "https://www.ons.gov.uk")
		So(DomainSetLang("https://www.cy.ons.gov.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("https://cy.ons.gov.uk", "", "en"), ShouldEqual, "https://www.ons.gov.uk")
		So(DomainSetLang("www.foo-bar.baz.co.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.foo-bar.baz.co.uk/foo/bar/baz")
		So(DomainSetLang("cy.foo-bar.baz.co.uk", "", "en"), ShouldEqual, "https://www.foo-bar.baz.co.uk")
		So(DomainSetLang("https://www.cy.foo-bar.baz.co.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.foo-bar.baz.co.uk/foo/bar/baz")
		So(DomainSetLang("https://cy.foo-bar.baz.co.uk", "", "en"), ShouldEqual, "https://www.foo-bar.baz.co.uk")
		So(DomainSetLang("https://cy.foo-bar.baz.co.uk", "http://foo:12345/bar/baz/qux", "en"), ShouldEqual, "https://www.foo-bar.baz.co.uk/bar/baz/qux")
	})

	Convey("Welsh domain requested", t, func() {
		So(DomainSetLang("www.ons.gov.uk", "", "cy"), ShouldEqual, "https://cy.ons.gov.uk")
		So(DomainSetLang("ons.gov.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("https://www.ons.gov.uk", "", "cy"), ShouldEqual, "https://cy.ons.gov.uk")
		So(DomainSetLang("https://ons.gov.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("www.cy.ons.gov.uk", "", "cy"), ShouldEqual, "https://cy.ons.gov.uk")
		So(DomainSetLang("cy.ons.gov.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("https://www.cy.ons.gov.uk", "", "cy"), ShouldEqual, "https://cy.ons.gov.uk")
		So(DomainSetLang("https://cy.ons.gov.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("www.foo-bar.baz.co.uk", "", "cy"), ShouldEqual, "https://cy.foo-bar.baz.co.uk")
		So(DomainSetLang("cy.foo-bar.baz.co.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.foo-bar.baz.co.uk/foo/bar/baz")
		So(DomainSetLang("https://www.cy.foo-bar.baz.co.uk", "", "cy"), ShouldEqual, "https://cy.foo-bar.baz.co.uk")
		So(DomainSetLang("https://cy.foo-bar.baz.co.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.foo-bar.baz.co.uk/foo/bar/baz")
		So(DomainSetLang("https://cy.foo-bar.baz.co.uk", "http://foo:12345/bar/baz/qux", "cy"), ShouldEqual, "https://cy.foo-bar.baz.co.uk/bar/baz/qux")
	})

	Convey("Unsupported domain requested", t, func() {
		So(DomainSetLang("www.ons.gov.uk", "", "foo"), ShouldEqual, "https://www.ons.gov.uk")
		So(DomainSetLang("ons.gov.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("https://www.ons.gov.uk", "", "foo"), ShouldEqual, "https://www.ons.gov.uk")
		So(DomainSetLang("https://ons.gov.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("www.cy.ons.gov.uk", "", "foo"), ShouldEqual, "https://www.ons.gov.uk")
		So(DomainSetLang("cy.ons.gov.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("https://www.cy.ons.gov.uk", "", "foo"), ShouldEqual, "https://www.ons.gov.uk")
		So(DomainSetLang("https://cy.ons.gov.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(DomainSetLang("www.foo-bar.baz.co.uk", "", "foo"), ShouldEqual, "https://www.foo-bar.baz.co.uk")
		So(DomainSetLang("cy.foo-bar.baz.co.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.foo-bar.baz.co.uk/foo/bar/baz")
		So(DomainSetLang("https://www.cy.foo-bar.baz.co.uk", "", "foo"), ShouldEqual, "https://www.foo-bar.baz.co.uk")
		So(DomainSetLang("https://cy.foo-bar.baz.co.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.foo-bar.baz.co.uk/foo/bar/baz")
		So(DomainSetLang("https://cy.foo-bar.baz.co.uk", "http://foo:12345/bar/baz/qux", "foo"), ShouldEqual, "https://www.foo-bar.baz.co.uk/bar/baz/qux")
	})
}

func TestHasFields(t *testing.T) {
	type testStruct struct {
		id   string
		name string
	}
	testData := testStruct{
		id:   "1234567",
		name: "Test Data",
	}

	Convey("That true is returned because 'name' is present", t, func() {
		So(HasField(testData, "name"), ShouldEqual, true)
	})
	Convey("That false is returned because 'dataset' is not present", t, func() {
		So(HasField(testData, "dataset"), ShouldEqual, false)
	})
}

func TestConcatenateStrings(t *testing.T) {
	Convey("That the returned value is https://www.ons.gov.uk/datasets/cpih01 for filterable pages", t, func() {
		got := ConcatenateStrings("www.ons.gov.uk", "/datasets/", "cpih01")
		want := "www.ons.gov.uk/datasets/cpih01"
		So(got, ShouldEqual, want)
	})
}

func TestNotLastItem(t *testing.T) {
	Convey("That true is returned when index is 1 and length is 3", t, func() {
		got := NotLastItem(3, 1)
		want := true
		So(got, ShouldEqual, want)
	})
	Convey("That false is returned when index is 0 and length is 1", t, func() {
		got := NotLastItem(1, 0)
		want := false
		So(got, ShouldEqual, want)
	})

}

func TestDatePeriodFormat(t *testing.T) {
	Convey("Given a time-series monthly string", t, func() {
		want := "Jun - Sep 2010"
		got := DatePeriodFormat("2010 JUN-SEP")
		So(got, ShouldEqual, want)

	})
	Convey("Given a time-series yearly string", t, func() {
		want := "2019"
		got := DatePeriodFormat("2019")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series month string", t, func() {
		want := "Feb 2018"
		got := DatePeriodFormat("2018 FEB")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series Q1 string", t, func() {
		want := "Jan - Mar 2019"
		got := DatePeriodFormat("Q1 2019")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series Q2 string", t, func() {
		want := "Apr - Jun 2019"
		got := DatePeriodFormat("Q2 2019")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series Q3 string", t, func() {
		want := "Jul - Sep 2019"
		got := DatePeriodFormat("Q3 2019")
		So(got, ShouldEqual, want)
	})
	Convey("Given a time-series Q4 string", t, func() {
		want := "Oct - Dec 2019"
		got := DatePeriodFormat("Q4 2019")
		So(got, ShouldEqual, want)
	})

}

func TestTruncateToMaximumCharacters(t *testing.T) {
	Convey("That text is not truncated when it contains fewer characters than maxLength", t, func() {
		got := TruncateToMaximumCharacters("Hello world", 50)
		want := "Hello world"
		So(got, ShouldEqual, want)
	})
	Convey("That text is truncated when exceeding maxLength", t, func() {
		got := TruncateToMaximumCharacters("Hello world", 2)
		want := "He..."
		So(got, ShouldEqual, want)
	})
	Convey("That leading/trailling whitespace is removed when text is truncated", t, func() {
		got := TruncateToMaximumCharacters("The space after 'space' should not be included", 10)
		want := "The space..."
		So(got, ShouldEqual, want)
	})
}

func TestTrimPrefixedPeriod(t *testing.T) {
	Convey("That the returned string is myString", t, func() {
		got := TrimPrefixedPeriod(".myString")
		want := "myString"
		So(got, ShouldEqual, want)
	})
	Convey("That all periods are trimmed", t, func() {
		got := TrimPrefixedPeriod(".....myString")
		want := "myString"
		So(got, ShouldEqual, want)
	})
	Convey("That only leading periods are trimmed", t, func() {
		got := TrimPrefixedPeriod(".string with periods on the end....")
		want := "string with periods on the end...."
		So(got, ShouldEqual, want)
	})
	Convey("That a string with no period is returned", t, func() {
		got := TrimPrefixedPeriod("a string with no period")
		want := "a string with no period"
		So(got, ShouldEqual, want)
	})
}
