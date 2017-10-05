package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/Machiel/slugify"

	"github.com/roscopecoltran/configor"
	"github.com/roscopecoltran/krakend-admin/db"
	"github.com/roscopecoltran/krakend-admin/models/provider/restful"
	"github.com/roscopecoltran/krakend-admin/utils/download"
)

var (
	Cfg          Config
	strSlugifier *slugify.Slugifier
)

var (
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

func init() {
	strSlugifier = slugify.New(slugify.Configuration{ReplaceCharacter: '-'})
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

	db.Store.Admin.Disabled = false
	db.Store.Gorm.IsTruncate = true
	db.Store.Gorm.IsAutoMigrate = true

	/*
		db, err := bolt.Open(*dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
		if err != nil {
			log.Fatal(err)
		}
		backend.Store.Bolt.DB = db
		defer db.Close()
	*/

	if err := db.New(defaultModels...); err != nil {
		log.Fatal("ERROR while init backend components:", err.Error())
	}

	for _, p := range Cfg.Providers {
		krakenBackend, gormBackend := convert(p)
		Cfg.Backends = append(Cfg.Backends, krakenBackend)
		if err := importNewProvider(p); err != nil {
			fmt.Println("error occured while importing provider: ", p.Name)
		}
		if ok := db.Store.Gorm.DB.NewRecord(gormBackend); ok {
			db.Store.Gorm.DB.Create(&gormBackend)
		}
	}

	// Create a wait group to report to and a work channel that takes URLs
	wg := new(sync.WaitGroup)
	work := make(chan string)

	// Create workers and start goroutines listening to work channel
	numWorkers := 25
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go download.Worker(work, wg)
	}

	if len(Cfg.Specs.Registry) > 0 {
		if err := createNewApiGuruProfile(Cfg.Specs.Registry); err != nil {
			fmt.Println("error occured while importing apis-guru profile.")
		}
	}

	// Add URLs to work channel from CSV
	download.AddWorkFromCsv(work, fmt.Sprintf("%v", *downloadCsv))

	if ok := db.Store.Gorm.DB.NewRecord(Cfg.Krakend); ok {
		db.Store.Gorm.DB.Create(&Cfg.Krakend)
	}

	for _, e := range Cfg.Krakend.Endpoints {
		if ok := db.Store.Gorm.DB.NewRecord(e); ok {
			db.Store.Gorm.DB.Create(&e)
		}
		for _, b := range e.Backend {
			if ok := db.Store.Gorm.DB.NewRecord(b); ok {
				db.Store.Gorm.DB.Create(&b)
			}
		}
	}

	// Wait for workers to finish
	close(work)
	wg.Wait()

	if len(Cfg.Specs.Registry) > 0 {
		restful.ParseSpecs(Cfg.Specs.Registry, Cfg.Patterns.Ignore)
	}

	fmt.Println("Done.")

	startAdminUI()

}
