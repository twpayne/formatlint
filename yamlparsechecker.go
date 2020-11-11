package main

import "gopkg.in/yaml.v3"

type yamlParseChecker struct{}

func (yamlParseChecker) Parse(data []byte) error {
	var value interface{}
	return yaml.Unmarshal(data, &value)
}

func (yamlParseChecker) Name() string {
	return "yaml"
}
