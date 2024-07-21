//go:build build
// +build build

package main

import (
	"fmt"
)

func init() {
	fmt.Println("This code runs only when the 'build' tag is used")
}
