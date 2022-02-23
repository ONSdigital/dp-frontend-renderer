package render

import (
	"context"
	"fmt"
	"html/template"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ONSdigital/dp-frontend-renderer/assets"

	"github.com/BurntSushi/toml"
	"github.com/ONSdigital/go-ns/common"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/c2h5oh/datasize"
	"github.com/gosimple/slug"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const legacyDatasetURIFormat = "/file?uri=%s/%s"

var bundle, _ = InitLocaleBundle()
var localizers = InitLocalizer()

// InitLocalizer is used to initialise the localizer
func InitLocalizer() map[string]*i18n.Localizer {
	m := make(map[string]*i18n.Localizer)
	for _, locale := range common.SupportedLanguages {
		m[locale] = i18n.NewLocalizer(bundle, locale)

	}
	return m
}

// InitLocaleBundle is used to initialise the locale bundle
func InitLocaleBundle() (*i18n.Bundle, error) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	for _, locale := range common.SupportedLanguages {
		filePath := "locales/active." + locale + ".toml"
		asset, err := assets.Asset(filePath)
		if err != nil {
			log.Error(context.Background(), "failed to get locale file", err)
		}
		bundle.ParseMessageFileBytes(asset, filePath)
	}

	return bundle, nil
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
		log.Error(context.Background(), "failed to parse time", err)
		return template.HTML(s)
	}
	localiseTime(&t)
	return template.HTML(t.Format("02 January 2006"))
}

func DateFormatYYYYMMDD(s string) template.HTML {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(context.Background(), "failed to parse time", err)
		return template.HTML(s)
	}
	localiseTime(&t)
	return template.HTML(t.Format("2006/01/02"))
}
func DateFormatYYYYMMDDNoSlash(s string) template.HTML {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Error(context.Background(), "failed to parse time", err)
		return template.HTML(s)
	}
	localiseTime(&t)
	return template.HTML(t.Format("20060102"))
}

func localiseTime(t *time.Time) time.Time {
	tz, err := time.LoadLocation("Europe/London")
	if err != nil {
		log.Error(context.Background(), "failed to load time zone location", err)
		return *t
	}
	return t.In(tz)
}

// DatePeriodFormat will format a time-series date period string to a human accessible format e.g.
// "2019 JAN-FEB" to "Jan - Feb 2019"
// "2010 Q1" to "Jan - Mar 2010"
func DatePeriodFormat(s string) string {
	dashIndex := strings.Index(s, "-")
	// 1. Add spaces around dash
	if dashIndex > -1 {
		charIndexAfterDash := dashIndex + 1
		s = s[:dashIndex] + " - " + s[charIndexAfterDash:]
	}

	//2. Replace Q1 Q2 Q3 Q4 with their quarterly month representation e.g. Apr - Jun
	Q1 := strings.Index(s, "Q1")
	if Q1 > -1 {
		Q1EndIndex := Q1 + 2
		s = s[:Q1] + "Jan - Mar" + s[Q1EndIndex:]
	}
	Q2 := strings.Index(s, "Q2")
	if Q2 > -1 {
		Q2EndIndex := Q2 + 2
		s = s[:Q2] + "Apr - Jun" + s[Q2EndIndex:]
	}
	Q3 := strings.Index(s, "Q3")
	if Q3 > -1 {
		Q3EndIndex := Q3 + 2
		s = s[:Q3] + "Jul - Sep" + s[Q3EndIndex:]

	}
	Q4 := strings.Index(s, "Q4")
	if Q4 > -1 {
		Q4EndIndex := Q4 + 2
		s = s[:Q4] + "Oct - Dec" + s[Q4EndIndex:]

	}

	year, err := strconv.Atoi(s[:4])
	// 3. Move year to end of string if present and insert a space
	if err == nil {
		// Not just displaying year but month as well
		if len(s) > 5 {
			// YYYY[space]DEC[space]-[space]JAN = 14 characters
			if len(s) == 14 {
				monthStart, _ := time.Parse("Jan", s[5:8])
				monthEnd, _ := time.Parse("Jan", s[11:])

				dateStart := monthStart.AddDate(year, 0, 0)
				dateEnd := monthEnd.AddDate(year, 0, 0)

				if dateStart.After(dateEnd) {
					dateEnd = dateEnd.AddDate(1, 0, 0)
					s = dateStart.Format("Jan 2006")
				} else {
					s = dateStart.Format("Jan")
				}
				s = s + " - " + dateEnd.Format("Jan 2006")
			} else {
				// YYYY[space] = 5 characters
				postYearIndex := 5
				s = s[postYearIndex:] + " " + s[:4]
			}
		}
	}
	// 4. Convert BLOCK CAPS to Title Caps
	timePeriodFormatted := strings.Title(strings.ToLower(s))
	return timePeriodFormatted
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

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
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
func Localise(key string, language string, plural int, templateArguments ...string) string {
	if key == "" {
		err := fmt.Errorf("key " + key + " not found in locale file")
		log.Error(context.Background(), "no locale look up key provided", err)
		return ""
	}
	if language == "" {
		language = "en"
	}

	// Configure template data for arguments in strings
	templateData := make(map[string]string)
	for i, argument := range templateArguments {
		stringIndex := strconv.Itoa(i)
		key := "arg" + stringIndex
		templateData[key] = argument
	}

	loc := localizers[language]
	translation := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    key,
		PluralCount:  plural,
		TemplateData: templateData,
	})
	return translation
}

func DomainSetLang(domain string, uri string, language string) string {
	languageSupported := false
	for _, locale := range common.SupportedLanguages {
		if locale == language {
			languageSupported = true
		}
	}

	// uri comes in inconsistently, remove domain and port if they come through in uri param
	var findEndpointRE = regexp.MustCompile(`https?://[^/]+(.*)`)
	if endpoint := findEndpointRE.FindStringSubmatch(uri); len(endpoint) == 2 {
		uri = endpoint[1]
	}

	url := domain + uri

	strippedURL := strings.Replace(url, "https://", "", 1)
	strippedURL = strings.Replace(strippedURL, "www.", "", 1)

	for _, locale := range common.SupportedLanguages {
		possibleLocaleURLPrefix := strippedURL[0:len(locale)]

		if possibleLocaleURLPrefix == locale {
			trimLength := len(locale) + 1
			strippedURL = strippedURL[trimLength:]
			break
		}
	}

	domainWithTranslation := ""
	if !languageSupported {
		err := fmt.Errorf("Language: " + language + " is not supported resolving to " + common.DefaultLang)
		log.Error(context.Background(), "language fail", err)
	}
	if language == common.DefaultLang || !languageSupported {
		domainWithTranslation = "https://www." + strippedURL
	} else {
		domainWithTranslation = "https://" + language + "." + strippedURL
	}

	return domainWithTranslation
}

// HasField checks to see if the field is present in the struct
func HasField(data interface{}, name string) bool {
	rv := reflect.ValueOf(data)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}

// ConcatenateStrings takes a number of string arguments and concatenates them
func ConcatenateStrings(tokens ...string) string {
	var result strings.Builder
	for _, token := range tokens {
		result.WriteString(token)
	}
	return result.String()
}

// NotLastItem returns true/false based on if the index equals the length
// Example of use is in JSON-LD partials, where we must determine whether or not a comma should be rendered in a range
func NotLastItem(length, index int) bool {
	if index < length-1 {
		return true
	}
	return false
}

// TruncateToMaximumCharacters returns a substring of parameter 'text' if the text is longer than the specified maximum length
func TruncateToMaximumCharacters(text string, maxLength int) string {
	if len(text) < maxLength {
		return text
	}

	truncatedText := text[0:maxLength]
	return strings.TrimSpace(truncatedText) + "..."
}
