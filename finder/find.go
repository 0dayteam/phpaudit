package finder

import (
	"context"
	"log"
	"phpaudit/config"
	"phpaudit/finder/find"
)

func FindCode(ctx context.Context, config config.FinderConf) <-chan *find.File {
	root := find.WithRoot(config.RootPath)
	names := find.Find(ctx,
		root,
		find.WithRegexPathMatch(".+\\.php$"),
		find.WithoutIgnoreDir(config.IgnoreDir),
		find.WithDebugLog(log.Printf),
	)
	return names
}
