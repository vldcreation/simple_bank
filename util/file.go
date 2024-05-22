package util

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFromYAML(fname string, v interface{}) error {
	b, err := os.ReadFile(fname)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(b, v)
}
