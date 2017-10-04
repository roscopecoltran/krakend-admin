package main

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

type LocalEnvironment struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Debug    bool   `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Instance string `gorm:"column:instance" mapstructure:"instance" storm:"instance" json:"instance" yaml:"instance" toml:"instance"`
	File     oss.OSS
	Envs     []LocalEnvironmentVariable `gorm:"many2many:gateway_local_envvars;" json:"variables,omitempty" yaml:"variables,omitempty" toml:"variables,omitempty"`
}

type LocalEnvironmentVariable struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`
	Disabled   bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	File       string `gorm:"column:file" storm:"file" json:"file" yaml:"file" toml:"file"`
	Key        string `gorm:"column:key" storm:"key" json:"key" yaml:"key" toml:"key"`
	Val        string `gorm:"column:val" storm:"val" json:"val" yaml:"val" toml:"val"`
}

type LocalFile struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled   bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug      bool   `gorm:"column:debug" default:"false" example:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Name       string `gorm:"column:filename" storm:"filename" json:"filename" yaml:"filename" toml:"filename"`
	Path       string `gorm:"column:filepath" storm:"filepath" json:"filepath" yaml:"filepath" toml:"filepath"`
	PrefixPath string `gorm:"column:prefix_path" storm:"prefix_path" json:"prefix_path" yaml:"prefix_path" toml:"prefix_path"`
	File       oss.OSS
}

type TaskQueue struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled    bool   `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug       bool   `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`
	Name        string `gorm:"column:name" storm:"name" json:"name" yaml:"name" toml:"name"`
	Description string `gorm:"column:description" storm:"description" json:"description" yaml:"description" toml:"description"`
	Code        string `gorm:"column:code" storm:"code" json:"code" yaml:"code" toml:"code" `
	Runner      string `gorm:"column:runner" storm:"runner" json:"runner" yaml:"runner" toml:"runner"`
	Workdir     string `gorm:"column:workdir" storm:"workdir" json:"workdir" yaml:"workdir" toml:"workdir"`
	Register    string `gorm:"column:register" storm:"register" json:"register" yaml:"register" toml:"register"`
	Type        string `gorm:"column:type" storm:"type" json:"type" yaml:"type" toml:"type"` // series, parallel
	Action      string `gorm:"column:action" storm:"action" json:"action" yaml:"action" toml:"action"`
	Order       int    `gorm:"column:order" storm:"order" json:"order" yaml:"order" toml:"order"`
	Priority    int    `gorm:"column:disabpriorityled" storm:"priority" json:"priority" yaml:"priority" toml:"priority"`

	//PreTasks  []TaskQConfig `gorm:"many2many:gateway_tasks_pre;" storm:"pre_tasks" json:"pre_tasks" yaml:"pre_tasks" toml:"pre_tasks"`
	//PostTasks []TaskQConfig `gorm:"many2many:gateway_tasks_post;" storm:"post_tasks" json:"post_tasks" yaml:"post_tasks" toml:"post_tasks"`
	//Envs      []EnvConfig   `gorm:"many2many:gateway_tasks_env;" storm:"post_tasks" json:"post_tasks" yaml:"post_tasks" toml:"post_tasks"`
}

var TaskQueueRunners = []string{"sh", "bash", "python", "ruby", "perl", "nodejs", "go"}

type TaskQueueRunner struct {
	gorm.Model `json:"-" yaml:"-" toml:"-"`

	Disabled bool `gorm:"column:disabled" default:"false" storm:"disabled" json:"disabled" yaml:"disabled" toml:"disabled"`
	Debug    bool `gorm:"column:debug" default:"false" storm:"debug" json:"debug" yaml:"debug" toml:"debug"`

	Name        string `gorm:"column:name" storm:"name" json:"name" yaml:"name" toml:"name"`
	Type        string `gorm:"column:type" storm:"type" json:"type" yaml:"type" toml:"type"`
	Description string `gorm:"column:description" storm:"description" json:"description" yaml:"description" toml:"description"`
}
