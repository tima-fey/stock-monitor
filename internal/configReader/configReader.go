package configReader

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Token string `yaml:"token"`
}

func (c *Config) Parse(path string) (err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(data, &c); err != nil {
		return
	}
	return err
}

