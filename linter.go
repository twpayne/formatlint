package main

// A Linter lints data.
type Linter interface {
	Name() string
	Lint(data []byte) error
}
