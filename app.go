package gotools

import "fmt"

type Tools struct {
	version   string
	developer string
}

var tools = newTools()

func GetToolsApi() *Tools {
	return tools
}

func newTools() *Tools {
	return &Tools{
		version: "0.1", developer: "yiwc",
	}
}

func Greeting() {
	fmt.Println("Welcome Using GO tools by yiwc")
}

func (t *Tools) GetVersion() string {
	return t.version
}
