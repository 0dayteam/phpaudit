package audit

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"phpaudit/finder"
	"syscall"
)

func Run(ctx context.Context, conf ScanConfig) {
	files := finder.FindCode(ctx, conf.FindConfig)

	exitSign := make(chan os.Signal, 1)
	signal.Notify(exitSign,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	for file := range files {
		go ParseFile(*file)
	}

	for {
		select {
		case <-exitSign:
			fmt.Print("exit")
			syscall.Exit(0)
		case <-ctx.Done():
			syscall.Exit(0)
		}
	}

}
