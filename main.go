package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func addWatchers(watcher *fsnotify.Watcher, dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
}

func runTests(dir string) {
	fmt.Printf("Running test cases...\n")
	cmd := exec.Command("go", "test", dir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error running tests: %v\n", err)
	}
	fmt.Printf("Test output:\n%s\n", string(output))
}

func startWatcher(watcher *fsnotify.Watcher, dir string) {
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
					runTests(dir) // Run tests for the modified file
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	const DIR = "./exercises"
	const TEST_DIR = DIR + "/..."

	startWatcher(watcher, TEST_DIR)

	fmt.Println("Watching for file changes...")
	err = addWatchers(watcher, DIR)
	if err != nil {
		log.Fatal(err)
	}

	runTests(TEST_DIR) // Run all tests initially

	<-make(chan struct{})

}
