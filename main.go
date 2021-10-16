package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bytixo/Asabira/logger"
	"github.com/bytixo/Asabira/watcher"
)

func main() {
	watcher.Start()

	// prevent the application from closing, until we receive ctrl+c
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	logger.Info("Closing Asabira")
}
