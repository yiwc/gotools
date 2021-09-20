package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/yiwc/gotools"
)

func TestHello(t *testing.T) {
	fmt.Println("OK")

	api := gotools.GetToolsApi()
	v := api.GetVersion()
	if v != "0.1" {
		os.Exit(1)
	}
}
