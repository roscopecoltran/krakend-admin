package backend

import (
	"fmt"
	// "log"
	"os"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/qor/qor"

	"github.com/athom/suitecase"
	"github.com/qor/action_bar"
	"github.com/qor/admin"
	"github.com/qor/media"
	"github.com/qor/publish2"
	"github.com/qor/validations"

	// "github.com/qor/help"
	// "github.com/qor/media/asset_manager"
	// "github.com/qor/sorting"
	// "github.com/qor/l10n"

	/*
		"github.com/qor/notification"
		"github.com/qor/notification/channels/database"

		"github.com/qor/assetfs"
		"github.com/qor/cache"
		"github.com/qor/filebox"
		"github.com/qor/help"
		"github.com/qor/i18n/exchange_actions"
		"github.com/qor/l10n"
		"github.com/qor/location"
		"github.com/qor/media"
		"github.com/qor/media/media_library"
		"github.com/qor/middlewares"
		"github.com/qor/page_builder"
		"github.com/qor/publish2"
		"github.com/qor/qor/resource"
		"github.com/qor/qor/utils"
		"github.com/qor/roles"
		"github.com/qor/serializable_meta"
		"github.com/qor/slug"
		"github.com/qor/transition"
		"github.com/qor/validations"
		"github.com/qor/variations"
		"github.com/qor/widget"
		"github.com/qor/worker"
	*/

	"github.com/k0kubun/pp"
)

type AdminConfig struct {
	UI        *admin.Admin
	ActionBar *action_bar.ActionBar
	Disabled  bool
	IsReady   bool
}

func New(tables ...interface{}) error {
	Store.Gorm.DB, _ = gorm.Open("sqlite3", "./shared/stores/krakend.db")

	if os.Getenv("DEBUG") != "" {
		Store.Gorm.DB.LogMode(true)
	}

	// Store.Gorm.DB = Store.Gorm.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff)

	media.RegisterCallbacks(Store.Gorm.DB)
	validations.RegisterCallbacks(Store.Gorm.DB)
	publish2.RegisterCallbacks(Store.Gorm.DB)
	// l10n.RegisterCallbacks(Store.Gorm.DB)
	// sorting.RegisterCallbacks(Store.Gorm.DB)

	if !Store.Admin.Disabled {

		// Admin = admin.New(&qor.Config{DB: db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff)})
		Store.Admin.UI = admin.New(&qor.Config{DB: Store.Gorm.DB})

		Store.Admin.UI.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin"})

		// Add Notification
		// Notification := notification.New(&notification.Config{})
		// Notification.RegisterChannel(database.New(&database.Config{DB: Store.Gorm.DB}))
		// Store.Admin.UI.NewResource(Notification)

		// Add Asset Manager, for rich editor
		// assetManager := Store.Admin.UI.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})

		// Add Help
		// Help := Store.Admin.UI.NewResource(&help.QorHelpEntry{})
		// Help.GetMeta("Body").Config = &admin.RichEditorConfig{AssetManager: assetManager}

		// Add Worker
		/*
			Worker := getWorker()
			exchange_actions.RegisterExchangeJobs(i18n.I18n, Worker)
			Admin.AddResource(Worker, &admin.Config{Menu: []string{"Site Management"}})
		*/

	}

	for _, table := range tables {

		// pp.Print(table)
		fmt.Println("Store.Gorm.IsTruncate", Store.Gorm.IsTruncate)
		if Store.Gorm.IsTruncate {
			if err := Store.Gorm.DB.DropTableIfExists(table).Error; err != nil {
				pp.Print(err)
				return err
			}
		}

		fmt.Println("Store.Gorm.IsAutoMigrate", Store.Gorm.IsAutoMigrate)
		if Store.Gorm.IsAutoMigrate {
			if err := Store.Gorm.DB.AutoMigrate(table).Error; err != nil {
				pp.Print(err)
				return err
			}
		}

		fmt.Println("Store.Admin.Disabled", Store.Admin.Disabled)
		if !Store.Admin.Disabled {
			varTypeName := getVarTypeName(table)
			varTypeNameStr := strings.Replace(varTypeName, "*", "", -1)
			sectionName := strings.Title(ExtractMenuSectionName(varTypeNameStr))
			fmt.Println(" -- varTypeNameStr=", varTypeNameStr, " > varTypeName=", varTypeName, " > sectionName=", sectionName)

			tableShortNameStr := strings.Replace(varTypeNameStr, sectionName, "", -1)
			Store.Admin.UI.AddResource(table, &admin.Config{Name: tableShortNameStr, Menu: []string{sectionName}})
			// section[sectionName] := Store.Admin.UI.AddResource(table, &admin.Config{Menu: []string{sectionName}})
			// section[sectionName].Meta(&admin.Meta{Name: "Gender", Config: &admin.SelectOneConfig{Collection: Genders, AllowBlank: true}})
			// section[sectionName].Meta(&admin.Meta{Name: "Category", Config: &admin.SelectOneConfig{AllowBlank: true}})
			// product.SearchAttrs("Name", "Code", "Category.Name", "Brand.Name")
			// product.IndexAttrs("MainImageURL", "Name", "Price", "VersionName", "PublishLiveNow")
		}

	}

	if !Store.Admin.Disabled {

		// category := Admin.AddResource(&models.Category{}, &admin.Config{Menu: []string{"Product Management"}, Priority: -3})
		// Store.Admin.UI.GetResource("ClassifyCategory")
		// category.Meta(&admin.Meta{Name: "Categories", Type: "select_many"})

		// Add Search Center Resources
		// Admin.AddSearchResource(product, user, order)

		// Add ActionBar
		// ActionBar = action_bar.New(Admin)
		// ActionBar.RegisterAction(&action_bar.Action{Name: "Admin Dashboard", Link: "/admin"})

		if len(Store.Admin.UI.GetResources()) > 0 {
			for _, resource := range Store.Admin.UI.GetResources() {
				fmt.Println("var.adminui.resource", resource)
			}
		}
	}

	return nil
}

func getVarTypeName(variable interface{}) string {
	if t := reflect.TypeOf(variable); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func ExtractMenuSectionName(camelInputStr string) string {
	if camelInputStr == "" {
		return "Global"
	}
	convStr := suitecase.ToSnakeCase(camelInputStr)
	parts := strings.Split(convStr, "_")
	// lastPart := info[len(parts)-1]
	if len(parts) > 0 {
		return parts[0]
	} else {
		return "Global"
	}
}

/*
func AddMenu(name string, link string) error {
	if name != "" && link != "" {
		Store.Admin.UI.AddMenu(&admin.Menu{Name: fmt.Sprintf("%s", strings.ToTitle(name)), Link: fmt.Sprintf("%s", link)})
		return nil
	} else {
		return errors.New("Missing name or link attribute.")
	}
}
*/

/*
func SetSearchAResources(resources ...interface{}) {
	Store.Admin.UI.AddSearchResource(resources)
}

func SetSearchAttributes(model interface{}, attributes ...string) {
	if len(attributes) > 0 {
		model.SearchAttrs(attributes) // "Name", "Code", "Category.Name", "Brand.Name")
	}
}
*/
