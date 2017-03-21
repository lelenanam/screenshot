package screenshot

import (
	"image"
	"image/color"
	// Packages image/gif, image/jpeg and image/png are not used explicitly in the code below,
	// but are imported for its initialization side-effect, which allows
	// image.Decode to understand these formatted images.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

// MinColorPercent is a lower bound for one particular color in percent to total number of colors in image
var MinColorPercent = 0.2

// MaxColorPercent is a upper bound for one particular color in percent to total number of colors in image
var MaxColorPercent = 0.4

// MinColorMapSize is a lower bound for size of color map
var MinColorMapSize = 1000

//Detect reports whether an Image m is screenshots
func Detect(m image.Image) bool {
	res := false
	colormap := make(map[color.Color]int)
	bounds := m.Bounds()
	total := (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			color := m.At(x, y)
			colormap[color]++
		}
	}
	colormapSize := len(colormap)

	for c, n := range colormap {
		percent := float64(n) / float64(total)
		if percent > MaxColorPercent {
			log.Printf("Map size = %d; Color:%v, percent = %.4f", colormapSize, c, percent)
			return true
		}
		if percent > MinColorPercent && colormapSize < MinColorMapSize {
			log.Printf("Map size = %d; Color:%v, percent = %.4f", colormapSize, c, percent)
			return true
		}
	}
	return res
}
