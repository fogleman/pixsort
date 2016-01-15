package pixsort

import (
	"image"
	"image/color"
	"image/color/palette"
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
			r, g, b, a := c.RGBA()
			if a < 255 {
				continue
			}
			if r == 255 && g == 255 && b == 255 {
				continue
			}
			result = append(result, Point{x, y, byte(r), byte(g), byte(b)})
		}
	}
	return w, h, result
}

func CreateFrame(m, w, h int, points []Point) *image.Paletted {
	im := image.NewPaletted(image.Rect(0, 0, w*m, h*m), palette.Plan9)
	index := color.Palette(palette.Plan9).Index(color.RGBA{255, 255, 255, 255})
	for y := 0; y < h*m; y++ {
		for x := 0; x < w*m; x++ {
			im.SetColorIndex(x, y, uint8(index))
		}
	}
	for _, p := range points {
		for y := 0; y < m; y++ {
			for x := 0; x < m; x++ {
				c := color.RGBA{p.R, p.G, p.B, 255}
				im.Set(p.X*m+x, p.Y*m+y, c)
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
