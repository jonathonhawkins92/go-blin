package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var (
	mainFile    = "/Users/jonathon/Development/go-blin/playground.go" // The main file of your application
	watchDirs   = []string{"/Users/jonathon/Development/go-blin"}
	excludeDirs = []string{"vendor", "node_modules"}
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	for _, dir := range watchDirs {
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				for _, excludeDir := range excludeDirs {
					if strings.HasPrefix(path, excludeDir) {
						return filepath.SkipDir
					}
				}
				return watcher.Add(path)
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	cmd := runServer()

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("Modified file:", event.Name)
				if filepath.Ext(event.Name) == ".go" {
					log.Println("Rebuilding...")
					cmd.Process.Kill()
					cmd = runServer()
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error:", err)
		}
	}
}

func runServer() *exec.Cmd {
	mainPath, err := filepath.Abs(mainFile)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("go", "run", mainPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start server: %v", err)
	} else {
		log.Println("Server started")
	}
	return cmd
}
