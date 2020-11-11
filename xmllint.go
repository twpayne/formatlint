package main

import "encoding/xml"

type xmlLinter struct{}

func (xmlLinter) Lint(data []byte) error {
	var value interface{}
	return xml.Unmarshal(data, &value)
}

func (xmlLinter) Name() string {
	return "xml"
}
