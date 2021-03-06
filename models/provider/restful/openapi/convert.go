package openapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Machiel/slugify"
	"github.com/dimiro1/flatmap"
	"github.com/gebv/typed"
	"github.com/k0kubun/pp"
	"github.com/paulvollmer/yaml2json/src"
	"github.com/pkg4go/convert"
	"github.com/roscopecoltran/configor"
	"github.com/roscopecoltran/krakend-admin/models/provider/engine/gateway"
	"github.com/roscopecoltran/krakend-admin/utils/errors"
	"github.com/roscopecoltran/krakend-admin/utils/types"
	// "github.com/roscopecoltran/krakend-admin/plugins/tokenizer"
	// "github.com/generaltso/linguist/tokenizer"
	// "github.com/emirpasic/gods/utils"
	// "github.com/SamuelKupferschmid/nbc"
	// "github.com/AlasdairF/Classifier"
	// "github.com/lytics/multibayes"
	// "github.com/mgudipati/urlclassify"
	// "github.com/sjwhitworth/golearn/base"
	// "github.com/sjwhitworth/golearn/evaluation"
	// "github.com/sjwhitworth/golearn/knn"
	// "github.com/russellsimpkins/ensemble"
)

type newKrakenBackends []gateway.Export

type SwaggerExports struct {
	Backends []SwaggerExport `json:"backends" yaml:"backends" toml:"backends"`
}

type SwaggerExport struct {
	AttrsCount        int               `default:"0" json:"-" yaml:"-" toml:"-"`                               // to remove
	Group             string            `json:"group,omitempty" yaml:"group,omitempty" toml:"group,omitempty"` // provider name, eg: github.com, bitbucket, aws.amazon
	Scheme            string            `json:"-" yaml:"-" toml:"-"`                                           // `json:"scheme,omitempty" yaml:"scheme,omitempty" toml:"scheme,omitempty"`
	UrlPattern        string            `required:"true" json:"url_pattern" yaml:"url_pattern" toml:"url_pattern"`
	Method            string            `default:"GET" json:"method,omitempty" yaml:"method,omitempty" toml:"method,omitempty"`
	Encoding          string            `default:"json" json:"encoding,omitempty" yaml:"encoding,omitempty" toml:"encoding,omitempty"`
	Host              []string          `json:"host,omitempty" yaml:"host,omitempty" toml:"host,omitempty"`
	Holders           map[string]string `json:"holders,omitempty" yaml:"holders,omitempty" toml:"holders,omitempty"`
	URLKeys           []string          `json:"url_keys,omitempty" yaml:"url_keys,omitempty" toml:"url_keys,omitempty"`
	QueryStringParams []string          `json:"query_string_params,omitempty" yaml:"query_string_params,omitempty" toml:"query_string_params,omitempty"`
	Params            []string          `json:"params,omitempty" yaml:"params,omitempty" toml:"params,omitempty"`
	Headers           []string          `json:"header,omitempty" yaml:"header,omitempty" toml:"header,omitempty"`
	Whitelist         []string          `json:"whitelist,omitempty" yaml:"whitelist,omitempty" toml:"whitelist,omitempty"`
	Blacklist         []string          `json:"blacklist,omitempty" yaml:"blacklist,omitempty" toml:"blacklist,omitempty"`
	IsCollection      bool              `json:"is_collection,omitempty" yaml:"is_collection,omitempty" toml:"is_collection,omitempty"`
	Target            string            `json:"target,omitempty" yaml:"target,omitempty" toml:"target,omitempty"`
	RespAttrCount     int               `default:"0" json:"-" yaml:"-" toml:"-"` // to remove
	RespAttrRef       string            `json:"-" yaml:"-" toml:"-"`
	RespAttrRefPath   string            `json:"-" yaml:"-" toml:"-"`
	ConcurrentCalls   int               `default:"1" json:"concurrent_calls" yaml:"concurrent_calls" toml:"concurrent_calls"`
	Mapping           map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty" toml:"mapping,omitempty"`
	Extra             struct {
		Provider    string            `json:"provider,omitempty" yaml:"provider,omitempty" toml:"provider,omitempty"` // x-providerName
		Description string            `json:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
		Version     string            `json:"version,omitempty" yaml:"version,omitempty" toml:"version,omitempty"`
		Slug        string            `json:"slug,omitempty" yaml:"slug,omitempty" toml:"slug,omitempty"`
		Schemes     []string          `json:"schemes,omitempty" yaml:"schemes,omitempty" toml:"schemes,omitempty"`
		Bodies      []string          `json:"body,omitempty" yaml:"body,omitempty" toml:"body,omitempty"`
		Encodings   []string          `json:"-" yaml:"-" toml:"-"`
		FormDatas   []string          `json:"form_data,omitempty" yaml:"form_data,omitempty" toml:"form_data,omitempty"`
		Produce     string            `json:"-" yaml:"-" toml:"-"`
		Consumes    []string          `json:"consumes,omitempty" yaml:"consumes,omitempty" toml:"consumes,omitempty"`
		Produces    []string          `json:"produces,omitempty" yaml:"produces,omitempty" toml:"produces,omitempty"`
		Consume     string            `json:"-" yaml:"-" toml:"-"`
		Mapping     map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty" toml:"mapping,omitempty"`
		Hosts       []string          `json:"hosts,omitempty" yaml:"hosts,omitempty" toml:"hosts,omitempty"`
		Topics      []string          `json:"topics,omitempty" yaml:"topics,omitempty" toml:"topics,omitempty"`
		Tokens      []string          `json:"tokens,omitempty" yaml:"tokens,omitempty" toml:"tokens,omitempty"`
		LocalFile   string            `json:"-" yaml:"-" toml:"-"`
	} `json:"extra_config" yaml:"extra_config" toml:"extra_config"`
}

func sliceToStrMap(elements []string) map[string]string {
	elementMap := make(map[string]string)
	for _, s := range elements {
		elementMap[s] = s
	}
	return elementMap
}

func sliceToIntMap(elements []string) map[string]int {
	elementMap := make(map[string]int)
	for _, s := range elements {
		elementMap[s]++
	}
	return elementMap
}

func getDefaultScheme(schemes []string) string {
	// fmt.Sprintf("schemes: %T", schemes)
	for _, v := range schemes {
		if v == "https" { // always push up https protocol
			return fmt.Sprintf("%v", v)
		}
	}
	return schemes[0]
}

func findKeyValue(input string, flatmap map[string]interface{}) string {
	for k, v := range flatmap {
		if input == k {
			return fmt.Sprintf("%v", v)
		}
	}
	return ""
}

func findKey(input string, flatmap map[string]interface{}) string {
	for k, _ := range flatmap {
		if input == k {
			return fmt.Sprintf("%k", k)
		}
	}
	return ""
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			if elements[v] != "" {
				result = append(result, elements[v])
			}
		}
	}
	sort.Strings(result)
	return result // Return the new slice.
}

func ConvertToKrakend(swaggerFile string, prefixPath string, format string) {
	_, err := os.Stat(swaggerFile)
	if err != nil {
		fmt.Printf("ConvertToProtobuf(...): File doesn't exist, skipping\n", swaggerFile)
		return
	}

	s := "10086"
	i, _ := convert.Int(s)
	debug := false
	if debug {
		fmt.Println("int: ", i)
	}

	indexData, err := ioutil.ReadFile(swaggerFile)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// fmt.Printf("%#v\n", tokenizer.Tokenize(indexData))

	dotSlugifier := slugify.New(slugify.Configuration{ReplaceCharacter: '-'})
	var jsonRaw json.RawMessage

	/*
		// from a json []byte
		typed, err := typed.Json(data)

		// from a file containing JSON
		typed, err := typed.JsonFile(path)
	*/
	switch strings.ToLower(format) {
	case "json":
		if !types.IsJSON(string(indexData)) {
			fmt.Println("error not a JSON file content...")
			return
		}
		jsonRaw = indexData
	case "yaml":
		if !types.IsYAML(string(indexData)) {
			fmt.Println("error not a YAML file content...")
			return
		}
		var err error
		var yamlParsed interface{}
		yamlParsed, err = yaml2json.BytesToYAMLDoc([]byte(indexData))
		if err != nil {
			fmt.Println("error while getting yamlParsed: ", err)
			errors.PrintError(err)
			return
		}
		jsonRaw, err = yaml2json.YAMLToJSON(yamlParsed)
		if err != nil {
			fmt.Println(yamlParsed)
			fmt.Println("error while getting jsonRaw: ", err)
			errors.PrintError(err)
			return
		}
	default:
		fmt.Println("format not supported")
		return
	}

	if jsonRaw != nil {

		swaggerExports := &SwaggerExports{}

		var mp map[string]interface{}
		if err := json.Unmarshal([]byte(jsonRaw), &mp); err != nil {
			log.Fatal(err)
		}
		fm, err := flatmap.FlattenWithConfig(mp, flatmap.Config{AddLengthForArrays: true})
		if err != nil {
			log.Fatal(err)
		}
		var ks []string
		for k := range fm {
			ks = append(ks, k)
		}
		sort.Strings(ks)

		typed, err := typed.Json(jsonRaw)
		if err != nil {
			log.Fatal(err)
		}

		if err := configor.Dump(fm, "flatten", "json,yaml", fmt.Sprintf("%s/flatten", prefixPath)); err != nil {
			log.Fatal("ERROR while dumping the config struct:", err.Error())
		}

		// tokens := tokenizer.TokenizeFile(fmt.Sprintf("%s/flatten/flatten.yaml", prefixPath), "yaml", 3)

		queue := make(map[string]bool, 0)
		for _, k := range ks {
			if strings.HasPrefix(k, "paths.") && (strings.Contains(k, "post.") || strings.Contains(k, "get.")) {
				parts := strings.Split(k, ".")
				path := parts[1]
				if queue[path] {
					continue
				}
				backend := &SwaggerExport{}

				var basePath string
				basePath = typed.String("basePath")
				if basePath == "/" {
					basePath = ""
				}

				backend.UrlPattern = fmt.Sprintf("%s%s", basePath, path) // path
				backend.Method = strings.ToUpper(parts[2])
				backend.Extra.LocalFile = swaggerFile
				backend.ConcurrentCalls = 1

				/*
					if len(tokens) > 0 {
						backend.Extra.Tokens = tokens
					}
				*/

				var paramsKeyLength, respAttrLength, respAttrRef string
				var paramsKeyLengthInt, respAttrLengthInt int

				host := typed.String("host")

				consumes := typed.Strings("consumes")
				produces := typed.Strings("produces")
				schemes := typed.StringsOr("schemes", []string{"https"})
				backend.Extra.Schemes = schemes

				providerName := findKeyValue("info.x-providerName", fm)
				backend.Extra.Provider = providerName

				apiDescription := findKeyValue("info.description", fm)
				backend.Extra.Description = apiDescription

				apiVersion := findKeyValue("info.version", fm) // strconv.ParseFloat(findKeyValue("info.version", fm), 64)
				backend.Extra.Version = apiVersion

				/*
					According to W3C:
					- The most correct is application/rss+xml
					- The most compatible is application/xml
				*/
				for k, consume := range consumes {
					mimeParts := strings.Split(consume, "/")
					if len(mimeParts) == 1 {
						consumes[k] = mimeParts[0]
						if mimeParts[0] == "json" {
							backend.Extra.Consume = consumes[k]
							backend.Encoding = consumes[k]
						}
					} else {
						consumes[k] = mimeParts[1]
						if mimeParts[1] == "json" {
							backend.Extra.Consume = consumes[k]
							backend.Encoding = consumes[k]
						}
					}

				}

				backend.Extra.Consumes = consumes
				for k, produce := range produces {
					mimeParts := strings.Split(produce, "/")
					if len(mimeParts) == 1 {
						produces[k] = mimeParts[0]
						if mimeParts[0] == "json" {
							backend.Extra.Produce = produces[k]
						}
					} else {
						produces[k] = mimeParts[1]
						if mimeParts[1] == "json" {
							backend.Extra.Produce = produces[k]
						}
					}

				}

				backend.Extra.Produces = produces
				backend.Extra.Encodings = backend.Extra.Produces
				var hosts []string
				for _, scheme := range schemes {
					if scheme == "https" {
						backend.Scheme = scheme
					}
					hosts = append(hosts, fmt.Sprintf("%s://%s", scheme, host))
				}
				backend.Extra.Hosts = hosts

				if backend.Scheme == "" {
					backend.Scheme = schemes[0]
				}

				if backend.Extra.Consume == "" && len(consumes) > 0 {
					backend.Extra.Consume = consumes[0]
				}

				if backend.Extra.Produce == "" && len(produces) > 0 {
					backend.Extra.Produce = produces[0]
				}

				if backend.Encoding == "" && len(produces) > 0 {
					backend.Encoding = produces[0]
				}

				// basePath
				backend.Host = []string{fmt.Sprintf("%s://%s", backend.Scheme, host)}

				paramsKeyLength = findKeyValue(fmt.Sprintf("paths.%s.%s.parameters.length", parts[1], parts[2]), fm)
				paramsKeyLengthInt, _ = strconv.Atoi(paramsKeyLength)
				backend.AttrsCount = paramsKeyLengthInt

				respAttrLength = findKeyValue(fmt.Sprintf("paths.%s.%s.responses.200.length", parts[1], parts[2]), fm)
				respAttrLengthInt, _ = strconv.Atoi(respAttrLength)
				backend.RespAttrCount = respAttrLengthInt

				var respAttrRefPath string
				respAttrRef = findKeyValue(fmt.Sprintf("paths.%s.%s.responses.200.schema.$ref", parts[1], parts[2]), fm)
				if respAttrRef != "" {
					respAttrRefPath = strings.Replace(respAttrRef, "#/", "", -1)
					respAttrRefPath = strings.Replace(respAttrRefPath, "/", ".", -1)
					backend.RespAttrRefPath = strings.ToLower(respAttrRefPath)
					backend.RespAttrRef = respAttrRef
				}

				var defObjectName string
				_, ok := typed.MapIf("definitions")
				if ok {
					if respAttrRefPath != "" {
						defParts := strings.Split(respAttrRefPath, ".")
						if len(defParts) > 1 {
							defObjectName = defParts[1]
							if debug {
								fmt.Println(" \n defParts[1]: ", defParts[1])
							}
						}
					}
				}
				var fieldList []string
				var fieldListNested []string
				if defObjectName != "" {
					prefixPath := fmt.Sprintf("definitions.%s.properties.", defObjectName)
					for k, _ := range fm {
						if strings.HasPrefix(k, prefixPath) && strings.HasSuffix(k, ".type") {
							propParts := strings.Split(k, ".")
							if debug {
								fmt.Println(" --------- MATCH -------------")
								fmt.Println("key: ", k, " / start with: ", prefixPath, " / end with: .type / len(propParts): ", len(propParts))
								pp.Println(propParts)
							}
							// check the lenght in order to append more dot style paths
							var subPaths []string
							for kp, vp := range propParts {
								if kp >= 3 {
									if kp == 4 && vp == "items" {
										continue
									}
									if vp == "type" || vp == "properties" {
										continue
									}
									if vp != "" {
										subPaths = append(subPaths, vp)
										fieldList = append(fieldList, vp)
									}
								}
							}
							fmt.Printf("######### subPaths: %s \n", strings.Join(subPaths, "."))
							fieldListNested = append(fieldListNested, strings.Join(subPaths, "."))
							if debug {
								fmt.Println("fieldList: ", fieldList)
							}
						}
					}
					if len(fieldList) > 0 {
						fmt.Println("######### fieldList:")
						pp.Print(fieldList)
					}
				}
				// Properties
				/*
					definitions.search-code.properties.items.items.properties.git_url.type: string
					definitions.search-code.properties.items.items.properties.html_url.type: string
					definitions.search-code.properties.items.items.properties.name.type: string
					definitions.search-code.properties.items.items.properties.path.type: string
					definitions.search-code.properties.items.items.properties.repository.properties.archive_url.type: string
					definitions.search-code.properties.items.items.properties.repository.properties.assignees_url.type: string
					definitions.search-code.properties.items.items.properties.repository.properties.blobs_url.type: string
					definitions.search-code.properties.items.items.properties.repository.properties.branches_url.type: string
					definitions.search-code.properties.items.items.properties.repository.properties.collaborators_url.type: string
					definitions.search-code.properties.items.items.properties.repository.properties.comments_url.type: string
					definitions.search-code.properties.items.items.properties.repository.properties.commits_url.type: string
					definitions.search-code.properties.items.items.properties.repository.properties.compare_url.type: string
				*/
				if len(fieldList) > 0 {
					backend.Whitelist = removeDuplicates(fieldListNested)
					backend.Extra.Mapping = sliceToStrMap(removeDuplicates(fieldListNested))
					//backend.Extra.Mapping = sliceToStrMap(removeDuplicates(fieldList))
					if debug {
						pp.Print(backend.Whitelist)
					}
				}
				swaggerExports.Backends = append(swaggerExports.Backends, *backend)
				queue[path] = true
			}
		}

		for k, bck := range swaggerExports.Backends {
			backPath := bck.UrlPattern
			// SLug
			title := findKeyValue("info.title", fm)
			version := findKeyValue("info.version", fm)
			swaggerExports.Backends[k].Extra.Slug = dotSlugifier.Slugify(fmt.Sprintf("%s %s %s", title, version, swaggerExports.Backends[k].UrlPattern))

			// Group
			swaggerExports.Backends[k].Group = dotSlugifier.Slugify(fmt.Sprintf("%s %s", title, version))

			// Parameters
			var paramsList, headerList, urlKeysList, queryStrList, formDataList, bodyList /*, holderList*/ []string
			for i := 0; i <= bck.AttrsCount; i++ { // The location of the parameter. Possible values are "query", "header", "path", "formData" or "body".
				attrIn := findKeyValue(fmt.Sprintf("paths.%s.%s.parameters[%d].in", backPath, strings.ToLower(bck.Method), i), fm)
				attrName := findKeyValue(fmt.Sprintf("paths.%s.%s.parameters[%d].name", backPath, strings.ToLower(bck.Method), i), fm)
				// attrRequired := findKeyValue(fmt.Sprintf("paths.%s.%s.parameters[%d].required", backPath, bck.Protocol, i), fm)
				if attrName != "" {
					if attrIn == "header" {
						headerList = append(headerList, attrName)
					} else if attrIn == "query" {
						queryStrList = append(queryStrList, attrName)
					} else if attrIn == "formData" { // "string", "number", "integer", "boolean", "array" or "file"
						formDataList = append(formDataList, attrName)
					} else if attrIn == "body" {
						bodyList = append(bodyList, attrName)
					} else if attrIn == "path" {
						urlKeysList = append(urlKeysList, attrName)
					} else {
						paramsList = append(paramsList, attrName)
					}
				}
			}
			var categoriesKeyLength, tagListLength string
			var categoriesKeyLengthInt, tagListLengthInt int
			var tagList []string

			categoriesKeyLength = findKeyValue("info.x-apisguru-categories.length", fm)
			categoriesKeyLengthInt, _ = strconv.Atoi(categoriesKeyLength)

			for i := 0; i <= categoriesKeyLengthInt; i++ {
				tagName := findKeyValue(fmt.Sprintf("info.x-apisguru-categories[%d]", i), fm)
				if tagName != "" {
					tagList = append(tagList, dotSlugifier.Slugify(tagName))
				}
			}

			tagListLength = findKeyValue(fmt.Sprintf("paths.%s.%s.tags.length", backPath, strings.ToLower(bck.Method)), fm)
			tagListLengthInt, _ = strconv.Atoi(tagListLength)

			for i := 0; i <= tagListLengthInt; i++ {
				tagLabel := findKeyValue(fmt.Sprintf("paths.%s.%s.tags[%d]", backPath, strings.ToLower(bck.Method), i), fm)
				if tagLabel != "" {
					tagList = append(tagList, dotSlugifier.Slugify(tagLabel))
				}
			}

			if len(tagList) > 0 {
				swaggerExports.Backends[k].Extra.Topics = removeDuplicates(tagList)
			}

			if len(paramsList) > 0 {
				swaggerExports.Backends[k].Params = removeDuplicates(paramsList)
			}
			if len(urlKeysList) > 0 {
				swaggerExports.Backends[k].URLKeys = removeDuplicates(urlKeysList)
			}
			if len(formDataList) > 0 {
				swaggerExports.Backends[k].Extra.Bodies = removeDuplicates(bodyList)
			}
			if len(formDataList) > 0 {
				swaggerExports.Backends[k].Extra.FormDatas = removeDuplicates(formDataList)
			}
			if len(headerList) > 0 {
				swaggerExports.Backends[k].Headers = removeDuplicates(headerList)
			}
			if len(queryStrList) > 0 {
				swaggerExports.Backends[k].QueryStringParams = removeDuplicates(queryStrList)
			}

			// utils.Sort(swaggerExports.Backends[k].Extra, utils.StringComparator)

		}
		if len(swaggerExports.Backends) > 0 && debug {
			fmt.Println("SwaggerExports: ")
			pp.Print(swaggerExports)
		}
		if debug {
			fmt.Println("Paths for swagger file: ", swaggerFile)
		}

		if err := configor.Dump(swaggerExports, "backends", "json,yaml", fmt.Sprintf("%s/krakend", prefixPath)); err != nil {
			log.Fatal("ERROR while dumping the config struct:", err.Error())
		}

	}
}

/*
	krakenBackend = append(krakenBackend, gateway.Export{
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
	})
*/
