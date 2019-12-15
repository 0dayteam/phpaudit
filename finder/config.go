package finder

type FindConfig struct {
	ignoreDir  []string
	ignoreFile []string
	rootPath   string
	timeout    int
}
