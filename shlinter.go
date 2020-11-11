package main

import (
	"bytes"

	"mvdan.cc/sh/v3/syntax"
)

type shLinter struct{}

func (shLinter) Lint(data []byte) error {
	_, err := syntax.NewParser().Parse(bytes.NewReader(data), "")
	return err
}

func (shLinter) Name() string {
	return "sh"
}
