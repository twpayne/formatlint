package main

import "gopkg.in/yaml.v3"

type yamlLinter struct{}

func (yamlLinter) Lint(data []byte) error {
	var value interface{}
	return yaml.Unmarshal(data, &value)
}

func (yamlLinter) Name() string {
	return "yaml"
}
