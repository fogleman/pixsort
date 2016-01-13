package main

import (
	"os"

	"github.com/fogleman/pixsort/pixsort"
)

func main() {
	pixsort.Run(os.Args[1])
}
