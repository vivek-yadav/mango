package config

import "github.com/vivek-yadav/mango/utils"

var CurrentConfig Config

type Config struct {
	Serve  ServeConfig     `mapstructure:"serve"`
	Dgraph DgraphConfig    `mapstructure:"dgraph"`
	Log    utils.LogConfig `mapstructure:"log"`
}

type ServeConfig struct {
	Port int `mapstructure:"port"`
}

type DgraphConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
