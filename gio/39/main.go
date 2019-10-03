package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/op"
	"gioui.org/op/paint"
)

var images [2]*image.RGBA
var Image *image.RGBA // This image is modified then uploaded to the window buffer at each click.

func main() {
	go func() {
		Image = image.NewRGBA(image.Rectangle{})
		w := app.NewWindow(app.WithTitle("gio39"))
		ops := new(op.Ops)
		for e := range w.Events() {
			switch e := e.(type) {
			case app.UpdateEvent:
				fmt.Println("update")
				ops.Reset()
				Draw(ops, e.Size)
				w.Update(ops)
			case pointer.Event:
				if e.Type == pointer.Press {
					fmt.Println("press")
					w.Invalidate()
				}
			}
		}
	}()
	app.Main()
}

func Draw(ops *op.Ops, size image.Point) {
	if Image.Bounds().Max != size {
		fmt.Println("Resize ", size)
		images[0] = image.NewRGBA(image.Rectangle{Max: size}) // reallocate backing image with new size
		images[1] = image.NewRGBA(image.Rectangle{Max: size})
	}
	images[0], images[1] = images[1], images[0]
	Image = images[0]
	c := drawNextFrame() // modifies Image: it changes the color of the visible rectangle on each call.
	fmt.Println("Draw, color:", c)
	paint.ImageOp{Image, Image.Bounds()}.Add(ops)
	paint.PaintOp{Rect: f32.Rectangle{Max: f32.Point{float32(size.X), float32(size.Y)}}}.Add(ops)
}

var frameIndex int
var red, green, blue = color.RGBA{255, 0, 0, 255}, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255}

// DrawNextFrame paints the upper left quadrant of the backing image red, green or blue.
func drawNextFrame() color.Color {
	colors := []color.RGBA{red, green, blue}
	c := colors[frameIndex%3]
	r := Image.Bounds()
	r.Max.X /= 2
	r.Max.Y /= 2
	draw.Draw(Image, r, &image.Uniform{c}, image.ZP, draw.Src)
	frameIndex++
	return c
}
