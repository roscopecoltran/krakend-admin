package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Machiel/slugify"
	"github.com/husarlabs/qparser"
	"github.com/k0kubun/pp"

	"github.com/roscopecoltran/krakend-admin/db"
	"github.com/roscopecoltran/krakend-admin/models/classify"
	"github.com/roscopecoltran/krakend-admin/models/provider"
	"github.com/roscopecoltran/krakend-admin/models/provider/engine/gateway"
	"github.com/roscopecoltran/krakend-admin/models/provider/restful"
	"github.com/roscopecoltran/krakend-admin/utils/download"
)

func convert(p provider.Provider) (gateway.Export, gateway.Backend) {

	dotSlugifier := slugify.New(slugify.Configuration{ReplaceCharacter: '-'})
	fmt.Println(dotSlugifier.Slugify("Hello, world!")) // Will print: hello.world

	paths := make([]map[string]string, 0)
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

	gormBackend := gateway.Backend{
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

	krakenBackend := gateway.Export{
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

func createURLs(u []provider.URL) error {
	for _, url := range u {
		if ok := db.Store.Gorm.DB.NewRecord(url); ok {
			if err := db.Store.Gorm.DB.Create(&url).Error; err != nil {
				fmt.Println("error: ", err)
			}
		}
	}
	return nil
}

// Cfg.Specs.Registry
func createNewApiGuruProfile(apis []restful.SpecsRegistryAPI) error {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(fmt.Sprintf("%v", *downloadCsv), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	for k, api := range apis {
		if ok := db.Store.Gorm.DB.NewRecord(api); ok {
			// pp.Print(api)
			if err := db.Store.Gorm.DB.Create(&api).Error; err != nil {
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
				fileName, fullPath, fullName := download.GetFilePathFromURL(contentURL)
				fmt.Println(" |- fileName: ", fileName, ", fullPath: ", fullPath, ", fullName: ", fullName)
				filePathAbsolute, err := filepath.Abs(fullName)
				if err != nil {
					continue
				}
				fmt.Printf("filePathAbsolute: %s \n", filePathAbsolute)
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

func createCategories(c []classify.Category) error {
	for _, cat := range c {
		category := classify.Category{}
		category.Name = cat.Name
		category.Code = strings.ToLower(cat.Name)
		category.Slug = fmt.Sprintf(strSlugifier.Slugify(cat.Name))
		if ok := db.Store.Gorm.DB.NewRecord(category); ok {
			if err := db.Store.Gorm.DB.Create(&category).Error; err != nil {
				fmt.Println("error: ", err)
				// return err
			}
		}
	}
	return nil
}

func importNewProvider(p provider.Provider) error {
	pp.Print(p)
	if ok := db.Store.Gorm.DB.NewRecord(p); ok {
		db.Store.Gorm.DB.Create(&p)
	}
	if len(p.Categories) > 0 {
		createCategories(p.Categories)
	}
	if len(p.URLS) > 0 {
		createURLs(p.URLS)
	}
	return nil
}
