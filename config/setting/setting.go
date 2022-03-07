package setting

import (
	"gimSec/basic/logging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Setting struct {
	RunMode  string   `yaml:"runMode"`
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
}

type server struct {
	HTTPPort int `yaml:"HTTPPort"`
}

type database struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	password string `yaml:"password"`
	host     string `yaml:"host"`
	dbName   string `yaml:"dbName"`
	port     string `yaml:"port"`
}

var Conf = &Setting{}

func init() {
	yamlFile, err := ioutil.ReadFile("config/" + "develop-gbits" + ".yaml")
	if err != nil {
		logging.Error(err)
	}

	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		panic(err)
	}

	logging.Info(Conf)
}
