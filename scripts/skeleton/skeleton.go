package skeleton

import (
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/MueR/adventofcode.go/util"
)

//go:embed tmpls/*.go
var fs embed.FS

// Run makes a skeleton main.go and main_test.go file for the given day and year
func Run(day, year int) {
	if day > 25 || day <= 0 {
		log.Fatalf("invalid -day value, must be 1 through 25, got %v", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	ts, err := template.ParseFS(fs, "tmpls/*.go")
	if err != nil {
		log.Fatalf("parsing tmpls directory: %s", err)
	}

	mainFilename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d/main.go", year, day))
	testFilename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d/main_test.go", year, day))
	inputFileName := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d/input.txt", year, day))

	err = os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}

	ensureNotOverwriting(mainFilename)
	ensureNotOverwriting(testFilename)
	ensureNotOverwriting(inputFileName)

	mainFile, err := os.Create(mainFilename)
	if err != nil {
		log.Fatalf("creating main.go file: %v", err)
	}
	testFile, err := os.Create(testFilename)
	if err != nil {
		log.Fatalf("creating main_test.go file: %v", err)
	}
	inputFile, err := os.Create(inputFileName)
	if err != nil {
		log.Fatalf("creating input.txt file: %v", err)
	}

	ts.ExecuteTemplate(mainFile, "main.go", nil)
	ts.ExecuteTemplate(testFile, "main_test.go", nil)
	it, err := os.Open(filepath.Join(util.Dirname(), "/tmpls/input.txt"))
	if err != nil {
		log.Fatalf("opening input.txt template: %v", err)
	}
	_, err = io.Copy(inputFile, it)
	if err != nil {
		log.Fatalf("copying input.txt template: %v", err)
	}

	fmt.Printf("templates made for %d-day%d\n", year, day)
}

func ensureNotOverwriting(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatalf("File already exists: %s", filename)
	}
}
