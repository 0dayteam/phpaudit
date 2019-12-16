package config

import "github.com/go-yaml/yaml"

type RunConfig struct {
	LogConfig  `yaml:"log_config"`
	FinderConf `yaml:"finder_config"`
}

func NewRunConfig(data []byte) (*RunConfig, error) {
	c := &RunConfig{}
	err := yaml.Unmarshal([]byte(data), c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
