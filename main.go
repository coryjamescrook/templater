package main

import (
	"log"
	"os"

	"github.com/coryjamescrook/templater/internal/templaters"
)

func main() {
	args := os.Args[1:]
	templateName := args[0]
	if templateName == "" {
		panic("you must provide a valid template name to build from")
	}

	outDir := args[1]
	if outDir == "" {
		panic("you must provide a valid output directory to build to")
	}

	tr := templaters.Templater{}

	log.Printf("Initializing template: `%s`\n", templateName)
	t := tr.CreateTemplate(templateName)

	log.Println("Collecting template data...")
	t.CollectData()

	log.Printf("Beginning template build to `%s`\n", outDir)
	t.Build(outDir)
}
