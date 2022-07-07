package main

import (
	"github.com/shalldie/gos3/internal/upload"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	upload.Setup()
}
