package config

import (
	"github.com/spf13/viper"
	"github.com/golang-jwt/jwt/v5"

	"sync"
	"strings"
	"time"
)

type (
	Config struct {
		App *App
		Database *Database
		Auth *Auth
		Hash *Hash
		Weather *Weather
		Redis *Redis
	}

	App struct {
		Address string
		Port string
	}

	Database struct {
		Host string
		Name string	
		User string
		Port string
		Password string
	}

	Auth struct {
		JWTSecret string
		SigningMethod jwt.SigningMethod
		Access Token
		Refresh Token
	}

	Token struct {
		ExpirationTime int
	}

	Hash struct {
		Cost int
	}

	Weather struct {
		Key string
		URL string
	}

	Redis struct {
		Host string
		Port string
		Password string
		DB int // db number to select on connection
		TaskExpiration time.Duration
		WeatherExpiration time.Duration
	}
)

var (
	once sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func () {
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")	

		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		viper.SetDefault("app.address", "localhost")
		viper.SetDefault("app.port", "8000")

		viper.SetDefault("database.name", "postrges")
		viper.SetDefault("database.user", "postgres")
		viper.SetDefault("database.port", "5432")
		viper.SetDefault("database.host", "db")

		viper.SetDefault("auth.access.expirationtime", 60*60)
		viper.SetDefault("auth.refresh.expirationtime", 60*60*24*7)
		viper.SetDefault("auth.signingmethod", jwt.SigningMethodHS256)

		viper.SetDefault("hash.cost", 10)

		viper.SetDefault("weather.key", "")

		viper.SetDefault("redis.host", "redis")
		viper.SetDefault("redis.port", "6379")
		viper.SetDefault("redis.db", 0)
		viper.SetDefault("redis.taskexpiration", 30*time.Minute)
		viper.SetDefault("redis.weatherexpiration", time.Hour)

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
