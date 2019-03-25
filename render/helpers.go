package render

import (
	"fmt"
	"github.com/ONSdigital/go-ns/log"
	"github.com/c2h5oh/datasize"
	"github.com/gosimple/slug"
	"html/template"
	"reflect"
	"strconv"
	"time"
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
