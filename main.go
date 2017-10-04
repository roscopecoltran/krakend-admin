package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qor/help"
	i18n_database "github.com/qor/i18n/backends/database"
	"github.com/qor/media/asset_manager"

	// "gopkg.in/olivere/elastic.v3"
	// "github.com/imdario/mergo"
	// "github.com/gin-gonic/contrib/cache"
	// "github.com/gin-contrib/secure"
	// "github.com/gregjones/httpcache"
	// "github.com/gin-contrib/static"

	// "github.com/qor/slug"
	"github.com/Machiel/slugify"
	"github.com/boltdb/bolt"
	"github.com/husarlabs/qparser"
	"github.com/k0kubun/pp"
	"github.com/roscopecoltran/configor"
	"github.com/roscopecoltran/krakend-admin/backend"
	"github.com/roscopecoltran/krakend-admin/models"
	// "github.com/roscopecoltran/krakend-admin/models/api/openapi"
	// "github.com/roscopecoltran/krakend-admin/models/api/raml"
)

const (
	_sep = string(os.PathSeparator)
)

var (
	Cfg          Config
	strSlugifier *slugify.Slugifier

	isSanitizeUrl = flag.Bool("sanitize-download", false, "sanitize downloaded file names.")
	downloadCsv   = flag.String("download-csv", "./shared/public/downloads/urls.csv", "download urks csv file")
	downloadDir   = flag.String("download-dir", "./shared/specs/registry", "download prefix path")

	dbPath      = flag.String("db", "./shared/stores/krakend.boltdb", "database absolute/relative path")
	ramlFile    = flag.String("raml", "./shared/testdata/specs/raml/github-api-v3.yaml", "location of the swagger spec file")
	swaggerFile = flag.String("swagger", "./shared/testdata/specs/swagger/2.0/github-api-v3.yaml", "location of the swagger spec file")
	annotate    = flag.Bool("options", true, "include (google.api.http) options for grpc-gateway")

	definitionsDir = flag.String("d", "./shared/specs/definitions", "filepath to the definitions yaml files")
	pathsDir       = flag.String("p", "./shared/specs/paths", "filepath to the paths yaml files")
	responsesDir   = flag.String("r", "./shared/specs/responses", "filepath to the responses yaml files")
	format         = flag.String("f", "yaml", "the file format to print to stdout. can be yaml or json")
)

var defaultConfigFiles = []string{
	"shared/testdata/service.yaml",
	"shared/testdata/downloads.yaml",
	"shared/testdata/specs.yaml",
	"shared/testdata/file_types.yaml",
	"shared/testdata/engines.yaml",
	"shared/testdata/specs/api-gurus.yaml",
	"shared/testdata/providers.yaml",
}

var defaultProviderModels = []interface{}{
	&Provider{}, &ProviderCategoryToKeyword{}, &ProviderCountryToHostname{},
	&ProviderQueryPattern{}, &ProviderPathBlocks{}, &ProviderPathBlocks{}, &ProviderPathExtraBlocks{},
}

var defaultLocalConfigModels = []interface{}{
	&LocalFile{}, &LocalEnvironment{}, &LocalEnvironmentVariable{},
}

var defaultClassifyModels = []interface{}{
	&models.ClassifyCategory{}, &models.ClassifyTag{}, &models.ClassifyTopic{},
}

var defaultTaskQModels = []interface{}{
	&TaskQueue{}, &TaskQueueRunner{},
}

var defaultSpecsModels = []interface{}{

	&SpecsConfig{}, &SpecsRegistryAPIVersion{}, &SpecsRegistryAPIVersionInfo{}, &SpecsRegistryAPIVersionInfoContact{}, &SpecsRegistryAPIVersionInfoXOrigin{}, &SpecsRegistryAPIVersionInfoXLogo{},
}

var defaultGatewayModels = []interface{}{
	&GatewayBackend{}, &GatewayEndpoint{}, &GatewayService{},
	&GatewayHeaders{}, &GatewayQueryStrings{}, &GatewayParameters{},
	&GatewayBlacklist{}, &GatewayWhitelist{}, &GatewayUrlKeys{},
	&GatewayHolders{}, &GatewayMappings{},
}

var defaultModels = []interface{}{

	&i18n_database.Translation{}, &help.QorHelpEntry{}, &asset_manager.AssetManager{},

	&SpecsConfig{}, &SpecsRegistry{},
	// &SpecsRegistryAPIVersion{}, &SpecsRegistryAPIVersionInfo{}, &SpecsRegistryAPIVersionInfoContact{}, &SpecsRegistryAPIVersionInfoXOrigin{}, &SpecsRegistryAPIVersionInfoXLogo{},

	&LocalFile{}, &LocalEnvironment{}, &LocalEnvironmentVariable{},
	&CommonCountry{}, &CommonCurrency{},
	&models.ClassifyCategory{}, &models.ClassifyTag{}, &models.ClassifyTopic{},
	&Provider{}, &ProviderCategoryToKeyword{}, &ProviderCountryToHostname{}, &ProviderQueryPattern{}, &ProviderPath{}, &ProviderPath{}, &ProviderPathExtraBlocks{}, &ProviderURL{},

	&GatewayBackend{}, &GatewayEndpoint{}, &GatewayService{},
	&GatewayHeaders{}, &GatewayQueryStrings{}, &GatewayParameters{}, &GatewayBlacklist{}, &GatewayWhitelist{},
	&GatewayUrlKeys{}, &GatewayHolders{}, &GatewayMappings{},
	&TaskQueue{}, &TaskQueueRunner{},
}

func init() {
	strSlugifier = slugify.New(slugify.Configuration{ReplaceCharacter: '-'})
}

func createURLs(u []ProviderURL) error {
	for _, url := range u {
		if ok := backend.Store.Gorm.DB.NewRecord(url); ok {
			if err := backend.Store.Gorm.DB.Create(&url).Error; err != nil {
				fmt.Println("error: ", err)
			}
		}
	}
	return nil
}

// Cfg.Specs.Registry
func createNewApiGuruProfile(apis []SpecsRegistryAPI) error {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(fmt.Sprintf("%v", *downloadCsv), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	for k, api := range apis {
		if ok := backend.Store.Gorm.DB.NewRecord(api); ok {
			pp.Print(api)
			if err := backend.Store.Gorm.DB.Create(&api).Error; err != nil {
				fmt.Println("error: ", err)
			}
			var contentURL string
			for m, v := range api.Versions {
				contentURL = fmt.Sprintf("%s\n", v.SwaggerURL)
				if _, err := f.Write([]byte(contentURL)); err != nil {
					fmt.Println("error: ", err)
				}
				contentURL = fmt.Sprintf("%s\n", v.SwaggerYamlURL)
				if _, err := f.Write([]byte(contentURL)); err != nil {
					fmt.Println("error: ", err)
				}
				fmt.Println("api key: ", k, "version key: ", m)
				fileName, fullPath, fullName := getFilePathFromURL(contentURL)
				fmt.Println(" |- fileName: ", fileName, ", fullPath: ", fullPath, ", fullName: ", fullName)
				filePathAbsolute, err := filepath.Abs(fullName)
				if err != nil {
					continue
				}
				fmt.Printf("filePathAbsolute:", filePathAbsolute)
				// loadSwaggerFile(filePathAbsolute)
				// openapi2Protobuf(filePathAbsolute, fullPath, true)
				//if err := configor.Load(&Cfg.Specs.Registry[k].Versions[m].OpenAPI, filePathAbsolute); err != nil {
				//	log.Fatal("ERROR while loading the config struct:", err.Error())
				//}
			}
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func createCategories(c []models.ClassifyCategory) error {
	for _, cat := range c {
		category := models.ClassifyCategory{}
		category.Name = cat.Name
		category.Code = strings.ToLower(cat.Name)
		category.Slug = fmt.Sprintf(strSlugifier.Slugify(cat.Name))
		if ok := backend.Store.Gorm.DB.NewRecord(category); ok {
			if err := backend.Store.Gorm.DB.Create(&category).Error; err != nil {
				fmt.Println("error: ", err)
				// return err
			}
		}
	}
	return nil
}

func importNewProvider(p Provider) error {
	pp.Print(p)
	if ok := backend.Store.Gorm.DB.NewRecord(p); ok {
		backend.Store.Gorm.DB.Create(&p)
	}
	if len(p.Categories) > 0 {
		createCategories(p.Categories)
	}
	if len(p.URLS) > 0 {
		createURLs(p.URLS)
	}
	return nil
}

func dumpConfig() {
	if err := configor.Dump(Cfg.Krakend, "krakend", "yaml,json", "./shared/config/dump"); err != nil {
		log.Fatal("ERROR while dumping the config struct:", err.Error())
	}
	if err := configor.Dump(Cfg.Service, "service", "yaml,json", "./shared/config/dump"); err != nil {
		log.Fatal("ERROR while dumping the config struct:", err.Error())
	}

	if err := configor.Dump(Cfg.Providers, "providers", "yaml,json", "./shared/config/dump"); err != nil {
		log.Fatal("ERROR while dumping the config struct:", err.Error())
	}

	if err := configor.Dump(Cfg, "all", "yaml,json", "./shared/config/dump"); err != nil {
		log.Fatal("ERROR while dumping the config struct:", err.Error())
	}
}

func main() {

	if err := configor.Load(&Cfg, defaultConfigFiles...); err != nil {
		log.Fatal("ERROR while loading the config struct:", err.Error())
	}

	flag.Parse()

	f, err := os.OpenFile(fmt.Sprintf("%s", *downloadCsv), os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	f.Truncate(0)
	f.Seek(0, 0)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	backend.Store.Admin.Disabled = false
	backend.Store.Gorm.IsTruncate = true
	backend.Store.Gorm.IsAutoMigrate = true

	db, err := bolt.Open(*dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	backend.Store.Bolt.DB = db
	defer db.Close()

	if err := backend.New(defaultModels...); err != nil {
		log.Fatal("ERROR while init backend components:", err.Error())
	}

	for _, p := range Cfg.Providers {
		krakenBackend, gormBackend := convert(p)
		Cfg.Backends = append(Cfg.Backends, krakenBackend)
		if err := importNewProvider(p); err != nil {
			fmt.Println("error occured while importing provider: ", p.Name)
		}
		if ok := backend.Store.Gorm.DB.NewRecord(gormBackend); ok {
			backend.Store.Gorm.DB.Create(&gormBackend)
		}
	}

	// Create a wait group to report to and a work channel that takes URLs
	wg := new(sync.WaitGroup)
	work := make(chan string)

	// Create workers and start goroutines listening to work channel
	numWorkers := 25
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(work, wg)
	}

	if len(Cfg.Specs.Registry) > 0 {
		if err := createNewApiGuruProfile(Cfg.Specs.Registry); err != nil {
			fmt.Println("error occured while importing apis-guru profile.")
		}
	}

	// Add URLs to work channel from CSV
	addWorkFromCsv(work, fmt.Sprintf("%v", *downloadCsv))

	/*
		parseRAML(*ramlFile)
		openapiSplit(*swaggerFile, *format)
	*/

	if ok := backend.Store.Gorm.DB.NewRecord(Cfg.Krakend); ok {
		backend.Store.Gorm.DB.Create(&Cfg.Krakend)
	}

	for _, e := range Cfg.Krakend.Endpoints {
		if ok := backend.Store.Gorm.DB.NewRecord(e); ok {
			backend.Store.Gorm.DB.Create(&e)
		}
		for _, b := range e.Backend {
			if ok := backend.Store.Gorm.DB.NewRecord(b); ok {
				backend.Store.Gorm.DB.Create(&b)
			}
		}
	}

	// Wait for workers to finish
	close(work)
	wg.Wait()

	if len(Cfg.Specs.Registry) > 0 {
		parseSpecsOpenAPI(Cfg.Specs.Registry)
	}

	fmt.Println("Done.")
	dumpConfig()

	mux := http.NewServeMux()
	if !backend.Store.Admin.Disabled {

		backend.Store.Admin.UI.SetSiteName("SniperKit Gateway - Administration Panel")
		//backend.Store.Admin.UI.SetAuth(auth.AdminAuth{})

		// Add Dashboard
		backend.Store.Admin.UI.MountTo("/admin", mux)
	}

	r := gin.Default()
	if !backend.Store.Admin.Disabled {
		r.Any("/admin/*w", gin.WrapH(mux))
	}
	r.Run()

}

func replaceSlugify(input string, replacementMap map[rune]string) string {
	replacementMap = map[rune]string{
		'a': "hello",
		'b': "hi",
	}
	replacementMapSlugifier := slugify.New(slugify.Configuration{
		ReplacementMap: replacementMap,
	})
	output := replacementMapSlugifier.Slugify(input)
	return output
}

func parseRawFile(url string, provider string, patterns map[string]string) (map[string]string, error) {
	keys := make(map[string]string)
	if provider == "" {
		return keys, errors.New("No provider name provided.")
	}

	if url == "" {
		url = fmt.Sprintf("https://raw.githubusercontent.com/ascimoo/searx/blob/%s/searx/engines/%s.py", provider, provider)
	}
	// src_remote_url: https://raw.githubusercontent.com/jibe-b/searx/blob/arxiv/searx/engines/arxiv.py
	return keys, nil
}

func convert(p Provider) (BackendExport, GatewayBackend) {

	dotSlugifier := slugify.New(slugify.Configuration{ReplaceCharacter: '-'})
	fmt.Println(dotSlugifier.Slugify("Hello, world!")) // Will print: hello.world

	paths := make([]map[string]string, 0)
	//paths = append(paths, p.TagsXpath)

	holders := make([]map[string]string, 0)
	mapping := make(map[string]string)
	parameters := make(map[string]string)
	query_strings := make(map[string]string)

	url_pattern := ""

	opts := &qparser.ParserOptions{
		LimitString: "l",
		PageString:  "pg",
	}

	parser := qparser.NewParser(opts)
	res, err := parser.ParseString(p.BaseURL)
	if err != nil {
		fmt.Println("res: ", res)
		fmt.Println("error: ", err)
	}

	// pp.Print(res)
	slug := fmt.Sprintf("%s %s %s", p.Name, p.BaseURL, p.Engine)

	gormBackend := GatewayBackend{
		Slug:                     dotSlugifier.Slugify(slug),
		Group:                    dotSlugifier.Slugify(p.Name),
		Method:                   "GET",
		HostStr:                  p.BaseURL,
		HostSanitizationDisabled: true,
		URLPattern:               url_pattern,
		HeaderStr:                "",
		QueryStringsStr:          "",
		ParametersStr:            "",
		BlacklistStr:             "",
		WhitelistStr:             "",
		// PathsStr:                 paths,
		Timeout:         time.Duration(p.Timeout) * time.Second,
		MappingStr:      "",
		HoldersStr:      "",
		Encoding:        p.Engine,
		URLKeysStr:      "",
		ConcurrentCalls: 1,
	}

	krakenBackend := BackendExport{
		Slug:   dotSlugifier.Slugify(slug),
		Group:  dotSlugifier.Slugify(p.Name),
		Method: "GET",
		Host:   []string{p.BaseURL},
		HostSanitizationDisabled: true,
		URLPattern:               url_pattern,
		Header:                   []string{},
		QueryStrings:             query_strings,
		Parameters:               parameters,
		Blacklist:                []string{},
		Whitelist:                []string{},
		Paths:                    paths,
		Timeout:                  time.Duration(p.Timeout) * time.Second,
		Mapping:                  mapping,
		Holders:                  holders,
		Encoding:                 p.Engine,
		URLKeys:                  []string{},
		ConcurrentCalls:          1,
	}

	return krakenBackend, gormBackend

}
