package render

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"time"

	"github.com/ONSdigital/go-ns/log"
	"github.com/c2h5oh/datasize"
	"github.com/gosimple/slug"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const legacyDatasetURIFormat = "/file?uri=%s/%s"

func HumanSize(size string) (string, error) {
	if size == "" {
		return "", nil
	}
	s, err := strconv.Atoi(size)
	if err != nil {
		return "", err
	}
	return datasize.ByteSize(s).HumanReadable(), nil
}

func SafeHTML(s string) template.HTML {
	return template.HTML(s)
}

func DateFormat(s string) template.HTML {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(err, nil)
		return template.HTML(s)
	}
	return template.HTML(t.Format("02 January 2006"))
}

func DateFormatYYYYMMDD(s string) template.HTML {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(err, nil)
		return template.HTML(s)
	}
	return template.HTML(t.Format("2006/01/02"))
}

func Last(x int, a interface{}) bool {
	return x == reflect.ValueOf(a).Len()-1
}

func Loop(n, m int) []int {
	arr := make([]int, m-n)
	v := n
	for i := 0; i < m-v; i++ {
		arr[i] = n
		n++
	}
	return arr
}

func Subtract(x, y int) int {
	return x - y
}

func Slug(s string) string {
	return slug.Make(s)
}

// LegacyDataSetDownloadURI builds a URI string for a legacy dataset download URI.
func LegacyDataSetDownloadURI(pageURI, filename string) string {
	// Concatenation of strings inside a Href tag causes the URI value to be HTML escaped.
	// The preference is for our links not to be escaped to maintain readability. To remedy this we build
	// the link inside this func which is then inserted into template.
	return fmt.Sprintf(legacyDatasetURIFormat, pageURI, filename)
}

// Markdown converts markdown to HTML
func Markdown(md string) template.HTML {
	// lot's of the markdown we currently have stored doesn't match markdown title specs
	// currently it has no space between the hashes and the title text e.g. ##Title
	// to use our new markdown parser we have add a space e.g. ## Title

	byteToCheck := byte('#')
	var previousByte byte
	isTitleTag := false

	for i := 0; i < len(md); i++ {
		// current byte iteration is a part of a title
		if isTitleTag {
			// if current iteration is a hash continue (until we find the the end of the hashes)
			if md[i] == byteToCheck {
				continue
			}

			isTitleTag = false
			previousByte = md[i]

			// if a space already exists after hashes continue and don't add a space
			if md[i] == ' ' {
				continue
			}

			// add a space between hashes and title
			md = md[0:i] + " " + md[i:]
			continue
		}

		// check current iteration is title tag
		if md[i] == byteToCheck {
			// if previous btye and current match
			if previousByte == md[i] {
				isTitleTag = true
				continue
			}
			previousByte = md[i]
		}

	}

	s := blackfriday.Run([]byte(fmt.Sprintf("%s", md)))
	return template.HTML(s)
}
