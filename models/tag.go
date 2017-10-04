package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
	// "github.com/qor/l10n"
	// "github.com/qor/sorting"
)

type ClassifyTag struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	// l10n.Locale
	// sorting.Sorting

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Name        string        `gorm:"column:name" storm:"name" json:"name" yaml:"name" toml:"name"`
	Description string        `gorm:"column:desc" storm:"desc" json:"desc" yaml:"desc" toml:"desc"`
	Slug        string        `gorm:"column:slug" storm:"slug" json:"slug" yaml:"slug" toml:"slug"`
	Tags        []ClassifyTag // `gorm:"many2many:classify_tags;" storm:"tags" json:"tags" yaml:"tags" toml:"tags"`
	// TagID       uint          `gorm:"column:tag_id" storm:"tag_id" json:"tag_id" yaml:"tag_id" toml:"tag_id"`
}

func (tag ClassifyTag) Validate(db *gorm.DB) {
	if strings.TrimSpace(tag.Name) == "" {
		db.AddError(validations.NewError(tag, "Name", "Name can not be empty"))
	}
}

func (tag ClassifyTag) DefaultPath() string {
	if len(tag.Slug) > 0 {
		return fmt.Sprintf("/tag/%s", tag.Slug)
	}
	return "/"
}
