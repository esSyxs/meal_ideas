package main

import (
	logger "System/Log"
	"System/food"
	"System/server"
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

var Conf Config

var configFlag string

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
	Init     bool   `yaml:"Init"`
}

type Server struct {
	Port string `yaml:"Port"`
}

func init() {
	flag.StringVar(&configFlag, "config", "config.yaml", "path to config")
	flag.Parse()
}

func main() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}

	if Conf.Db.Init {
		err := food.InitDB(
			Conf.Db.User,
			Conf.Db.Password,
			Conf.Db.Host,
			Conf.Db.Port,
			Conf.Db.DBName,
		)

		if err != nil {
			panic(err)
		}

		os.Exit(0)
	}

	server.Start(
		Conf.Server.Port,
		Conf.Db.User,
		Conf.Db.Password,
		Conf.Db.Host,
		Conf.Db.Port,
		Conf.Db.DBName,
	)
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
