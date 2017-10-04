package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
	"github.com/qor/sorting"
	"github.com/qor/validations"
)

type ClassifyCategory struct {
	gorm.Model `storm:"-" json:"-" yaml:"-" toml:"-"`
	l10n.Locale
	sorting.Sorting

	Disabled bool `default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Name        string `gorm:"not null;type:varchar(100);unique" storm:"name" json:"name" yaml:"name" toml:"name"`
	Description string `storm:"desc" json:"desc" yaml:"desc" toml:"desc"`
	Code        string `storm:"code" json:"code" yaml:"code" toml:"code"`
	Slug        string `storm:"slug" json:"slug" yaml:"slug" toml:"slug"`

	Categories []ClassifyCategory // `gorm:"many2many:classify_categories;" storm:"categories" json:"categories" yaml:"categories" toml:"categories"`
	// CategoryID uint               `gorm:"column:category_id" storm:"category_id" json:"category_id" yaml:"category_id" toml:"category_id"`
}

func (category ClassifyCategory) Validate(db *gorm.DB) {
	if strings.TrimSpace(category.Name) == "" {
		db.AddError(validations.NewError(category, "Name", "Name can not be empty"))
	}
}

func (category ClassifyCategory) DefaultPath() string {
	if len(category.Code) > 0 {
		return fmt.Sprintf("/category/%s", category.Code)
	}
	return "/"
}
