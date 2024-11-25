package main

import (
	"flag"
	"log"
	"os"

	"github.com/coryjamescrook/templater/internal/template"
)

func main() {
	templateName := flag.String("t", "", "the name of the template to build from")
	outDir := flag.String("o", "", "the path to the output directory to build the template")

	flag.Parse()

	if templateName == nil || *templateName == "" {
		panic("`template name` is required")
	}

	if outDir == nil || *outDir == "" {
		panic("`output directory` is required")
	}

	log.Printf("Initializing template: `%s`\n", *templateName)
	t := template.CreateTemplate(*templateName)

	log.Println("Collecting template data...")
	t.CollectData()

	log.Printf("Beginning template build to `%s`\n", *outDir)
	t.Build(*outDir)

	log.Println("Template build successful!")
	os.Exit(0)
}
