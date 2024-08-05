package config

import (
	"github.com/lapkomo2018/goServiceExample/internal/transport/grpc"
	"github.com/lapkomo2018/goServiceExample/internal/transport/rest"
	"github.com/lapkomo2018/goServiceExample/pkg/hash"
	"github.com/lapkomo2018/goServiceExample/pkg/validation"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	RestServer rest.Config
	GrpcServer grpc.Config
	Hash       hash.Config
	Validator  validation.Config
}

func LoadConfig() (*Config, error) {
	log.Println("Loading config...")

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	log.Println("Loaded config")
	return &config, nil
}
