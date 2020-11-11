package main

import "github.com/pelletier/go-toml"

type tomlParseChecker struct{}

func (tomlParseChecker) Parse(data []byte) error {
	var value interface{}
	return toml.Unmarshal(data, &value)
}

func (tomlParseChecker) Name() string {
	return "toml"
}
