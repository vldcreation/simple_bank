package app

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"github.com/vldcreation/simple_bank/consts"
	"github.com/vldcreation/simple_bank/util"
)

type Config struct {
	DB *DBConfig `yaml:"db" env:"db" mapstructure:"db"`
}

type DBConfig struct {
	Driver   string `yaml:"driver" env:"driver" mapstructure:"db_driver"`
	User     string `yaml:"user" env:"user" mapstructure:"db_user"`
	Password string `yaml:"password" env:"password" mapstructure:"db_password"`
	Host     string `yaml:"host" env:"host" mapstructure:"db_host"`
	Port     string `yaml:"port" env:"port" mapstructure:"db_port"`
	Database string `yaml:"database" env:"database" mapstructure:"db_database"`
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

func NewConfigFromEnv(path string) *Config {
	var (
		_cfg = &Config{}
	)

	if path == "" {
		path = consts.ConfigPath
	}

	log.Printf("path == %s", path)

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to load config %+v", err)
	}

	err = viper.Unmarshal(&_cfg.DB)
	if err != nil {
		log.Fatalf("unable to unmarshall config %+v", err)
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
