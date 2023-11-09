package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	LeetCode struct {
		UserSlug []string `yaml:"userSlug"`
	} `yaml:"leetcode"`
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
}

var C Config

func Init() {
	_, err := os.Stat("./config/config.yml")
	if os.IsNotExist(err) {
		os.Mkdir("./config", 0755)
		os.Create("./config/config.yml")
	}

	dataBytes, err := os.ReadFile("./config/config.yml")
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		panic(err)
	}

	C = config
}

func GetConfig() Config {
	return C
}
