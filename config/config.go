package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SshConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Cmd      string
	File     string
}

func ParseConfig(filePath string) (config *SshConfig) {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	config = &SshConfig{}
	err = yaml.Unmarshal(dat, config)
	if err != nil {
		panic(err)
	}
	return
}
