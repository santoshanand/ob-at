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
	LotSizeURL           string `mapstructure:"LOT_SIZE_CSV"`
	InstrumentZerodhaURL string `mapstructure:"INSTRUMENT_ZERODHA_CSV"`
	Username             string `mapstructure:"Z_USERNAME"`
	Password             string `mapstructure:"Z_PASSWORD"`
	ServerAddress        string `mapstructure:"SERVER_ADDRESS"`
	Environment          string `mapstructure:"ENVIRONMENT"`
	RunCron              bool   `mapstructure:"RUN_CRON"`
	PostgresDBURL        string `mapstructure:"POSTGRES_DB_URL"`
	NseHostURL           string `mapstructure:"NSE_HOST_URL"`
	ZerodhaHostURL       string `mapstructure:"ZERODHA_HOST_URL"`
	TestCron             bool   `mapstructure:"TEST_CRON"`
	LoadStocks           bool   `mapstructure:"LOAD_STOCKS"`
	BUY                  bool   `mapstructure:"BUY"`
	IndicesTradeInterval string `mapstructure:"INDICES_TRADE_INTERVAL"`
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
