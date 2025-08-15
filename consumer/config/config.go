package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Postgres Postgres
	Kafka    Kafka
	Redis    Redis
}

type Server struct {
	Host string
	Port int
}

type Postgres struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

type Kafka struct {
	Brokers string
	Topic   string
}

type Redis struct {
	Address string
	DB      int
}

func GetConfig() Config {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read .env file: " + err.Error())
	}

	return Config{
		Server: Server{
			Host: viper.GetString("SRV_HOST"),
			Port: viper.GetInt("SRV_PORT"),
		},
		Postgres: Postgres{
			Username: viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			Host:     viper.GetString("POSTGRES_HOST"),
			Port:     viper.GetString("POSTGRES_PORT"),
			DBName:   viper.GetString("POSTGRES_DB"),
		}, Kafka: Kafka{
			Brokers: viper.GetString("KAFKA_BROKERS"),
			Topic:   viper.GetString("KAFKA_TOPIC"),
		}, Redis: Redis{
			Address: viper.GetString("REDIS_ADDRESS"),
			DB:      viper.GetInt("REDIS_DB"),
		},
	}
}
