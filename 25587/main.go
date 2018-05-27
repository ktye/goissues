// Program 25576 demonstrates a bug in shiny/windriver (#25576)
//
// The program creates a main window (blue).
// For any click into the window and a green client window appears.
// Trying to close the client window does not make it disappear,
// but leaves it around unfunctinally:
// It does not redraw when resizing (try to make it bigger).
// Closing the main window always works as expected.
package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

func main() {
	driver.Main(func(s screen.Screen) {
		createWindow(s, false)
	})
}

func createWindow(s screen.Screen, isClient bool) {

	title := "Main Window"
	if isClient {
		title = "Client Window"
	}
	opt := screen.NewWindowOptions{Width: 300, Height: 300, Title: title}

	w, err := s.NewWindow(&opt)
	if err != nil {
		log.Print(err)
		return
	}
	defer w.Release()

	var b screen.Buffer
	defer func() {
		if b != nil {
			b.Release()
		}
	}()

	for {
		switch e := w.NextEvent().(type) {

		case lifecycle.Event:
			if e.To == lifecycle.StageDead {
				log.Print(e)
				return
			}

		case paint.Event:
			// lock()
			w.Upload(image.Point{}, b, b.Bounds())
			w.Publish()
			// unlock()

		case size.Event:
			// lock()
			if b != nil {
				b.Release()
			}
			b, err = s.NewBuffer(e.Size())
			if err != nil {
				log.Print(err)
			}

			// Paint the main window blue, and client windows green.
			c := color.RGBA{0, 0, 255, 0}
			if isClient {
				c = color.RGBA{0, 255, 0, 255}
			}
			m := b.RGBA()
			draw.Draw(m, m.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
			// unlock

		case mouse.Event:
			if isClient == false {
				if e.Button > 0 {
					if e.Direction == mouse.DirPress {
						log.Print("Create a new window")
						go createWindow(s, true)
					}
				}
			}
		}
	}
}
