package app

import (
	"fmt"
	"log"
	"sync"

	"github.com/vldcreation/simple_bank/consts"
	"github.com/vldcreation/simple_bank/util"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

var (
	once sync.Once
	_cfg *Config
)

func NewConfig() *Config {
	fpath := []string{consts.ConfigPath}
	once.Do(func() {
		c, err := readCfg("env.yaml", fpath...)
		if err != nil {
			log.Fatal(err)
		}

		_cfg = c
	})

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
