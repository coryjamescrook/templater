# Templater

A basic CLI tool for creating new projects with your own templates.

## Getting Started

Copy the appropriate executable file from the `bin` directory to the root of your templates repository.

If you wish to use the default configuration, your templates should be inside the `templates` directory, each template in its own subdirectory. Like so:

```
  /your-repo-dir
    templater (the templater executable)
    /templates
      /go-http-example (project template dir)
        template.yaml (the template definition file)
        go.mod.template (template for your go module file)
        main.go.template (template for your primary application)
```

Let's take a quick look at the `template.yaml` file in our example template from above:

```yaml
name: go-http-example
data_schema:
  properties:
    GoVersion:
      type: "string"
      required: true
      default: "1.22.6"
    ModuleName:
      type: "string"
      required: true
    ProjectName:
      type: "string"
      required: true
      default: "httpservice"
    DefaultHttpPort:
      type: "string"
      required: true
      default: "8080"
```

Each property defined by your template file will be parsed and prompt the CLI user for input for each when creating a new project based on that template. These properties will later be available to your template files.

Project `.template` files are processed and parsed with Go's [text/template module](https://pkg.go.dev/text/template). The template variables you define in your `template.yaml` file will be present in a struct passed to the template while rendering, and can be accessed like in the following example:

```go
// main.go.template
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultHttpPort string = "{{ .DefaultHttpPort }}"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = defaultHttpPort
	}
	httpAddr := ":" + httpPort

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world! You're looking at {{ .ProjectName }}"))
	})

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("error occurred starting server: %v", err)
	}
	fmt.Printf("http server listening at %s\n", httpAddr)
}
```

As you can see in the above example template file, we are accessing the template variables using the `.` notation inside of the `{{  }}` curly brackets. For more details on how to best use go's text templating, reference the [module documentation](https://pkg.go.dev/text/template).

## Versioning

We use [Semantic Versioning](http://semver.org/) for versioning. For the versions
available, see the [tags on this
repository](https://github.com/coryjamescrook/templater/tags).

## Authors

- **Cory James Crook** - Creator -
  [coryjamescrook](https://github.com/coryjamescrook)

## License

This project is licensed under the [MIT](https://choosealicense.com/licenses/mit)
MIT License - see the [LICENSE.md](LICENSE.md) file for
details
