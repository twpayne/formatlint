package main

import "github.com/pelletier/go-toml"

type tomlLinter struct{}

func (tomlLinter) Lint(data []byte) error {
	var value interface{}
	return toml.Unmarshal(data, &value)
}

func (tomlLinter) Name() string {
	return "toml"
}
