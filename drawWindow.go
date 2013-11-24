package main

import (
	"github.com/lxn/walk"
	"image"
	_ "image/png"
)

type drawWindow struct {
	*walk.MainWindow
	paintWidget *walk.CustomWidget
}

func (mw *drawWindow) drawStuff(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
	bounds := mw.paintWidget.ClientBounds()

	bmp, err := walk.NewBitmap(bounds.Size())
	if err != nil {
		panic(err)
	}

	succeeded := false
	defer func() {
		if !succeeded {
			bmp.Dispose()
		}
	}()

	cv := NewCanvas(image.Rect(0, 0, bounds.Width, bounds.Height))
	zoom := 400.0
	center := complex(-0.75, -0.25)
	colorizer := createColorizer("gradient1.png")
	drawFractal(cv, zoom, center, colorizer, 50)

	fractalBmp, err := walk.NewBitmapFromImage(cv)
	if err != nil {
		panic(err)
	}

	brush, err := walk.NewBitmapBrush(fractalBmp)
	if err != nil {
		panic(err)
	}
	defer brush.Dispose()

	if err := canvas.FillRectangle(brush, bounds); err != nil {
		panic(err)
	}

	return nil
}
