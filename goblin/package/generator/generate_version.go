//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	version := fmt.Sprintf("1.0.%s", time.Now().Format("20060102"))

	f, err := os.Create("version.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, `package main

// Version is automatically generated
const Version = "%s"
`, version)
	if err != nil {
		panic(err)
	}
}
