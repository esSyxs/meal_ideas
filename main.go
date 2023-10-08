package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Db DBConf `yaml:"DBConf"`
}

type DBConf struct {
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	DBName   string `yaml:"DBName"`
}

var Conf Config

var configFlag string
var createDB bool

func init() {
	flag.StringVar(&configFlag, "config", "config.yaml", "path to config")
	flag.Parse()
}

func main() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}

}

func loadConfig() error {
	yfile, err := os.ReadFile(configFlag)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yfile, &Conf)
	if err != nil {
		return err
	}

	// remove later
	fmt.Println(Conf)

	return nil
}
