package main

import "fmt"

//go:generate go run generate_version.go

var Version string

func main() {
	fmt.Printf("App Version: %s\n", Version)
}
