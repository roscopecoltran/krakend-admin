package local

import (
	"github.com/jinzhu/gorm"
)

var TaskQueueRunners = []string{"sh", "bash", "python", "ruby", "perl", "nodejs", "go"}

type TaskQueue struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Name        string        `gorm:"column:name" storm:"name" json:"name" yaml:"name" toml:"name"`
	Description string        `gorm:"column:description" storm:"description" json:"description" yaml:"description" toml:"description"`
	Code        string        `gorm:"column:code" storm:"code" json:"code" yaml:"code" toml:"code" `
	Runner      string        `gorm:"column:runner" storm:"runner" json:"runner" yaml:"runner" toml:"runner"`
	Workdir     string        `gorm:"column:workdir" storm:"workdir" json:"workdir" yaml:"workdir" toml:"workdir"`
	Register    string        `gorm:"column:register" storm:"register" json:"register" yaml:"register" toml:"register"`
	Type        string        `gorm:"column:type" storm:"type" json:"type" yaml:"type" toml:"type"` // series, parallel
	Action      string        `gorm:"column:action" storm:"action" json:"action" yaml:"action" toml:"action"`
	Order       int           `gorm:"column:order" storm:"order" json:"order" yaml:"order" toml:"order"`
	Priority    int           `gorm:"column:disabpriorityled" storm:"priority" json:"priority" yaml:"priority" toml:"priority"`
	PreTasks    []TaskQueue   `storm:"pre_tasks" json:"pre_tasks" yaml:"pre_tasks" toml:"pre_tasks"`
	PostTasks   []TaskQueue   `storm:"post_tasks" json:"post_tasks" yaml:"post_tasks" toml:"post_tasks"`
	Envs        []Environment `gorm:"many2many:gateway_tasks_env;" storm:"post_tasks" json:"post_tasks" yaml:"post_tasks" toml:"post_tasks"`
}

type TaskQueueRunner struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Name        string `gorm:"column:name" storm:"name" json:"name" yaml:"name" toml:"name"`
	Type        string `gorm:"column:type" storm:"type" json:"type" yaml:"type" toml:"type"`
	Description string `gorm:"column:description" storm:"description" json:"description" yaml:"description" toml:"description"`
}
