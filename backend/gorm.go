package backend

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	//"github.com/qor/l10n"
	"github.com/qor/media"
	"github.com/qor/publish2"
	//"github.com/qor/sorting"
	"github.com/qor/validations"
)

type GormStorageConfig struct {
	DB            *gorm.DB
	Config        DBConfig
	Disabled      bool
	IsReady       bool
	IsTruncate    bool
	IsAutoMigrate bool
}

type DBConfig struct {
	Name     string `env:"DBName" default:"gateway"`
	Adapter  string `env:"DBAdapter" default:"sqlite"`
	Host     string `env:"DBHost" default:"localhost"`
	Port     string `env:"DBPort" default:"3306"`
	User     string `env:"DBUser" default:"root"`
	Password string `env:"DBPassword" default:"password"`
}

func NewRDB() {
	var err error
	var DB *gorm.DB
	dbConfig := Store.Gorm.Config

	if Store.Gorm.Config.Adapter == "mysql" {
		DB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
	} else if Store.Gorm.Config.Adapter == "postgres" {
		DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name))
	} else if Store.Gorm.Config.Adapter == "sqlite" {
		DB, err = gorm.Open("sqlite3", fmt.Sprintf("%v/%v", os.TempDir(), dbConfig.Name))
	} else {
		panic(errors.New("not supported database adapter"))
	}

	if err == nil {
		if os.Getenv("DEBUG") != "" {
			Store.Gorm.DB.LogMode(true)
		}
		// l10n.RegisterCallbacks(DB)
		// sorting.RegisterCallbacks(DB)
		validations.RegisterCallbacks(DB)
		media.RegisterCallbacks(DB)
		publish2.RegisterCallbacks(DB)

		Store.Gorm.DB = DB

	} else {
		panic(err)
	}

}

func TruncateTables(tables ...interface{}) {
	for _, table := range tables {
		if err := Store.Gorm.DB.DropTableIfExists(table).Error; err != nil {
			panic(err)
		}
		Store.Gorm.DB.AutoMigrate(table)
	}
}
