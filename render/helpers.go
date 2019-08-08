package render

import (
	"fmt"
	"html/template"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/ONSdigital/dp-frontend-renderer/config"
	"github.com/ONSdigital/go-ns/log"
	"github.com/c2h5oh/datasize"
	"github.com/gosimple/slug"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const legacyDatasetURIFormat = "/file?uri=%s/%s"

var bundle = InitLocaleBundle()

// InitLocaleBundle is used to initialise the locale bundle
func InitLocaleBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	workingDir, err := os.Getwd()
	if err != nil {
		log.Error(err, nil)
	}
	parts := strings.SplitAfter(workingDir, "dp-frontend-renderer/")
	projectRootDir := parts[0]
	for _, locale := range config.SupportedLanguages {
		filePath := projectRootDir + "/assets/locales/active." + locale + ".toml"
		bundle.MustLoadMessageFile(filePath)
	}
	return bundle
}

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
	re := regexp.MustCompile(`(##+)([^\s#])`)

	modifiedMarkdown := strings.Builder{}
	for _, line := range strings.Split(md, "\n") {
		modifiedMarkdown.WriteString(fmt.Sprintf("%s\n", re.ReplaceAllString(line, "$1 $2")))
	}

	s := blackfriday.Run([]byte(fmt.Sprintf("%s", modifiedMarkdown.String())))
	return template.HTML(s)
}

// Localise localises text based on a key
func Localise(key string, language string, plural int) string {
	if key == "" {
		return ""
	}
	if language == "" {
		language = "en"
	}

	// Call i18n to get the translations
	loc := i18n.NewLocalizer(bundle, language)
	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID:   key,
		PluralCount: plural,
	})
	return translation
}
