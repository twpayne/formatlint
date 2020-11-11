package main

import (
	"go/parser"
	"go/token"
)

type goLinter struct{}

func (goLinter) Lint(data []byte) error {
	_, err := parser.ParseFile(token.NewFileSet(), "", data, parser.AllErrors)
	return err
}

func (goLinter) Name() string {
	return "go"
}
