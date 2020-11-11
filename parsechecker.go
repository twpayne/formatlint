package main

// A ParseChecker checks that data parses.
type ParseChecker interface {
	Name() string
	Parse(data []byte) error
}
