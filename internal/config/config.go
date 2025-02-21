package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type NotificationServiceConfig struct{
	Env string `yaml:"env"`
	GRPCServer `yaml:"grpc_server"`
}

type GRPCServer struct{
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func MustLoad () *NotificationServiceConfig{
	var cfg NotificationServiceConfig
	
	configPath := os.Getenv("NOTIF_CONFIG_PATH")
	if configPath == ""{
		log.Fatal("config path is empty")
	}

	if _, err := os.Stat(configPath); err != nil{
		log.Fatal("config file was not found")
	}

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil{
		log.Fatal("failed to read config file")
	}

	return &cfg
}