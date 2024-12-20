package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type Exercise struct {
	Name      string
	Completed bool
	Path      string
}

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

func runTests(dir string) bool {
	fmt.Printf("Running test cases for %s...\n", dir)
	cmd := exec.Command("go", "test", "./"+dir)
	output, err := cmd.CombinedOutput()
	fmt.Printf("Test output:\n%s\n", string(output))
	return err == nil
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

func getExercises(dir string) ([]Exercise, error) {
	var exercises []Exercise
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() && isExerciseDir(entry.Name()) {
			exercises = append(exercises, Exercise{
				Name:      entry.Name(),
				Completed: false,
				Path:      filepath.Join(dir, entry.Name()),
			})
		}
	}
	return exercises, nil
}

func isExerciseDir(name string) bool {
	// You can customize this based on your exercise naming convention
	// For example: exercise_01, exercise_02, etc.
	return true
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	const DIR = "./exercises"
	const TEST_DIR = DIR + "/..."

	exercises, err := getExercises(DIR)
	if err != nil {
		log.Fatal(err)
	}

	currentExercise := 0
	fmt.Printf("Starting with exercise: %s\n", exercises[currentExercise].Name)

	startWatcher(watcher, exercises[currentExercise].Path)
	err = addWatchers(watcher, exercises[currentExercise].Path)
	if err != nil {
		log.Fatal(err)
	}

	// Run initial test for the first exercise
	if runTests(exercises[currentExercise].Path) {
		fmt.Printf("Exercise %s completed! Press Enter to continue to next exercise...\n",
			exercises[currentExercise].Name)
		// You might want to add user input handling here
	}

	<-make(chan struct{})

}
