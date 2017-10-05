package provider

import (
	"github.com/jinzhu/gorm"
	"github.com/roscopecoltran/krakend-admin/models/classify"
)

type Provider struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	AlternativeTitleXpath   string `gorm:"column:alternative_title_xpath" json:"alternative_title_xpath,omitempty" yaml:"alternative_title_xpath,omitempty" toml:"alternative_title_xpath,omitempty"`
	AlternativeTorrentXpath string `gorm:"column:alternative_torrent_xpath" json:"alternative_torrent_xpath,omitempty" yaml:"alternative_torrent_xpath,omitempty" toml:"alternative_torrent_xpath,omitempty"`
	APIKey                  string `gorm:"column:api_key" json:"api_key,omitempty" yaml:"api_key,omitempty" toml:"api_key,omitempty"`
	BaseURL                 string `gorm:"column:base_url" json:"base_url,omitempty" example:"https://wiki.archlinux.org" yaml:"base_url,omitempty" example:"https://wiki.archlinux.org" toml:"base_url,omitempty" example:"https://wiki.archlinux.org"`
	BaseYoutubeURL          string `gorm:"column:base_youtube_url" json:"base_youtube_url,omitempty" yaml:"base_youtube_url,omitempty" toml:"base_youtube_url,omitempty"`
	CalendarNameXpath       string `gorm:"column:calendar_name_xpath" json:"calendar_name_xpath,omitempty" yaml:"calendar_name_xpath,omitempty" toml:"calendar_name_xpath,omitempty"`
	ContentMiscXpath        string `gorm:"column:content_misc_xpath" json:"content_misc_xpath,omitempty" yaml:"content_misc_xpath,omitempty" toml:"content_misc_xpath,omitempty"`
	ContentQuery            string `gorm:"column:content_query" json:"content_query,omitempty" yaml:"content_query,omitempty" toml:"content_query,omitempty"`
	ContentXpath            string `gorm:"column:content_xpath" json:"content_xpath,omitempty" yaml:"content_xpath,omitempty" toml:"content_xpath,omitempty"`
	DefaultHostname         string `gorm:"column:default_hostname" json:"default_hostname,omitempty" yaml:"default_hostname,omitempty" toml:"default_hostname,omitempty"`
	DescriptionXpath        string `gorm:"column:description_xpath" json:"description_xpath,omitempty" yaml:"description_xpath,omitempty" toml:"description_xpath,omitempty"`
	Description             string `gorm:"column:desc" json:"desc,omitempty" yaml:"desc,omitempty" toml:"desc,omitempty"`
	EmbeddedURL             string `gorm:"column:embedded_url" json:"embedded_url,omitempty" yaml:"embedded_url,omitempty" toml:"embedded_url,omitempty"`
	Engine                  string `gorm:"column:engine" json:"engine,omitempty" example:"archlinux" yaml:"engine,omitempty" example:"archlinux" toml:"engine,omitempty" example:"archlinux"`
	FirstPageNum            int    `gorm:"column:first_page_num" json:"first_page_num,omitempty" yaml:"first_page_num,omitempty" toml:"first_page_num,omitempty"`
	GuestClientID           string `gorm:"column:guest_client_id" json:"guest_client_id,omitempty" yaml:"guest_client_id,omitempty" toml:"guest_client_id,omitempty"`
	ImageURLXpath           string `gorm:"column:image_url_xpath" json:"image_url_xpath,omitempty" yaml:"image_url_xpath,omitempty" toml:"image_url_xpath,omitempty"`
	ImagesPath              string `gorm:"column:images_path" json:"images_path,omitempty" yaml:"images_path,omitempty" toml:"images_path,omitempty"`
	ImagesURL               string `gorm:"column:images_url" json:"images_url,omitempty" yaml:"images_url,omitempty" toml:"images_url,omitempty"`
	ImagesXpath             string `gorm:"column:images_xpath" json:"images_xpath,omitempty" yaml:"images_xpath,omitempty" toml:"images_xpath,omitempty"`
	Inactive                bool   `gorm:"column:inactive" json:"inactive,omitempty" yaml:"inactive,omitempty" toml:"inactive,omitempty"`
	LabelXpath              string `gorm:"column:label_xpath" json:"label_xpath,omitempty" yaml:"label_xpath,omitempty" toml:"label_xpath,omitempty"`
	LanguageFallbackXpath   string `gorm:"column:language_fallback_xpath" json:"language_fallback_xpath,omitempty" yaml:"language_fallback_xpath,omitempty" toml:"language_fallback_xpath,omitempty"`
	LanguageSupport         bool   `gorm:"column:language_support" json:"language_support,omitempty" example:"true" yaml:"language_support,omitempty" example:"true" toml:"language_support,omitempty" example:"true"`
	LinkXpath               string `gorm:"column:link_xpath" json:"link_xpath,omitempty" yaml:"link_xpath,omitempty" toml:"link_xpath,omitempty"`
	MagnetXpath             string `gorm:"column:magnet_xpath" json:"magnet_xpath,omitempty" yaml:"magnet_xpath,omitempty" toml:"magnet_xpath,omitempty"`
	MapAddressXpath         string `gorm:"column:map_address_xpath" json:"map_address_xpath,omitempty" yaml:"map_address_xpath,omitempty" toml:"map_address_xpath,omitempty"`
	MapHostnameStart        string `gorm:"column:map_hostname_start" json:"map_hostname_start,omitempty" yaml:"map_hostname_start,omitempty" toml:"map_hostname_start,omitempty"`
	MapNear                 string `gorm:"column:map_near" json:"map_near,omitempty" yaml:"map_near,omitempty" toml:"map_near,omitempty"`
	MapNearPhone            string `gorm:"column:map_near_phone" json:"map_near_phone,omitempty" yaml:"map_near_phone,omitempty" toml:"map_near_phone,omitempty"`
	MapNearTitle            string `gorm:"column:map_near_title" json:"map_near_title,omitempty" yaml:"map_near_title,omitempty" toml:"map_near_title,omitempty"`
	MapNearURL              string `gorm:"column:map_near_url" json:"map_near_url,omitempty" yaml:"map_near_url,omitempty" toml:"map_near_url,omitempty"`
	MapPhoneXpath           string `gorm:"column:map_phone_xpath" json:"map_phone_xpath,omitempty" yaml:"map_phone_xpath,omitempty" toml:"map_phone_xpath,omitempty"`
	MapWebsiteTitleXpath    string `gorm:"column:map_website_title_xpath" json:"map_website_title_xpath,omitempty" yaml:"map_website_title_xpath,omitempty" toml:"map_website_title_xpath,omitempty"`
	MapWebsiteURLXpath      string `gorm:"column:map_website_url_xpath" json:"map_website_url_xpath,omitempty" yaml:"map_website_url_xpath,omitempty" toml:"map_website_url_xpath,omitempty"`
	MapsPath                string `gorm:"column:maps_path" json:"maps_path,omitempty" yaml:"maps_path,omitempty" toml:"maps_path,omitempty"`
	Name                    string `gorm:"column:name" json:"name,omitempty" example:"arch linux wiki" yaml:"name,omitempty" example:"arch linux wiki" toml:"name,omitempty" example:"arch linux wiki"`
	NbPerPage               int    `gorm:"column:nb_per_page" json:"nb_per_page,omitempty" yaml:"nb_per_page,omitempty" toml:"nb_per_page,omitempty"`
	NumberOfResults         int    `gorm:"column:number_of_results" json:"number_of_results,omitempty" yaml:"number_of_results,omitempty" toml:"number_of_results,omitempty"`
	OntentXpath             string `gorm:"column:ontent_xpath" json:"ontent_xpath,omitempty" yaml:"ontent_xpath,omitempty" toml:"ontent_xpath,omitempty"`
	PageSize                int    `gorm:"column:page_size" json:"page_size,omitempty" yaml:"page_size,omitempty" toml:"page_size,omitempty"`
	Paging                  bool   `gorm:"column:paging" json:"paging,omitempty" example:"true" yaml:"paging,omitempty" example:"true" toml:"paging,omitempty" example:"true"`
	PdbeEntryURL            string `gorm:"column:pdbe_entry_url" json:"pdbe_entry_url,omitempty" yaml:"pdbe_entry_url,omitempty" toml:"pdbe_entry_url,omitempty"`
	PdbePreviewURL          string `gorm:"column:pdbe_preview_url" json:"pdbe_preview_url,omitempty" yaml:"pdbe_preview_url,omitempty" toml:"pdbe_preview_url,omitempty"`
	PdbeSolrURL             string `gorm:"column:pdbe_solr_url" json:"pdbe_solr_url,omitempty" yaml:"pdbe_solr_url,omitempty" toml:"pdbe_solr_url,omitempty"`
	PhotoURL                string `gorm:"column:photo_url" json:"photo_url,omitempty" yaml:"photo_url,omitempty" toml:"photo_url,omitempty"`
	PodTitleXpath           string `gorm:"column:pod_title_xpath" json:"pod_title_xpath,omitempty" yaml:"pod_title_xpath,omitempty" toml:"pod_title_xpath,omitempty"`
	PreferredRankXpath      string `gorm:"column:preferred_rank_xpath" json:"preferred_rank_xpath,omitempty" yaml:"preferred_rank_xpath,omitempty" toml:"preferred_rank_xpath,omitempty"`
	PropertyAddress         string `gorm:"column:property_address" json:"property_address,omitempty" yaml:"property_address,omitempty" toml:"property_address,omitempty"`
	PropertyPhone           string `gorm:"column:property_phone" json:"property_phone,omitempty" yaml:"property_phone,omitempty" toml:"property_phone,omitempty"`
	PropertyRowXpath        string `gorm:"column:property_row_xpath" json:"property_row_xpath,omitempty" yaml:"property_row_xpath,omitempty" toml:"property_row_xpath,omitempty"`
	PropertyXpath           string `gorm:"column:property_xpath" json:"property_xpath,omitempty" yaml:"property_xpath,omitempty" toml:"property_xpath,omitempty"`
	PublishedDateXpath      string `gorm:"column:publishedDate_xpath" json:"publishedDate_xpath,omitempty" yaml:"publishedDate_xpath,omitempty" toml:"publishedDate_xpath,omitempty"`
	RedirectPath            string `gorm:"column:redirect_path" json:"redirect_path,omitempty" yaml:"redirect_path,omitempty" toml:"redirect_path,omitempty"`
	ResultBaseURL           string `gorm:"column:result_base_url" json:"result_base_url,omitempty" yaml:"result_base_url,omitempty" toml:"result_base_url,omitempty"`
	ResultURL               string `gorm:"column:result_url" json:"result_url,omitempty" yaml:"result_url,omitempty" toml:"result_url,omitempty"`
	ResultXpath             string `gorm:"column:result_xpath" json:"result_xpath,omitempty" yaml:"result_xpath,omitempty" toml:"result_xpath,omitempty"`
	ResultsQuery            string `gorm:"column:results_query" json:"results_query,omitempty" yaml:"results_query,omitempty" toml:"results_query,omitempty"`
	ResultsXpath            string `gorm:"column:results_xpath" json:"results_xpath,omitempty" yaml:"results_xpath,omitempty" toml:"results_xpath,omitempty"`
	Safesearch              bool   `gorm:"column:safesearch" json:"safesearch,omitempty" yaml:"safesearch,omitempty" toml:"safesearch,omitempty"`
	SearchPath              string `gorm:"column:search_path" json:"search_path,omitempty" yaml:"search_path,omitempty" toml:"search_path,omitempty"`
	SearchPostfix           string `gorm:"column:search_postfix" json:"search_postfix,omitempty" yaml:"search_postfix,omitempty" toml:"search_postfix,omitempty"`
	SearchString            string `gorm:"column:search_string" json:"search_string,omitempty" yaml:"search_string,omitempty" toml:"search_string,omitempty"`
	SearchType              string `gorm:"column:search_type" json:"search_type,omitempty" yaml:"search_type,omitempty" toml:"search_type,omitempty"`
	SearchURL               string `gorm:"column:search_url" json:"search_url,omitempty" yaml:"search_url,omitempty" toml:"search_url,omitempty"`
	SearchURLWithTime       string `gorm:"column:search_url_with_time" json:"search_url_with_time,omitempty" yaml:"search_url_with_time,omitempty" toml:"search_url_with_time,omitempty"`
	Shortcut                string `gorm:"column:shortcut" json:"shortcut,omitempty" example:"al" yaml:"shortcut,omitempty" example:"al" toml:"shortcut,omitempty" example:"al"`
	SiteURL                 string `gorm:"column:site_url" json:"site_url,omitempty" yaml:"site_url,omitempty" toml:"site_url,omitempty"`
	SpellingSuggestionXpath string `gorm:"column:spelling_suggestion_xpath" json:"spelling_suggestion_xpath,omitempty" yaml:"spelling_suggestion_xpath,omitempty" toml:"spelling_suggestion_xpath,omitempty"`
	SrcRemoteURL            string `gorm:"column:src_remote_url" json:"src_remote_url,omitempty" yaml:"src_remote_url,omitempty" toml:"src_remote_url,omitempty"`
	SuggestionXpath         string `gorm:"column:suggestion_xpath" json:"suggestion_xpath,omitempty" yaml:"suggestion_xpath,omitempty" toml:"suggestion_xpath,omitempty"`
	SupportedLanguagesURL   string `gorm:"column:supported_languages_url" json:"supported_languages_url,omitempty" yaml:"supported_languages_url,omitempty" toml:"supported_languages_url,omitempty"`
	ThumbURL                string `gorm:"column:thumb_url" json:"thumb_url,omitempty" yaml:"thumb_url,omitempty" toml:"thumb_url,omitempty"`
	ThumbnailXpath          string `gorm:"column:thumbnail_xpath" json:"thumbnail_xpath,omitempty" yaml:"thumbnail_xpath,omitempty" toml:"thumbnail_xpath,omitempty"`
	TimeRangeCustomAttr     string `gorm:"column:time_range_custom_attr" json:"time_range_custom_attr,omitempty" yaml:"time_range_custom_attr,omitempty" toml:"time_range_custom_attr,omitempty"`
	TimeRangeString         string `gorm:"column:time_range_string" json:"time_range_string,omitempty" yaml:"time_range_string,omitempty" toml:"time_range_string,omitempty"`
	TimeRangeSearch         string `gorm:"column:time_range_search" json:"time_range_search,omitempty" yaml:"time_range_search,omitempty" toml:"time_range_search,omitempty"`
	TimeRangeURL            string `gorm:"column:time_range_url" json:"time_range_url,omitempty" yaml:"time_range_url,omitempty" toml:"time_range_url,omitempty"`
	Timeout                 int    `gorm:"column:timeout" json:"timeout,omitempty" yaml:"timeout,omitempty" toml:"timeout,omitempty"`
	TitleQuery              string `gorm:"column:title_query" json:"title_query,omitempty" yaml:"title_query,omitempty" toml:"title_query,omitempty"`
	TitleXpath              string `gorm:"column:title_xpath" json:"title_xpath,omitempty" yaml:"title_xpath,omitempty" toml:"title_xpath,omitempty"`
	TorrentXpath            string `gorm:"column:torrent_xpath" json:"torrent_xpath,omitempty" yaml:"torrent_xpath,omitempty" toml:"torrent_xpath,omitempty"`
	URL                     string `gorm:"column:url" json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty"`
	URLMap                  string `gorm:"column:url_map" json:"url_map,omitempty" yaml:"url_map,omitempty" toml:"url_map,omitempty"`
	URLQuery                string `gorm:"column:url_query" json:"url_query,omitempty" yaml:"url_query,omitempty" toml:"url_query,omitempty"`
	URLXpath                string `gorm:"column:url_xpath" json:"url_xpath,omitempty" yaml:"url_xpath,omitempty" toml:"url_xpath,omitempty"`
	ValueXpath              string `gorm:"column:value_xpath" json:"value_xpath,omitempty" yaml:"value_xpath,omitempty" toml:"value_xpath,omitempty"`
	Weight                  int    `gorm:"column:weight" json:"weight,omitempty" yaml:"weight,omitempty" toml:"weight,omitempty"`
	WikidataIdsXpath        string `gorm:"column:wikidata_ids_xpath" json:"wikidata_ids_xpath,omitempty" yaml:"wikidata_ids_xpath,omitempty" toml:"wikidata_ids_xpath,omitempty"`
	WikilinkXpath           string `gorm:"column:wikilink_xpath" json:"wikilink_xpath,omitempty" yaml:"wikilink_xpath,omitempty" toml:"wikilink_xpath,omitempty"`
	XpathCategory           string `gorm:"column:xpath_category" json:"xpath_category,omitempty" yaml:"xpath_category,omitempty" toml:"xpath_category,omitempty"`
	XpathDownloads          string `gorm:"column:xpath_downloads" json:"xpath_downloads,omitempty" yaml:"xpath_downloads,omitempty" toml:"xpath_downloads,omitempty"`
	XpathFilesize           string `gorm:"column:xpath_filesize" json:"xpath_filesize,omitempty" yaml:"xpath_filesize,omitempty" toml:"xpath_filesize,omitempty"`
	XpathLeeches            string `gorm:"column:xpath_leeches" json:"xpath_leeches,omitempty" yaml:"xpath_leeches,omitempty" toml:"xpath_leeches,omitempty"`
	XpathResults            string `gorm:"column:xpath_results" json:"xpath_results,omitempty" example:"//ul[@class="mw-search-results"]/li" yaml:"xpath_results,omitempty" example:"//ul[@class="mw-search-results"]/li" toml:"xpath_results,omitempty" example:"//ul[@class="mw-search-results"]/li"`
	XpathSeeds              string `gorm:"column:xpath_seeds" json:"xpath_seeds,omitempty" yaml:"xpath_seeds,omitempty" toml:"xpath_seeds,omitempty"`
	XpathTorrentLinks       string `gorm:"column:xpath_torrent_links" json:"xpath_torrent_links,omitempty" yaml:"xpath_torrent_links,omitempty" toml:"xpath_torrent_links,omitempty"`

	AcceptHeader         []AcceptHeader      `gorm:"many2many:provider_accept_headers;" json:"accept_header,omitempty" yaml:"accept_header,omitempty" toml:"accept_header,omitempty"`
	AcceptHeaderStr      string              `gorm:"column:accept_header_str" json:"accept_header_str,omitempty" yaml:"accept_header_str,omitempty" toml:"accept_header_str,omitempty"`
	Categories           []classify.Category `gorm:"column:categories" json:"categories,omitempty" yaml:"categories,omitempty" toml:"categories,omitempty"`
	CategoriesStr        string              `gorm:"column:categories_str" json:"categories_str,omitempty" yaml:"categories_str,omitempty" toml:"categories_str,omitempty"`
	CategoryToKeyword    []CategoryToKeyword `gorm:"many2many:provider_category_to_keywords;" json:"category_to_keyword,omitempty" yaml:"category_to_keyword,omitempty" toml:"category_to_keyword,omitempty"`
	CategoryToKeywordStr string              `gorm:"column:category_to_keyword_str" json:"category_to_keyword_str,omitempty" yaml:"category_to_keyword_str,omitempty" toml:"category_to_keyword_str,omitempty"`
	CountryToHostname    []CountryToHostname `gorm:"many2many:provider_country_to_hostnames;" json:"country_to_hostname,omitempty" yaml:"country_to_hostname,omitempty" toml:"country_to_hostname,omitempty"`
	CountryToHostnameStr string              `gorm:"column:country_to_hostname_str" json:"country_to_hostname_str,omitempty" yaml:"country_to_hostname_str,omitempty" toml:"country_to_hostname_str,omitempty"`
	ImageSizes           []ImageSize         `gorm:"many2many:provider_image_sizes;" json:"image_sizes,omitempty" yaml:"image_sizes,omitempty" toml:"image_sizes,omitempty"`
	ImageSizesStr        string              `gorm:"column:image_sizes_str" json:"image_sizes_str,omitempty" yaml:"image_sizes_str,omitempty" toml:"image_sizes_str,omitempty"`
	SafeSearchTypes      []SafeSearchType    `gorm:"many2many:provider_safe_search_types;" json:"safesearch_types,omitempty" yaml:"safesearch_types,omitempty" toml:"safesearch_types,omitempty"`
	SafesearchTypesStr   string              `gorm:"column:safesearch_types_str" json:"safesearch_types_str,omitempty" yaml:"safesearch_types_str,omitempty" toml:"safesearch_types_str,omitempty"`
	SearchCategory       []SearchCategory    `gorm:"many2many:provider_search_categories;" json:"search_category,omitempty" yaml:"search_category,omitempty" toml:"search_category,omitempty"`
	SearchCategoryStr    string              `gorm:"column:search_category_str" json:"search_category_str,omitempty" yaml:"search_category_str,omitempty" toml:"search_category_str,omitempty"`
	ShorcutDict          []Shortcut          `gorm:"many2many:provider_shortcuts;" json:"shorcut_dict,omitempty" yaml:"shorcut_dict,omitempty" toml:"shorcut_dict,omitempty"`
	ShorcutDictStr       string              `gorm:"column:shorcut_dict_str" json:"shorcut_dict_str,omitempty" yaml:"shorcut_dict_str,omitempty" toml:"shorcut_dict_str,omitempty"`
	TimeRangeDict        []TimeRange         `gorm:"many2many:provider_time_ranges;" json:"time_range_dict,omitempty" yaml:"time_range_dict,omitempty" toml:"time_range_dict,omitempty"`
	TimeRangeDictStr     string              `gorm:"column:time_range_dict_str" json:"time_range_dict_str,omitempty" yaml:"time_range_dict_str,omitempty" toml:"time_range_dict_str,omitempty"`

	Extract Nodes `gorm:"many2many:provider_nodes;" json:"extract,omitempty" yaml:"extract,omitempty" toml:"extract,omitempty"`
	URLS    []URL `gorm:"many2many:provider_urls;" json:"urls,omitempty" yaml:"urls,omitempty" toml:"urls,omitempty"`
}

type URL struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	IsAPI   bool   `gorm:"column:is_api" default:"false" storm:"is_api" json:"is_api" yaml:"is_api" toml:"is_api"`
	Lang    string `gorm:"column:lang" default:"all" storm:"lang" json:"lang" yaml:"lang" toml:"lang"`
	Url     string `gorm:"column:url" storm:"url" json:"url" yaml:"url" toml:"url"`
	Base    string `gorm:"column:base" storm:"base" json:"base" yaml:"base" toml:"base"`
	Search  string `gorm:"column:search" storm:"search" json:"search" yaml:"search" toml:"search"`
	Item    string `gorm:"column:item" storm:"item" json:"item" yaml:"item" toml:"item"`
	Sitemap string `gorm:"column:sitemap" storm:"sitemap" json:"sitemap" yaml:"sitemap" toml:"sitemap"`
}

type Nodes struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Active   int    `gorm:"column:active" default:"1" example:"2" storm:"active" json:"active" yaml:"active" toml:"active"`
	Encoding string `gorm:"column:encoding" storm:"encoding" json:"encoding" yaml:"encoding" toml:"encoding"`

	Leafs []Path `gorm:"many2many:provider_path_leafs;" storm:"leafs" json:"leafs" yaml:"leafs" toml:"leafs"`
	Paths []Path `gorm:"many2many:provider_paths;"" storm:"paths" json:"paths" yaml:"paths" toml:"paths"`
}

type Path struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Name       string `gorm:"column:name" json:"name" yaml:"name" toml:"name"`
	Path       string `gorm:"column:path" json:"path" yaml:"path" toml:"path"`
}

type PathBlocks struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Article             string            `gorm:"column:article" json:"article,omitempty" yaml:"article,omitempty" toml:"article,omitempty"`
	Author              string            `gorm:"column:author" json:"author,omitempty" yaml:"author,omitempty" toml:"author,omitempty"`
	Content             string            `gorm:"column:content" json:"content,omitempty" yaml:"content,omitempty" toml:"content,omitempty"`
	ContentMisc         string            `gorm:"column:content_misc" json:"content_misc,omitempty" yaml:"content_misc,omitempty" toml:"content_misc,omitempty"`
	RelatedLinks        string            `gorm:"column:related_links" json:"related_links,omitempty" yaml:"related_links,omitempty" toml:"related_links,omitempty"`
	Results             string            `gorm:"column:results" json:"results,omitempty" yaml:"results,omitempty" toml:"results,omitempty"`
	SpellingSuggestions string            `gorm:"column:spelling_suggestions" json:"spelling_suggestions,omitempty" yaml:"spelling_suggestions,omitempty" toml:"spelling_suggestions,omitempty"`
	Suggestions         string            `gorm:"column:suggestions" json:"suggestions,omitempty" yaml:"suggestions,omitempty" toml:"suggestions,omitempty"`
	Tags                string            `gorm:"column:tags" json:"tags,omitempty" yaml:"tags,omitempty" toml:"tags,omitempty"`
	Title               string            `gorm:"column:title" json:"title,omitempty" yaml:"title,omitempty" toml:"title,omitempty"`
	Thumbnail           string            `gorm:"column:thumbnail" json:"thumbnail,omitempty" yaml:"thumbnail,omitempty" toml:"thumbnail,omitempty"`
	URL                 string            `gorm:"column:url" json:"url,omitempty" yaml:"url,omitempty" toml:"url,omitempty"`
	Extra               []PathExtraBlocks `gorm:"many2many:provider_path_extras;" json:"extra,omitempty" yaml:"extra,omitempty" toml:"extra,omitempty"`
}

type PathExtraBlocks struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool   `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Name     string `gorm:"column:name" json:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Path     string `gorm:"column:path" json:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
}

type QueryPattern struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	LimitString  string `gorm:"column:limit_str" storm:"limit_str" json:"limit_str" yaml:"limit_str" toml:"limit_str"`
	PageString   string `gorm:"column:page_string" storm:"page_string" json:"page_string" yaml:"page_string" toml:"page_string"`
	ExpandString string `gorm:"column:expand_string" storm:"expand_string" json:"expand_string" yaml:"expand_string" toml:"expand_string"`
	FieldsString string `gorm:"column:fields_string" storm:"fields_string" json:"fields_string" yaml:"fields_string" toml:"fields_string"`
	OrderString  string `gorm:"column:order_string" storm:"order_string" json:"order_string" yaml:"order_string" toml:"order_string"`
	AscString    string `gorm:"column:asc_string" storm:"asc_string" json:"asc_string" yaml:"asc_string" toml:"asc_string"`
	DescString   string `gorm:"column:desc_string" storm:"desc_string" json:"desc_string" yaml:"desc_string" toml:"desc_string"`
	QueryString  string `gorm:"column:query_string" storm:"query_string" json:"query_string" yaml:"query_string" toml:"query_string"`
	ParamString  string `gorm:"column:param_string" storm:"param_string" json:"param_string" yaml:"param_string" toml:"param_string"`
	LimitValue   int    `gorm:"column:limit_value" storm:"limit_value" json:"limit_value" yaml:"limit_value" toml:"limit_value"`
	PageValue    int    `gorm:"column:page_value" storm:"page_value" json:"page_value" yaml:"page_value" toml:"page_value"`
	LeftBracket  rune   `gorm:"column:left_bracket" storm:"left_bracket" json:"left_bracket" yaml:"left_bracket" toml:"left_bracket"`
	RightBracket rune   `gorm:"column:right_bracket" storm:"right_bracket" json:"right_bracket" yaml:"right_bracket" toml:"right_bracket"`
	Separator    rune   `gorm:"column:separator" storm:"separator" json:"separator" yaml:"separator" toml:"separator"`
	KVSeparator  rune   `gorm:"column:kv_sperator" storm:"kv_sperator" json:"kv_sperator" yaml:"kv_sperator" toml:"kv_sperator"`
}

type SafeSearchTypes struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool             `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool             `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	List       []SafeSearchType `gorm:"many2many:provider_safe_search_types;" json:"safesearch_types,omitempty" yaml:"safesearch_types,omitempty" toml:"safesearch_types,omitempty"`
}

type SafeSearchType struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Key        string `gorm:"column:key" json:"key" yaml:"key" toml:"key"`
	Value      string `gorm:"column:value" json:"value" yaml:"value" toml:"value"`
}

type TimeRanges struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool        `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool        `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	List       []TimeRange `gorm:"many2many:provider_time_ranges;" json:"time_ranges,omitempty" yaml:"time_ranges,omitempty" toml:"time_ranges,omitempty"`
}

type TimeRange struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Key        string `gorm:"column:key" json:"key" yaml:"key" toml:"key"`
	Value      string `gorm:"column:value" json:"value" yaml:"value" toml:"value"`
}

type ImageSizes struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool        `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool        `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	List       []ImageSize `gorm:"many2many:provider_image_sizes;" json:"image_sizes,omitempty" yaml:"image_sizes,omitempty" toml:"image_sizes,omitempty"`
}

type ImageSize struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Label      string `gorm:"column:label" json:"label" yaml:"label" toml:"label"`
	Key        string `gorm:"column:key" json:"key" yaml:"key" toml:"key"`
	Value      string `gorm:"column:value" json:"value" yaml:"value" toml:"value"`
}

type AcceptHeaders struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool           `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool           `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	List       []AcceptHeader `gorm:"many2many:provider_accept_headers;" json:"accept_headers,omitempty" yaml:"accept_headers,omitempty" toml:"accept_headers,omitempty"`
}

type AcceptHeader struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Value      string `gorm:"column:value" json:"value" yaml:"value" toml:"value"`
}

type CountryToHostnames struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool                `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool                `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	List       []CountryToHostname `gorm:"many2many:provider_country_to_hostnames;" json:"country_to_hostnames,omitempty" yaml:"country_to_hostnames,omitempty" toml:"country_to_hostnames,omitempty"`
}

type CountryToHostname struct {
	gorm.Model  `json:"-" yaml:"-" toml:"-"`
	Disabled    bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug       bool   `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	CountryName string `gorm:"column:country_name" json:"country_name" yaml:"country_name" toml:"country_name"`
	CountryIso2 string `gorm:"column:country_iso2" json:"country_iso2" yaml:"country_iso2" toml:"country_iso2"`
	CountryIso3 string `gorm:"column:country_iso3" json:"country_iso3" yaml:"country_iso3" toml:"country_iso3"`
	Hostname    string `gorm:"column:hostname" json:"hostname" yaml:"hostname" toml:"hostname"`
	TLD         string `gorm:"column:tld" json:"tld" yaml:"tld" toml:"tld"`
	Proxy       string `gorm:"column:proxy" json:"proxy" yaml:"proxy" toml:"proxy"`
}

type SearchCategories struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool             `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool             `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	List       []SearchCategory `gorm:"many2many:provider_search_categories;" json:"category_to_keywords,omitempty" yaml:"category_to_keywords,omitempty" toml:"category_to_keywords,omitempty"`
}

type SearchCategory struct {
	gorm.Model   `json:"-" yaml:"-" toml:"-"`
	Disabled     bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug        bool   `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Keyword      string `gorm:"column:keyword" json:"keyword" yaml:"keyword" toml:"keyword"`
	CategoryName string `gorm:"column:category_name" json:"category_name" yaml:"category_name" toml:"category_name"`
}

type CategoryToKeywords struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool                `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool                `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	List       []CategoryToKeyword `gorm:"many2many:provider_category_to_keywords;" json:"category_to_keywords,omitempty" yaml:"category_to_keywords,omitempty" toml:"category_to_keywords,omitempty"`
}

type CategoryToKeyword struct {
	gorm.Model   `json:"-" yaml:"-" toml:"-"`
	Disabled     bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug        bool   `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	CategoryName string `gorm:"column:category_name" json:"category_name" yaml:"category_name" toml:"category_name"`
	Keyword      string `gorm:"column:keyword" json:"keyword" yaml:"keyword" toml:"keyword"`
}

type Shortcuts struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool       `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool       `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Shortcuts  []Shortcut `gorm:"many2many:provider_shortcuts;" json:"shortucts,omitempty" yaml:"shortucts,omitempty" toml:"shortucts,omitempty"`
}

type Shortcut struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `gorm:"column:column" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Shortcut   string `gorm:"column:shortcut" json:"shortcut" yaml:"shortcut" toml:"shortcut"`
	Input      string `gorm:"column:input" json:"input" yaml:"input" toml:"input"`
}
