package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ServerConfig struct {
	System System `yaml:"system"`
	Mysql Mysql `yaml:"mysql"`
}

var config *ServerConfig
var DataBase *sql.DB

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &config)
}

func GetDBConfiguration() *Mysql {
	return &config.Mysql
}

func GetSystemConfiguration() *System {
	return &config.System
}

func LoadDB(db *sql.DB) {
	DataBase = db
}