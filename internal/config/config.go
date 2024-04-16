package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-required:"true"`
	HTTPServer `yaml:"http_server"`
	Storage    `yaml:"storage"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8282"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Storage struct {
	ConnectionString string `yaml:"connection_string" env-required:"true"`
	CarsTable        string `yaml:"cars_table"`
}

var Cfg Config

func MustLoad() {
	configPath := (`.\config\local.yaml`)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist: ", configPath)
	}

	if err := cleanenv.ReadConfig(configPath, &Cfg); err != nil {
		log.Fatal("cannot read config: ", err)
	}
}
