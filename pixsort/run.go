package pixsort

import "fmt"

func Run(path string) {
	im, err := LoadPNG(path)
	if err != nil {
		panic(err)
	}
	w, h, points := GetPoints(im)
	fmt.Println(len(points))
	model := NewModel(points)
	fmt.Println(model.Energy())
	maxTemp := 100.0
	minTemp := 0.1
	steps := 100000000
	model = Anneal(model, maxTemp, minTemp, steps).(*Model)
	fmt.Println(model.Energy())
	SaveGIF(path+".gif", 8, w, h, model.Points)
}
