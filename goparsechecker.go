package main

import (
	"go/parser"
	"go/token"
)

type goParseChecker struct{}

func (goParseChecker) Parse(data []byte) error {
	_, err := parser.ParseFile(token.NewFileSet(), "", data, parser.AllErrors)
	return err
}

func (goParseChecker) Name() string {
	return "go"
}
