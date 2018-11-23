package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	apl386 "github.com/ktye/iv/font"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func main() {
	fnt, err := truetype.Parse(apl386.APL385())
	if err != nil {
		log.Fatal(err)
	}
	opt := truetype.Options{
		Size: 36,
		DPI:  75,
	}
	face := truetype.NewFace(fnt, &opt)

	dst := image.NewRGBA(image.Rect(0, 0, 24, 36))
	draw.Draw(dst, dst.Bounds(), image.White, image.Point{}, draw.Over)

	d := font.Drawer{
		Dst:  dst,
		Src:  image.Black,
		Face: face,
		Dot:  fixed.Point26_6{0, face.Metrics().Ascent},
	}
	d.DrawString("≢")

	// Write the glyph to raster.png.
	w, err := os.Create("freetype.png")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	if err := png.Encode(w, dst); err != nil {
		log.Fatal(err)
	}
}
