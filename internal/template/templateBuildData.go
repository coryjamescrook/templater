package template

import (
	"fmt"
	"log"
)

type TemplateBuildData map[string]interface{}

func BuildDataForTemplate(t *Template) {
	d := TemplateBuildData{}

	// cli prompt to collect data properties
	for propName, propDef := range t.def.DataSchema.Properties {
		// collect input for value
		var input string
		inputMsg := fmt.Sprintf("Enter %s value for %s", propDef.Type, propName)
		if propDef.Default != "" {
			inputMsg += fmt.Sprintf(" (default: %s)", propDef.Default)
		}
		inputMsg += ": "
		fmt.Print(inputMsg)
		fmt.Scanln(&input)

		if input == "" && propDef.Default != "" {
			input = propDef.Default
		}

		if propDef.Required && input == "" {
			log.Fatalf("`%s` is required for the template: %s", propName, t.def.Name)
		}

		// set this value for the template data
		d[propName] = transformInput(propDef.Type, input)
	}
	// once data is collected, set the `data` value on the `Template` instance
	t.data = d
}
