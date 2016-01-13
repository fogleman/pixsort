package pixsort

import (
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"os"
)

func LoadPNG(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return png.Decode(file)
}

func GetPoints(im image.Image) (int, int, []Point) {
	var result []Point
	bounds := im.Bounds()
	w := bounds.Size().X
	h := bounds.Size().Y
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := im.At(x, y)
			r, _, _, _ := c.RGBA()
			if r < 128 {
				result = append(result, Point{x, y})
			}
		}
	}
	return w, h, result
}

func CreateFrame(m, w, h int, points []Point) *image.Paletted {
	var palette []color.Color
	palette = append(palette, color.RGBA{255, 255, 255, 255})
	palette = append(palette, color.RGBA{0, 0, 0, 255})
	im := image.NewPaletted(image.Rect(0, 0, w*m, h*m), palette)
	for _, point := range points {
		for y := 0; y < m; y++ {
			for x := 0; x < m; x++ {
				im.SetColorIndex(point.X*m+x, point.Y*m+y, 1)
			}
		}
	}
	return im
}

func SaveGIF(path string, m, w, h int, points []Point) error {
	g := gif.GIF{}
	g.Image = append(g.Image, CreateFrame(m, w, h, nil))
	g.Delay = append(g.Delay, 10)
	for i := range points {
		g.Image = append(g.Image, CreateFrame(m, w, h, points[:i+1]))
		g.Delay = append(g.Delay, 10)
	}
	g.Image = append(g.Image, CreateFrame(m, w, h, points))
	g.Delay = append(g.Delay, 100)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return gif.EncodeAll(file, &g)
}
