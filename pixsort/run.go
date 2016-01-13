package pixsort

import (
	"fmt"
	"math"
	"time"
)

func Run(path string, quality int) {
	iterations := int(math.Pow(2, float64(quality)))
	im, err := LoadPNG(path)
	if err != nil {
		panic(err)
	}
	w, h, points := GetPoints(im)
	fmt.Printf("Sorting %d pixels...\n", len(points))
	fmt.Printf("Quality = %d (%d iterations)\n", quality, iterations)
	model := NewModel(points)
	fmt.Printf("Initial Score = %d\n", int(model.Energy()))
	maxTemp := 10.0
	minTemp := 0.1
	start := time.Now()
	model = Anneal(model, maxTemp, minTemp, iterations).(*Model)
	elapsed := time.Since(start).Seconds()
	fmt.Printf("%c[2K", 27)
	fmt.Printf("  Final Score = %d\n", int(model.Energy()))
	fmt.Printf(" Elapsed Time = %.2fs\n", elapsed)
	out := fmt.Sprintf("%s.%d.gif", path, int(model.Energy()))
	SaveGIF(out, 8, w, h, model.Points)
}
