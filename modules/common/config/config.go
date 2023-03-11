package config

import (
	"os"
	"path"
	"runtime"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// Config -
type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	PostgresDBURL string `mapstructure:"POSTGRES_DB_URL"`
}

// NewConfig -
func NewConfig() *Config {
	return loadConfig(".env")
}

func loadConfig(fileName string) *Config {
	conf := &Config{}
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	dirPath := path.Join(d, "../../../")
	if os.Getenv("ENVIRONMENT") == "prod" {
		dirPath = "."
	}
	viper.AddConfigPath(dirPath)
	viper.SetConfigName(fileName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(conf)
	if err != nil {
		panic(err)
	}
	return conf
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(NewConfig),
)
