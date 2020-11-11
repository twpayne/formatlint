package main

import "encoding/json"

type jsonLinter struct{}

func (jsonLinter) Lint(data []byte) error {
	var value interface{}
	return json.Unmarshal(data, &value)
}

func (jsonLinter) Name() string {
	return "json"
}
