// Program 25671 demonstrates a bug in shiny/windriver (#25671)
//
// The program creates a main window (blue).
// For any click into the window and a green client window appears.
// Pressing the ESC key asks shiny to close the current window
// by sending a lifecycle event.
// The event loop exits and calls the deferred Release method on the window.
// If this is done on the main window, it works as expected.
// However if it is done on a client window, it does not disappear.
// It is left in an unusable state (try to make it bigger), as the event loop is
// not running anymore.
package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
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
				log.Printf("received %v: window %p should be closing now", e, w)
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

		case key.Event:
			// We ask to close the window when the ESC key is pressed.
			if e.Code == key.CodeEscape && e.Direction == key.DirPress {
				log.Printf("ESC: send lifecycle.StageDead to window %p", w)
				e := lifecycle.Event{To: lifecycle.StageDead}
				w.Send(e)
			}

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
