package models

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
	// "github.com/qor/l10n"
	// "github.com/qor/sorting"
)

type ClassifyTopic struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	// l10n.Locale
	// sorting.Sorting

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Name        string `gorm:"column:name" storm:"name" json:"name" yaml:"name" toml:"name"`
	Description string `gorm:"column:desc" storm:"desc" json:"desc" yaml:"desc" toml:"desc"`
	Code        string `gorm:"column:code" storm:"code" json:"code" yaml:"code" toml:"code"`
	Slug        string `gorm:"column:slug" storm:"slug" json:"slug" yaml:"slug" toml:"slug"`

	Topics []ClassifyTopic // `gorm:"many2many:classify_topics;" storm:"topics" json:"topics" yaml:"topics" toml:"topics"`
	// TopicID uint            `gorm:"column:topic_id" storm:"topic_id" json:"topic_id" yaml:"topic_id" toml:"topic_id"`
}

func (topic ClassifyTopic) Validate(db *gorm.DB) {
	if strings.TrimSpace(topic.Name) == "" {
		db.AddError(validations.NewError(topic, "Name", "Name can not be empty"))
	}
}

func (topic ClassifyTopic) DefaultPath() string {
	if len(topic.Code) > 0 {
		return fmt.Sprintf("/topic/%s", topic.Code)
	}
	return "/"
}
