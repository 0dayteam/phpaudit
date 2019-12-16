package audit

import (
	"context"
	"os"
	"os/signal"
	"phpaudit/config"
	"phpaudit/finder"
	"syscall"
)

func Run(ctx context.Context, conf config.RunConfig) {
	log.Debug("start scan file")
	files := finder.FindCode(ctx, conf.FinderConf)

	exitSign := make(chan os.Signal, 1)
	signal.Notify(exitSign,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	log.Debug("start audit")
	for file := range files {

		go ParseFile(*file)
	}

	for {
		select {
		case <-exitSign:
			log.Info("exit; ok")
			syscall.Exit(0)
		case <-ctx.Done():
			syscall.Exit(0)
		}
	}

}
