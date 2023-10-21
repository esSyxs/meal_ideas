package main

import (
	logger "System/Log"
	"System/server"
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Db     DBConf `yaml:"DBConf"`
	Server Server `yaml:"Server"`
}

type DBConf struct {
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	DBName   string `yaml:"DBName"`
}

type Server struct {
	Port string `yaml:"Port"`
}

var Conf Config

var configFlag string

func init() {
	flag.StringVar(&configFlag, "config", "config.yaml", "path to config")
	flag.Parse()
}

func main() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}

	server.Start(Conf.Server.Port)
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
	logger.Default.Println("configuration for debugging", Conf)

	return nil
}
