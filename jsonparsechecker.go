package main

import "encoding/json"

type jsonParseChecker struct{}

func (jsonParseChecker) Parse(data []byte) error {
	var value interface{}
	return json.Unmarshal(data, &value)
}

func (jsonParseChecker) Name() string {
	return "json"
}
