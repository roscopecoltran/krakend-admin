// sitemap
package sitemap

/*
TODO:
	auto detect files
	auto unzip files
	iterable lazy initialized collections
	integrate RSS feeds?
*/

import (
	//"encoding/xml"
	"bytes"
	"errors"
	"fmt"
	"gocrawl/datastore"
	"io"
	"net/http"
	//	"math"
	//	"net/url"
	//	"strconv"
	//	"strings"
	"github.com/roscopecoltran/krakend-admin/sitemap/netrule"
	"github.com/roscopecoltran/krakend-admin/utils"
	//"time"
	//"io/ioutil"
	//"os"
)

/*
// Internal xml structures for parsing sitemap indices
type sitemap_index_entry struct {
	XML_name      xml.Name `xml:"sitemap"`
	Index_URL_raw string   `xml:"loc"`
	Last_mod      string   `xml:"lastmod"`
}

type sitemap_index_set struct {
	XML_name  xml.Name              `xml:"sitemapindex"`
	Index_set []sitemap_index_entry `xml:"sitemap"`
}

// Data structure containing information in a sitemap index

type SitemapIndex struct {
	URLs map[string]SitemapIndexURL
}

type SitemapIndexURL struct {
	Location     string
	LastModified time.Time
}

// Data structure containing information in a sitemap.
type Sitemap struct {
	URLs map[string]SitemapURL
}

type SitemapURL struct {
	Location         string
	LastModified    time.Time
	ChangeFrequency string
	Priority         float32
}


// Error type for error streams
type SitemapURL_Parse_Error struct {
	Field       string
	Err_Msg     string
	Parsed_Text string
	Location    string
}

func (e *SitemapURL_Parse_Error) Error() string {
	return e.Err_Msg
}

func NewSitemapURL_Parse_Error(field, err_msg, parsed_text, location string) *SitemapURL_Parse_Error {
	var full_err_msg = fmt.Sprintf("%s (%s): %s", location, field, err_msg)
	return &SitemapURL_Parse_Error{Field: field, Err_Msg: full_err_msg, Parsed_Text: parsed_text, Location: location}
}

func ParseSitemapIndexXml(xml_text string, err_stream chan *SitemapURL_Parse_Error) (sitemap_index *SitemapIndex, err error) {
	sitemap_index = &SitemapIndex{URLs: make(map[string]SitemapIndexURL)}
	var output sitemap_index_set
	if err = xml.Unmarshal([]byte(xml_text), &output); err != nil {
		return sitemap_index, err
	}
	for _, element := range output.Index_set {
		var sitemap_index_url = SitemapIndexURL{}

		var url_unescaped, err = parseURL(element.Index_URL_raw, err_stream)
		if err != nil {
			continue
		}

		sitemap_index_url.Location = url_unescaped
		sitemap_index_url.LastModified, _ = parseLastModified(element.Last_mod, url_unescaped, err_stream)

		sitemap_index.URLs[url_unescaped] = sitemap_index_url
	}
	return sitemap_index, nil

}

// ParseSitemap will convert a string of plain text or xml from a sitemap to a Sitemap structure.
// This function will never return an error.
// All errors during processing the valid plain text or xml document will be dumped to the err_stream.
func ParseSitemap(text string, err_stream chan *SitemapURL_Parse_Error) (sitemap *Sitemap, err error) {
	if err_stream != nil {
		defer close(err_stream)
	}

	if sitemap, err = parseSitemapXml(text, err_stream); err != nil {
		sitemap, err = parseSitemapPlaintext(text, err_stream)
	}

	return sitemap, err
}

// ParseSitemapPlaintext will convert a string of plain text from a sitemap to a Sitemap structure.
// This function will never return an error.
// All errors during processing the valid plain text document will be dumped to the err_stream.
func ParseSitemapPlaintext(plain_text string, err_stream chan *SitemapURL_Parse_Error) (sitemap *Sitemap, err error) {
	if err_stream != nil {
		defer close(err_stream)
	}
	sitemap, err = parseSitemapPlaintext(plain_text, err_stream)
	return sitemap, err

}

// parseSitemapPlaintext will convert a string of plain text from a sitemap to a Sitemap structure.
// This internal function will never return an error.
// All errors during processing the valid plain text document will be dumped to the err_stream.
func parseSitemapPlaintext(plain_text string, err_stream chan *SitemapURL_Parse_Error) (sitemap *Sitemap, err error) {
	sitemap = &Sitemap{URLs: make(map[string]SitemapURL)}
	var urls_split = strings.Split(plain_text, "\n")
	for _, url_raw := range urls_split {
		var sitemap_url = SitemapURL{LastModified: time.Time{}, ChangeFrequency: "", Priority: 1.0}
		var url_unescaped, err = parseURL(url_raw, err_stream)
		if err != nil {
			continue
		}

		sitemap_url.Location = url_unescaped

		sitemap.URLs[url_unescaped] = sitemap_url
	}
	return sitemap, err
}

// ParseSitemapXml will convert a string of xml text from a sitemap to ta Sitemap structure.
// This function will return an error only when the document is not valid xml.
// All errors during processing the valid xml document will be dumped to the err_stream.
func ParseSitemapXml(xml_text string, err_stream chan *SitemapURL_Parse_Error) (sitemap *Sitemap, err error) {
	if err_stream != nil {
		defer close(err_stream)
	}
	sitemap, err = parseSitemapXml(xml_text, err_stream)
	return sitemap, err

}

// parseSitemapXml will convert a string of xml text from a sitemap to a Sitemap structure.
// This internal function will return an error only when the document is not valid xml.
// All errors during processing the valid xml document will be dumped to the err_stream.
func parseSitemapXml(xml_text string, err_stream chan *SitemapURL_Parse_Error) (sitemap *Sitemap, err error) {
	sitemap = &Sitemap{URLs: make(map[string]SitemapURL)}
	var output sitemap_url_set
	if err = xml.Unmarshal([]byte(xml_text), &output); err != nil {
		return sitemap, err
	}
	for _, element := range output.URL_set {
		var sitemap_url = SitemapURL{}

		var url_unescaped, err = parseURL(element.URL_raw, err_stream)
		if err != nil {
			continue
		}

		sitemap_url.Location = url_unescaped
		sitemap_url.Priority, _ = parsePriority(element.Priority, url_unescaped, err_stream)
		sitemap_url.ChangeFrequency, _ = parseChangeFrequency(element.Change_freq, url_unescaped, err_stream)
		sitemap_url.LastModified, _ = parseLastModified(element.Last_mod, url_unescaped, err_stream)

		sitemap.URLs[url_unescaped] = sitemap_url
	}
	return sitemap, nil
}

//*/
// Object reform

// Sitemap deals with any sitemap file in xml format for index, and xml or plain text format for pages, with or without gzip compression for both
type Sitemap struct {
	sitemap_datastore datastore.DataStore

	access_rule   netrule.NetAccessRule
	sitemap_index *SitemapIndex
	sitemap_page  *SitemapPage
	err           error
}

// Resets to the beginning of the document, if it is an index then it will also reset to the beginning of the index.
func (sitemap *Sitemap) Reset() (err error) {
	sitemap.err = nil
	if sitemap.sitemap_index == nil && sitemap.sitemap_page != nil {
		return sitemap.sitemap_page.Reset()
	} else if sitemap.sitemap_index != nil {
		sitemap.sitemap_page = nil
		return sitemap.sitemap_index.Reset()

	}
	return nil
}

// Contains the error from the previous Next(), NextPage(), or NextOnPage() call
func (sitemap *Sitemap) Err() (err error) {
	return sitemap.err
}

// Contains the item from the previous Next() or NextOnPage() call
func (sitemap *Sitemap) Item() (item *SitemapURL) {
	if sitemap.sitemap_page != nil {
		return sitemap.sitemap_page.Item()
	}
	return nil
}

// Contains the item from the previous NextPage() call
func (sitemap *Sitemap) IndexItem() (item *SitemapIndexURL) {
	if sitemap.sitemap_index != nil {
		return sitemap.sitemap_index.Item()
	}
	return nil
}

// Seemless next, will not return false until the entire index is exhausted or an error has occurred,
// and will get the next page in the index based on the ruleset
// False return value and nil Err() indicate end of stream.
func (sitemap *Sitemap) Next() (success bool) {
	if sitemap.sitemap_page != nil {
		success = sitemap.sitemap_page.Next()
		if success || sitemap.sitemap_page.Err() != nil {
			sitemap.err = sitemap.sitemap_page.Err()
			return success
		}
	}

	// got here due to sitemap_page == nil || (!success && sitemap_page.Err() == nil)
	// lets try to grab next entry in sitemap_index
	if sitemap.sitemap_index != nil {
		success = sitemap.NextPage()
		if !success && sitemap.Err() != nil {
			return false
		}

		if sitemap.sitemap_page != nil {
			success = sitemap.sitemap_page.Next()
			if success || sitemap.sitemap_page.Err() != nil {
				sitemap.err = sitemap.sitemap_page.Err()
				return success
			}
		}
	}
	return false

}

// Checks if the access_rule condition is met, blocks until it is met
func canNetAccess(access_rule netrule.NetAccessRule) {
	if access_rule != nil {
		var access_chan = access_rule.AccessChannel()
		if access_chan != nil {
			for {
				select {
				case <-access_chan:
					return
				default:
				}
			}
		}
	}
}

// Grabs the next page in the index and sets sitemap_page to it, returns false on error or when the index is exhausted
// False return value and nil Err() indicate end of stream.
func (sitemap *Sitemap) NextPage() (success bool) {
	if sitemap.sitemap_index != nil {
		success := sitemap.sitemap_index.Next()
		// End of index stream
		if !success {
			sitemap.err = sitemap.sitemap_index.Err()
			sitemap.sitemap_page = nil
			return false
		} else { // Next entry
			index_item := sitemap.sitemap_index.Item()
			access_rule := sitemap.access_rule
			// rule for grabbing next has been met
			canNetAccess(access_rule)

			response, err := http.Get(index_item.Location())
			if err != nil {
				sitemap.err = err
				sitemap.sitemap_page = nil
				return false
			}
			defer response.Body.Close()
			sitemap.sitemap_page, err = NewSitemapPage(response.Body)
			if err != nil {
				sitemap.err = err
				sitemap.sitemap_page = nil
				return false
			}
			if access_rule != nil {
				go access_rule.Accessed()
			}
			return true
		}
	}
	sitemap.err = nil
	return false
}

// Grabs the next item on this page, returns false on error or when all contents of the page are exhausted
// False return value and nil Err() indicate end of stream.
func (sitemap *Sitemap) NextOnPage() (success bool) {
	if sitemap.sitemap_page != nil {
		success = sitemap.sitemap_page.Next()
		sitemap.err = sitemap.sitemap_page.Err()
		return success
	}
	sitemap.err = nil
	return false
}

func NewSitemap(reader io.Reader) (sitemap *Sitemap, err error) {
	sitemap = &Sitemap{}
	err = sitemap.ParseSitemap(reader)
	if err != nil {
		return nil, err
	}
	return sitemap, nil
}

func (sitemap *Sitemap) DataStore() (data_store datastore.DataStore) {
	return sitemap.sitemap_datastore
}

func (sitemap *Sitemap) SetDataStore(data_store datastore.DataStore) {
	sitemap.Close()
	if data_store != nil {
		sitemap.sitemap_datastore = data_store
	} else {

	}
}

func (sitemap *Sitemap) ParseSitemap(reader io.Reader) (err error) {
	sitemap.Close()
	sitemap_index, mem_seek, index_err := newSitemapIndex(reader)
	if index_err != nil {
		sitemap_page, _, page_err := newSitemapPageNoGzip(mem_seek)
		if page_err != nil {
			return errors.New(fmt.Sprintf("sitemap: reader was not in a sitemap format\n index_err: %s\n page_err: %s", index_err, page_err))
		}
		sitemap.sitemap_page = sitemap_page
	} else {
		sitemap.sitemap_index = sitemap_index
	}
	return nil
}

func (sitemap *Sitemap) Close() (err error) {

	return nil
}

// Sitemap index deals with a single sitemap index file in xml format, with or without gzip compression
type SitemapIndex struct {
	sitemap_datastore datastore.DataStore

	sitemap_iterator SitemapIndexIterator
}

func (sitemap_index *SitemapIndex) Reset() (err error) {
	return sitemap_index.sitemap_iterator.Reset()
}
func (sitemap_index *SitemapIndex) Err() (err error) {
	return sitemap_index.sitemap_iterator.Err()
}
func (sitemap_index *SitemapIndex) Item() (item *SitemapIndexURL) {
	return sitemap_index.sitemap_iterator.Item()
}

func (sitemap_index *SitemapIndex) Next() (success bool) {
	return sitemap_index.sitemap_iterator.Next()
}

const max_sitemap_index_size = 10485760

var max_sitemap_index_size_error = errors.New(fmt.Sprintf("Reader uncompressed size exceeds sitemap index maximum of %d", max_sitemap_index_size))

// external assumes nothing and copied the stream regardless
func NewSitemapIndex(reader io.Reader) (sitemap_index *SitemapIndex, err error) {
	sitemap_index, _, err = newSitemapIndex(reader)
	return sitemap_index, err
}

// assumes nothing and copies stream
func newSitemapIndex(reader io.Reader) (sitemap_index *SitemapIndex, internal_reader *bytes.Reader, err error) {
	sitemap_index = &SitemapIndex{}
	internal_reader, err = sitemap_index.parseSitemapIndex(reader)
	if err != nil {
		return nil, internal_reader, err
	}
	return sitemap_index, internal_reader, nil
}

// assumes direct stream, does not copy
func newSitemapIndexNoGzip(mem_seek *bytes.Reader) (sitemap_index *SitemapIndex, internal_reader *bytes.Reader, err error) {
	sitemap_index = &SitemapIndex{}
	internal_reader, err = sitemap_index.parseSitemapIndexNoGzip(mem_seek)
	if err != nil {
		return nil, internal_reader, err
	}
	return sitemap_index, internal_reader, nil
}

func (sitemap_index *SitemapIndex) DataStore() (data_store datastore.DataStore) {
	return sitemap_index.sitemap_datastore
}

func (sitemap_index *SitemapIndex) SetDataStore(data_store datastore.DataStore) {
	sitemap_index.Close()
	if data_store != nil {
		sitemap_index.sitemap_datastore = data_store
	} else {

	}
}

func (sitemap_index *SitemapIndex) ParseSitemapIndex(reader io.Reader) (err error) {
	_, err = sitemap_index.parseSitemapIndex(reader)
	return err
}

func (sitemap_index *SitemapIndex) parseSitemapIndex(reader io.Reader) (byte_reader *bytes.Reader, err error) {
	sitemap_index.Close()
	mem_seek, gzip_err := utils.UnGzipTransform(reader, max_sitemap_index_size+1)
	if gzip_err != nil && !gzip_err.GzipFailed() {
		return mem_seek, errors.New(fmt.Sprintf("sitemap: ungzipping failed: %s", gzip_err))
	}
	byte_reader, parse_err := sitemap_index.parseSitemapIndexNoGzip(mem_seek)
	if parse_err != nil {
		err = errors.New(fmt.Sprintf("sitemap: failed to parse reader\ngzip:%s\n%s", gzip_err, parse_err))
	}
	return byte_reader, err
}

func (sitemap_index *SitemapIndex) parseSitemapIndexNoGzip(mem_seek *bytes.Reader) (byte_reader *bytes.Reader, err error) {
	// validate uncompressed size
	if mem_seek.Len() > max_sitemap_index_size {
		return mem_seek, max_sitemap_index_size_error
	}

	err = sitemap_index.determineIteratorFormat(mem_seek)
	return mem_seek, err
}

// should assume a format of copied stream
func (sitemap_index *SitemapIndex) determineIteratorFormat(reader io.ReadSeeker) (err error) {
	xml_reader, xml_err := newXmlSitemapIndexIterator(reader)
	sitemap_index.sitemap_iterator = xml_reader
	if xml_err != nil {
		return errors.New(fmt.Sprintf("failed to recognize data format, xml error: %s\n", xml_err))

	}
	return nil
}

func (sitemap_index *SitemapIndex) Close() (err error) {

	return nil
}

// Sitemap page deals with a single sitemap file in xml or plain text format, with or without gzip compression
type SitemapPage struct {
	sitemap_datastore datastore.DataStore

	sitemap_iterator SitemapPageIterator
}

func (sitemap_page *SitemapPage) Reset() (err error) {
	return sitemap_page.sitemap_iterator.Reset()
}
func (sitemap_page *SitemapPage) Err() (err error) {
	return sitemap_page.sitemap_iterator.Err()
}
func (sitemap_page *SitemapPage) Item() (item *SitemapURL) {
	return sitemap_page.sitemap_iterator.Item()
}

func (sitemap_page *SitemapPage) Next() (success bool) {
	return sitemap_page.sitemap_iterator.Next()
}

const max_sitemap_page_size = 52428800

var max_sitemap_page_size_error = errors.New(fmt.Sprintf("Reader uncompressed size exceeds sitemap page maximum of %d", max_sitemap_page_size))

// external assumes nothing and copied the stream regardless
func NewSitemapPage(reader io.Reader) (sitemap *SitemapPage, err error) {
	sitemap, _, err = newSitemapPage(reader)
	return sitemap, err
}

// assumes nothing, copies stream
func newSitemapPage(reader io.Reader) (sitemap *SitemapPage, internal_reader *bytes.Reader, err error) {
	sitemap = &SitemapPage{}
	internal_reader, err = sitemap.parseSitemapPage(reader)
	if err != nil {
		return nil, internal_reader, err
	}
	return sitemap, internal_reader, nil
}

// assumes direct stream, does not copy
func newSitemapPageNoGzip(mem_seek *bytes.Reader) (sitemap *SitemapPage, internal_reader *bytes.Reader, err error) {
	sitemap = &SitemapPage{}
	internal_reader, err = sitemap.parseSitemapPageNoGzip(mem_seek)
	if err != nil {
		return nil, internal_reader, err
	}
	return sitemap, internal_reader, nil
}

func (sitemap *SitemapPage) DataStore() (data_store datastore.DataStore) {
	return sitemap.sitemap_datastore
}

func (sitemap *SitemapPage) SetDataStore(data_store datastore.DataStore) {
	sitemap.Close()
	if data_store != nil {
		sitemap.sitemap_datastore = data_store
	} else {

	}
}

func (sitemap *SitemapPage) ParseSitemapPage(reader io.Reader) (err error) {
	_, err = sitemap.parseSitemapPage(reader)
	return err
}

func (sitemap *SitemapPage) parseSitemapPage(reader io.Reader) (byte_reader *bytes.Reader, err error) {
	sitemap.Close()
	mem_seek, gzip_err := utils.UnGzipTransform(reader, max_sitemap_page_size+1)
	if gzip_err != nil && !gzip_err.GzipFailed() {
		return mem_seek, errors.New(fmt.Sprintf("sitemap: ungzipping failed: %s", gzip_err))
	}
	byte_reader, parse_err := sitemap.parseSitemapPageNoGzip(mem_seek)
	if parse_err != nil {
		err = errors.New(fmt.Sprintf("sitemap: failed to parse reader\ngzip:%s\n%s", gzip_err, parse_err))
	}
	return byte_reader, err
}

func (sitemap *SitemapPage) parseSitemapPageNoGzip(mem_seek *bytes.Reader) (byte_reader *bytes.Reader, err error) {
	// validate uncompressed size
	if mem_seek.Len() > max_sitemap_page_size {
		return mem_seek, max_sitemap_page_size_error
	}

	err = sitemap.determineIteratorFormat(mem_seek)
	return mem_seek, err
}

// should assume a format of copied stream
func (sitemap *SitemapPage) determineIteratorFormat(reader io.ReadSeeker) (err error) {
	xml_reader, xml_err := newXmlSitemapPageIterator(reader)
	sitemap.sitemap_iterator = xml_reader
	if xml_err != nil {
		plain_reader, plain_err := newPlainSitemapPageIterator(reader)
		sitemap.sitemap_iterator = plain_reader
		if plain_err != nil {
			return errors.New(fmt.Sprintf("failed to recognize data format, xml error: %s\nplain text error:%s", xml_err, plain_err))
		}
	}
	return nil
}

func (sitemap *SitemapPage) Close() (err error) {

	return nil
}
