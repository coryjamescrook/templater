package template

import (
	"log"
	"strconv"
)

func transformInput(inputType string, input string) interface{} {
	var transformed interface{}

	switch inputType {
	case "string":
		transformed = input
	case "boolean":
		t, err := strconv.ParseBool(input)
		if err != nil {
			panic(err)
		}
		transformed = t
	case "integer":
		i, err := strconv.ParseInt(input, 10, 0)
		if err != nil {
			panic(err)
		}
		transformed = i
	default:
		log.Fatalf("%s does not have transform logic configured", inputType)
	}

	return transformed
}
