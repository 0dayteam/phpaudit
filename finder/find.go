package finder

import (
	"context"
	"log"
	"phpaudit/finder/find"
)

func FindCode(ctx context.Context, config FindConfig) <-chan *find.File {
	root := find.WithRoot(config.rootPath)
	names := find.Find(ctx,
		root,
		find.WithRegexPathMatch(".+\\.php$"),
		find.WithoutIgnoreDir(config.ignoreDir),
		find.WithDebugLog(log.Printf),
	)
	return names
}
