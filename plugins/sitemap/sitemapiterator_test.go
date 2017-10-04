package sitemap

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchrcom/testify/assert"
)

/*
func Test_(t *testing.T) {

}

func Benchmark_(b *testing.B) {

}
*/

var plainText_Valid = `https://plus.google.com/103392052013027674779
https://plus.google.com/109779601340825835502
https://plus.google.com/106093162011629633254
https://plus.google.com/108879466279914852493
http://www.venuecom.com?beef=stew&whee=phew%C3%BC`

func Test_plainSitemapPageIterator_Normal(t *testing.T) {
	reader := bytes.NewReader([]byte(plainText_Valid))
	iterator, err := newPlainSitemapPageIterator(reader)
	assert.Nil(t, err, fmt.Sprintf("unexpected error from newPlainSitemapPageIterator: %s", err))

	test_SitemapPageIterator_Normal(t, iterator)
}

var xmlText_Valid = `<?xml version="1.0"?>
<urlset xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
 <url>
    <loc>https://plus.google.com/103392052013027674779</loc>
    <priority>0.5</priority>
    <changefreq>weekly</changefreq>
    <lastmod>2013-11-27T00:06:52-05:00</lastmod>
  </url>
<url>
    <loc>https://plus.google.com/109779601340825835502</loc>
    <priority>0.25</priority>
    <changefreq>monthly</changefreq>
    <lastmod>2013-11-27T00:06:52-05:00</lastmod>
  </url>
 <url>
    <loc>https://plus.google.com/106093162011629633254</loc>
    <priority>0.75</priority>
    <changefreq>yearly</changefreq>
    <lastmod>2013-11-27T00:06:52-05:00</lastmod>
  </url>
<url>
    <loc>https://plus.google.com/108879466279914852493</loc>
    <priority>1</priority>
    <changefreq>daily</changefreq>
    <lastmod>2013-11-27T00:06:52-05:00</lastmod>
  </url>
<url>
    <loc>http://www.venuecom.com?beef=stew&amp;whee=phew%C3%BC</loc>
    <priority>1</priority>
    <changefreq>daily</changefreq>
    <lastmod>2013-11-27T00:06:52-05:00</lastmod>
  </url>
</urlset>`

func Test_xmlSitemapPageIterator_Normal(t *testing.T) {
	reader := bytes.NewReader([]byte(xmlText_Valid))
	iterator, err := newXmlSitemapPageIterator(reader)
	assert.Nil(t, err, fmt.Sprintf("unexpected error from newXmlSitemapPageIterator: %s", err))

	test_SitemapPageIterator_Normal(t, iterator)
}

func test_SitemapPageIterator_Normal(t *testing.T, iterator SitemapPageIterator) {
	var success = iterator.Next()
	assert.True(t, success, fmt.Sprintf("unexpected false from iterator.Next(), err was %s", iterator.Err()))
	for success {
		assert.NotNil(t, iterator.Item(), fmt.Sprintf("unexpected nil from iterator.Item()"))
		assert.Nil(t, iterator.Err(), fmt.Sprintf("unexpected error from iterator.Next(): %s", iterator.Err()))
		success = iterator.Next()
	}
	assert.Nil(t, iterator.Err(), fmt.Sprintf("unexpected error from iterator.Next(): %s", iterator.Err()))

	err := iterator.Reset()
	assert.Nil(t, err, fmt.Sprintf("unexpected error from iterator.Reset(): %s", err))
}

func test_SitemapIndexIterator_Normal(t *testing.T, iterator SitemapIndexIterator) {
	var success = iterator.Next()
	assert.True(t, success, fmt.Sprintf("unexpected false from iterator.Next(), err was %s", iterator.Err()))
	for success {
		assert.NotNil(t, iterator.Item(), fmt.Sprintf("unexpected nil from iterator.Item()"))
		assert.Nil(t, iterator.Err(), fmt.Sprintf("unexpected error from iterator.Next(): %s", iterator.Err()))
		success = iterator.Next()
	}
	assert.Nil(t, iterator.Err(), fmt.Sprintf("unexpected error from iterator.Next(): %s", iterator.Err()))

	err := iterator.Reset()
	assert.Nil(t, err, fmt.Sprintf("unexpected error from iterator.Reset(): %s", err))
}

// Case 1 is an important loc parsing error
// Case 2 is a nonimportant priority parsing error
// Case 3 is a nonimportant changefreq parsing error
// Case 4 is a nonimportant lastmod parsing error
var xmlText_Invalid = `<?xml version="1.0"?>
<urlset xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
<url>
    <loc>https://plus.google.com/108879466279914852493</loc>
    <priority>1</priority>
    <changefreq>daily</changefreq>
    <lastmod>2013-11-27T00:06:52-05:00</lastmod>
  </url>
<url>
    <loc>plus.google.com/103392052013027674779</loc>
    <priority>0.5</priority>
    <changefreq>weekly</changefreq>
    <lastmod>2013-11-27T00:06:52-05:00</lastmod>
  </url>
<url>
    <loc>https://plus.google.com/109779601340825835502</loc>
    <priority>asdf</priority>
    <changefreq>monthly</changefreq>
    <lastmod>2013-11-27T00:06:52-05:00</lastmod>
  </url>
 <url>
    <loc>https://plus.google.com/106093162011629633254</loc>
    <priority>0.75</priority>
    <changefreq>asdf</changefreq>
    <lastmod>2013-11-27T00:06:52-05:00</lastmod>
  </url>
<url>
    <loc>https://plus.google.com/108879466279914852493</loc>
    <priority>1</priority>
    <changefreq>daily</changefreq>
    <lastmod>asdf</lastmod>
  </url>
</urlset>`

func Test_xmlSitemapPageIterator_ErrorConditions(t *testing.T) {
	reader := bytes.NewReader([]byte(xmlText_Invalid))
	iterator, err := newXmlSitemapPageIterator(reader)
	assert.Nil(t, err, fmt.Sprintf("unexpected error from newXmlSitemapPageIterator: %s", err))

	assert.True(t, iterator.Next(), fmt.Sprintf("expected true from iterator.Next() Case 0"))
	assert.Nil(t, iterator.Err(), fmt.Sprintf("unexpected error from iterator.Next() Case 0: %s", iterator.Err()))

	assert.False(t, iterator.Next(), fmt.Sprintf("expected false from iterator.Next() Case 1"))
	assert.NotNil(t, iterator.Err(), fmt.Sprintf("expected error from iterator.Next() Case 1"))

	assert.True(t, iterator.Next(), fmt.Sprintf("expected true from iterator.Next() Case 2"))
	assert.NotNil(t, iterator.Err(), fmt.Sprintf("expected error from iterator.Next() Case 2"))

	assert.True(t, iterator.Next(), fmt.Sprintf("expected true from iterator.Next() Case 3"))
	assert.NotNil(t, iterator.Err(), fmt.Sprintf("expected error from iterator.Next() Case 3"))

	assert.True(t, iterator.Next(), fmt.Sprintf("expected true from iterator.Next() Case 4"))
	assert.NotNil(t, iterator.Err(), fmt.Sprintf("expected error from iterator.Next() Case 4"))
}

var xmlText_WrongFormat = `https://plus.google.com/103392052013027674779`

func Test_xmlSitemapPageIterator_WrongFormat(t *testing.T) {
	reader := bytes.NewReader([]byte(xmlText_WrongFormat))
	_, err := newXmlSitemapPageIterator(reader)
	assert.NotNil(t, err, fmt.Sprintf("expected error from newXmlSitemapPageIterator"))

}
