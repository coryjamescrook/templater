package template

type TemplateDefDataSchemaProperty struct {
	Type     string `yaml:"type"`
	Required bool   `yaml:"required"`
	Default  string `yaml:"default"`
}

type TemplateDefDataSchema struct {
	Properties map[string]TemplateDefDataSchemaProperty `yaml:"properties"`
}

type TemplateDef struct {
	Name       string                `yaml:"name"`
	DataSchema TemplateDefDataSchema `yaml:"data_schema"`
}
