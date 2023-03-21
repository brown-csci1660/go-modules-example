package main

import (
	"examplelib" // After adding this, run "go get examplelib" from this dir
)

func main() {
	// Call something from the library
	examplelib.RunCommand("/tmp/out", []string{"ls", "-la"})
}
