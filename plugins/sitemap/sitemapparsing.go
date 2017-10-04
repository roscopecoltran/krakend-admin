// sitemapparsing.go
package sitemap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// Error type to determine if parse failure was fatal (location)
// or merely optional info (change frequency, last modified, priority).
type urlParseError struct {
	error_msg  string
	url_failed bool
}

func newUrlParseError(msg string, url_fail bool) (err *urlParseError) {
	return &urlParseError{error_msg: msg, url_failed: url_fail}
}

func (err *urlParseError) Error() (msg string) {
	return err.error_msg
}

func (err *urlParseError) UrlParseFailed() (failed bool) {
	return err.url_failed
}

// struct containing validated data for a sitemap url node
type SitemapURL struct {
	location         string
	last_modified    time.Time
	change_frequency string
	priority         float32
}

func (sitemap_url *SitemapURL) Location() string {
	return sitemap_url.location
}
func (sitemap_url *SitemapURL) LastModified() time.Time {
	return sitemap_url.last_modified
}
func (sitemap_url *SitemapURL) ChangeFrequency() string {
	return sitemap_url.change_frequency
}
func (sitemap_url *SitemapURL) Priority() float32 {
	return sitemap_url.priority
}

// Internal xml structures for parsing sitemaps
type sitemap_url_set struct {
	XML_name xml.Name            `xml:"urlset"`
	URL_set  []sitemap_url_entry `xml:"url"`
}

type sitemap_url_entry struct {
	XML_name    xml.Name `xml:"url"`
	URL_raw     string   `xml:"loc"`
	Last_mod    string   `xml:"lastmod"`
	Change_freq string   `xml:"changefreq"`
	Priority    string   `xml:"priority"`
}

// Will attempt to parse all values of the sitemap URL struct, will try all fields and return the first error that occurred
func (entry *sitemap_url_entry) parseSitemapURL() (sitemap_url *SitemapURL, err *urlParseError) {
	sitemap_url = &SitemapURL{}
	var temp_err error = nil
	sitemap_url.location, temp_err = ParseURL(entry.URL_raw)
	if temp_err != nil && err == nil {
		err = newUrlParseError(temp_err.Error(), true)
	}
	sitemap_url.priority, temp_err = parsePriority(entry.Priority)
	if temp_err != nil && err == nil {
		err = newUrlParseError(temp_err.Error(), false)
	}
	sitemap_url.change_frequency, temp_err = parseChangeFrequency(entry.Change_freq)
	if temp_err != nil && err == nil {
		err = newUrlParseError(temp_err.Error(), false)
	}
	sitemap_url.last_modified, temp_err = parseLastModified(entry.Last_mod)
	if temp_err != nil && err == nil {
		err = newUrlParseError(temp_err.Error(), false)
	}

	return sitemap_url, err

}

// struct containing validated data for a sitemap index url node
type SitemapIndexURL struct {
	location      string
	last_modified time.Time
}

func (sitemap_index_url *SitemapIndexURL) Location() string {
	return sitemap_index_url.location
}
func (sitemap_index_url *SitemapIndexURL) LastModified() time.Time {
	return sitemap_index_url.last_modified
}

// Internal xml structures for parsing sitemap indices
type sitemap_index_entry struct {
	XML_name xml.Name `xml:"sitemap"`
	URL_raw  string   `xml:"loc"`
	Last_mod string   `xml:"lastmod"`
}

// Will attempt to parse all values of the sitemap URL struct, will try all fields and return the first error that occurred
func (entry *sitemap_index_entry) parseSitemapIndexURL() (sitemap_index_url *SitemapIndexURL, err *urlParseError) {
	sitemap_index_url = &SitemapIndexURL{}
	var temp_err error = nil
	sitemap_index_url.location, temp_err = ParseURL(entry.URL_raw)
	if temp_err != nil && err == nil {
		err = newUrlParseError(temp_err.Error(), true)
	}
	sitemap_index_url.last_modified, temp_err = parseLastModified(entry.Last_mod)
	if temp_err != nil && err == nil {
		err = newUrlParseError(temp_err.Error(), false)
	}

	return sitemap_index_url, err

}

// parse validation functions
var valid_urls = regexp.MustCompile(`^((((https?|ftps?|gopher|telnet|nntp)://)|(mailto:|news:))(%[0-9A-Fa-f]{2}|[-()_.!~*';/?:@&=+$,A-Za-z0-9])+)([).!';/?:,][\t ])?$`)

func ParseURL(url_text string) (parsed_url string, err error) {
	parsed_url, err = url.QueryUnescape(url_text)
	if err != nil || !valid_urls.MatchString(url_text) {
		if err == nil {
			err = errors.New(fmt.Sprintf("\"%s\" is not a valid url", url_text))
		}
		return "", err
	}
	return parsed_url, nil
}

// Appends append_text to url_text while maintaining one forward slash between the two.
// Slashes in append_text are maintained
func AppendURL(url_text string, append_text []string) (url string, err error) {
	append_len := len(append_text)
	if append_len <= 0 {
		return url_text, nil
	}
	last_rune, width := utf8.DecodeLastRuneInString(url_text)
	if last_rune == utf8.RuneError && width == 1 {
		return "", errors.New(fmt.Sprintf("'%s' is not a valid utf8 string", url_text))
	}
	if append_len == 1 && last_rune == '/' {
		url = url_text + append_text[0]
		return url, nil

	} else if append_len > 1 && last_rune == '/' {
		url = strings.Join(append([]string{url_text + append_text[0]}, append_text[1:]...), "/")
		return url, nil
	}
	url = strings.Join(append([]string{url_text}, append_text...), "/")
	return url, nil
}

func parsePriority(priority_text string) (parsed_priority float32, err error) {
	parsed_float, err := strconv.ParseFloat(priority_text, 32)
	if err != nil {
		return 1.0, err
	}
	parsed_priority = float32(math.Min(1.0, math.Max(0.0, parsed_float)))
	return parsed_priority, nil

}

// Valid values for the change frequency field of a sitemap.xml
var valid_freqs = regexp.MustCompile("(?i)(always)|(hourly)|(daily)|(weekly)|(monthly)|(yearly)|(never)")

func parseChangeFrequency(change_freq_text string) (parsed_change_freq string, err error) {
	if !valid_freqs.MatchString(change_freq_text) {
		err = errors.New(fmt.Sprintf("Change Frequency \"%s\" could not be matched to a valid value.", change_freq_text))
		return "", err
	}
	parsed_change_freq = strings.ToLower(change_freq_text)
	return parsed_change_freq, nil
}

// Valid values for ISO8601 date formats
var iso8601_layouts = [...]string{
	//YYYY-MM-DDThh:mm:ss.sTZD (eg 1997-07-16T19:20:30.45+01:00)
	"2006-01-_2T15:04:05.9Z07:00",
	//YYYY-MM-DDThh:mm:ss.sTZD (eg 1997-07-16T19:20:30.45+0100)
	"2006-01-_2T15:04:05.9Z0700",
	//YYYY-MM-DDThh:mm:ssTZD (eg 1997-07-16T19:20:30+01:00)
	"2006-01-_2T15:04:05Z07:00",
	//YYYY-MM-DDThh:mm:ssTZD (eg 1997-07-16T19:20:30+0100)
	"2006-01-_2T15:04:05Z0700",
	//YYYY-MM-DDThh:mmTZD (eg 1997-07-16T19:20+01:00)
	"2006-01-_2T15:04Z07:00",
	//YYYY-MM-DDThh:mmTZD (eg 1997-07-16T19:20+0100)
	"2006-01-_2T15:04Z0700",
	//YYYY-MM-DD (eg 1997-07-16)
	"2006-01-_2",
	//YYYY-MM (eg 1997-07)
	"2006-01",
	//YYYY (eg 1997)
	"2006",
}

func parseLastModified(last_mod_text string) (parsed_last_mod time.Time, err error) {
	for _, layout := range iso8601_layouts {
		parsed_last_mod, err = time.Parse(layout, last_mod_text)
		if err == nil {
			break
		}
	}
	if err != nil {
		return time.Time{}, err
	}
	return parsed_last_mod, nil
}
