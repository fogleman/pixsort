package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fogleman/pixsort/pixsort"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Usage: pixsort image.png [quality]")
		return
	}
	path := args[0]
	quality := 24
	if len(args) == 2 {
		q, err := strconv.ParseInt(args[1], 0, 0)
		if err != nil || q < 0 || q > 30 {
			fmt.Println("Quality value must be between 0 and 30")
			return
		}
		quality = int(q)
	}
	pixsort.Run(path, quality)
}
