package config

type FinderConf struct {
	IgnoreDir  []string `yaml:"ignore_dir"`
	IgnoreFile []string `yaml:"ignore_file"`
	RootPath   string   `yaml:"root_path,omitempty"`
	Timeout    int      `yaml:"timeout"`
}
