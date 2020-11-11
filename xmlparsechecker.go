package main

import "encoding/xml"

type xmlParseChecker struct{}

func (xmlParseChecker) Parse(data []byte) error {
	var value interface{}
	return xml.Unmarshal(data, &value)
}

func (xmlParseChecker) Name() string {
	return "xml"
}
