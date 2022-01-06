package config

import (
	"io/ioutil"
	"log"
	"path"

	"github.com/pangud/pangud/pkg/shared/config"
	"gopkg.in/yaml.v2"
)

const appCfgName = "pangud.yml"

func LoadConfigFromFile(cfgDir string) {
	file := path.Join(cfgDir, appCfgName)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Could not read config file")
	}
	var cfg config.Config
	yaml.Unmarshal(data, &cfg)
	config.Set(&cfg)
}
