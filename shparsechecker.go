package main

import (
	"bytes"

	"mvdan.cc/sh/v3/syntax"
)

type shParseChecker struct{}

func (shParseChecker) Parse(data []byte) error {
	_, err := syntax.NewParser().Parse(bytes.NewReader(data), "")
	return err
}

func (shParseChecker) Name() string {
	return "sh"
}
