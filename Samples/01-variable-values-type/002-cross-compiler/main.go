package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime: %s\narchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}
