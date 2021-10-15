package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type SingleHost struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Key  string `yaml:"privateKey"`
}

type HostList struct {
	List []SingleHost `yaml:"list"`
}

func GetConfig() []SingleHost {
	config := HostList{}
	pwd, _ := os.Getwd()
	cfg, err := ioutil.ReadFile(pwd + `/utils/config/config.yaml`)
	if err != nil {
		log.Fatalf("io error: %v", err)
	}
	err = yaml.UnmarshalStrict(cfg, &config)
	if err != nil {
		log.Fatalf("decode error: %v", err)
	}
	return config.List
}
