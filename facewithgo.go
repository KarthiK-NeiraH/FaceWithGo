package main

import (
	"github.com/fogleman/gg"
)

func main() {
	// create a new image context with a size of 512x512
	dc := gg.NewContext(512, 512)

	// set the background color to light gray
	dc.SetRGB(0.9, 0.9, 0.9)
	dc.Clear()

	// draw the face outline
	dc.SetRGB(0, 0, 0)
	dc.DrawEllipse(256, 256, 200, 250)
	dc.Fill()

	// draw the eyes
	dc.SetRGB(1, 1, 1)
	dc.DrawCircle(200, 200, 40)
	dc.DrawCircle(312, 200, 40)
	dc.Fill()

	// draw the pupils
	dc.SetRGB(0, 0, 0)
	dc.DrawCircle(200, 200, 20)
	dc.DrawCircle(312, 200, 20)
	dc.Fill()

	// draw the mouth
	dc.SetRGB(1, 0, 0)
	dc.DrawEllipse(256, 350, 80, 40)
	dc.Fill()

	// save the image to a file
	dc.SavePNG("face.png")
}
