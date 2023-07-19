package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	cfg := Config{}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper could not read config file:%v", err)
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
