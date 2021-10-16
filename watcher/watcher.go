package watcher

import (
	//"time"

	"os"
	"path/filepath"
	"strings"

	"github.com/bytixo/Asabira/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/gen2brain/beeep"
)

var (
	watcher *fsnotify.Watcher
)

func Start() {
	path := "C:/Users/Bytix/AppData/Local/Discord/app-1.0.9003/modules/"

	logger.Info("Starting to watch path:", path)

	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	if err := filepath.Walk(path, watchDir); err != nil {
		logger.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//logger.Info("EVENT: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					filename := FileName(event.Name)
					if filename == ".js" || filename == ".asar" {
						logger.Error("IMPORTANT FILE MODIFIED", event)
						err := beeep.Alert("Asabira", "You are potentially infected with a token grabber please reinstall discord asap", "error.png")
						if err != nil {
							logger.Fatal(err)
						}
					}
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Fatal("Watcher error: ", err)
			}
		}
	}()

	<-done
}

func FileName(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n == -1 {
		return s
	}
	return s[n:]
}

//https://gist.github.com/sdomino/74980d69f9fa80cb9d73
func watchDir(path string, fi os.FileInfo, err error) error {

	if fi.IsDir() && fi.Name() == "node_modules" {
		return filepath.SkipDir
	}
	if fi.Mode().IsDir() {
		logger.Info("Watching: ", fi.Name())
		return watcher.Add(path)
	}

	return nil
}
