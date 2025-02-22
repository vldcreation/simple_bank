package app

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
	"github.com/vldcreation/simple_bank/consts"
	"github.com/vldcreation/simple_bank/util"
)

type Config struct {
	APP   *APPConfig   `yaml:"app" env:"app" mapstructure:"app"`
	DB    *DBConfig    `yaml:"db" env:"db" mapstructure:"db"`
	Token *TokenConfig `yaml:"token" env:"token" mapstructure:"token"`
}

type DBConfig struct {
	Driver   string `yaml:"driver" json:"driver" env:"DB_DRIVER" mapstructure:"db_driver"`
	User     string `yaml:"user" json:"user" env:"DB_USER" mapstructure:"db_user"`
	Password string `yaml:"password" json:"password" env:"DB_PASSWORD" mapstructure:"db_password"`
	Host     string `yaml:"host" json:"host" env:"DB_HOST" mapstructure:"db_host"`
	Port     string `yaml:"port" json:"port" env:"DB_PORT" mapstructure:"db_port"`
	Database string `yaml:"database" json:"database" env:"DB_DATABASE" mapstructure:"db_database"`
}

type APPConfig struct {
	Name string `yaml:"name" json:"name" env:"APP_NAME" mapstructure:"app_name"`
	Port string `yaml:"port" json:"port" env:"APP_PORT" mapstructure:"app_port"`
	Env  string `yaml:"env" json:"env" env:"APP_ENV" mapstructure:"app_env"`
}

type TokenConfig struct {
	Generator           string        `yaml:"generator" json:"generator" env:"token_generator" mapstructure:"token_generator"`
	SecretKey           string        `yaml:"secret_key" json:"secret_key" env:"token_secret_key" mapstructure:"token_secret_key"`
	AccessTokenDuration time.Duration `yaml:"access_token_duration" json:"token_access_token_duration" env:"token_access_token_duration" mapstructure:"token_access_token_duration"`
}

var (
	once sync.Once
	_cfg *Config
)

func NewConfigFromYaml(path string) *Config {
	if path == "" {
		path = consts.ConfigPath
	}

	fpath := []string{path}
	once.Do(func() {
		c, err := readCfg("env.yaml", fpath...)
		if err != nil {
			log.Fatal(err)
		}

		_cfg = c
	})

	return _cfg
}

// NOTE: this NewConfig parser only work with single struct
// TODO: handle nested struct
func NewConfigFromEnv(path string) *Config {
	var (
		configDB = DBConfig{}
		_cfg     = &Config{}
	)

	if path == "" {
		path = consts.ConfigPath
	}

	log.Printf("path == %s", path)

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.ConfigFileUsed()

	err := viper.ReadInConfig()
	if err != nil {
		// load config from environtment variables
		if err = util.SetEnvValue("env", &configDB); err != nil {
			log.Fatalf("unable to parse config from environtment vars: %+v", err)
		}

		_cfg.DB = &configDB
	} else {
		err = viper.Unmarshal(&_cfg.DB)
		if err != nil {
			log.Fatalf("unable to parse config %+v", err)
		}
	}

	log.Printf("DB %+v", _cfg.DB)

	if _cfg.DB == nil {
		log.Fatalf("env parse error")
	}

	return _cfg
}

// *config.Configuration: configuration ptr object
// error: error operation
func readCfg(fname string, ps ...string) (*Config, error) {
	var cfg *Config
	var errs []error

	for _, p := range ps {
		f := fmt.Sprint(p, fname)

		err := util.ReadFromYAML(f, &cfg)
		if err != nil {
			errs = append(errs, fmt.Errorf("file %s error %s", f, err.Error()))
			continue
		}
		break
	}

	if cfg == nil {
		return nil, fmt.Errorf("file config parse error %v", errs)
	}

	return cfg, nil
}
