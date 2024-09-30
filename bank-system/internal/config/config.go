package config

import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	Server struct {
		Port string
	}
}

// LoadConfig загружает конфигурацию из файла config.yaml
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")  // имя файла (без расширения)
	viper.SetConfigType("yaml")    // расширение файла
	viper.AddConfigPath("./config") // путь к папке config

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil
}
