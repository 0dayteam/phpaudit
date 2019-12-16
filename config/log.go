package config

type LogConfig struct {
	LogType  string `yaml:"log_type"`
	LogPath  string `yaml:"log_path"`
	LogLevel uint32 `yaml:"log_level"`
}
