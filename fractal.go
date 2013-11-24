package main

import (
	"image/color"
	"math"
	"math/cmplx"
)

func toCmplx(x, y int, zoom float64, center complex128) complex128 {
	return center + complex(float64(x)/zoom, float64(y)/zoom)
}

func mandelbrot(c complex128, iter int) float64 {
	z := complex(0, 0)
	for i := 0; i < iter; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 1000 {
			return 1000
		}
	}
	return cmplx.Abs(z)
}

func createColorizer(filename string) func(float64) color.Color {
	gradient := CanvasFromFile(filename)
	limit := gradient.Bounds().Size().Y - 1
	return func(mag float64) color.Color {
		m := int(math.Max(math.Min(300*mag, float64(limit)), 1))
		return gradient.At(0, m)
	}
}

func drawFractal(canvas *Canvas, zoom float64, center complex128, colorizer func(float64) color.Color, iter int) {
	size := canvas.Bounds().Size()
	for x := 0; x < size.X; x += 1 {
		for y := 0; y < size.Y; y += 1 {
			c := toCmplx(x-size.X/2, y-size.Y/2, zoom, center)
			mag := mandelbrot(c, iter)
			color := colorizer(mag)
			canvas.Set(x, y, color)
		}
	}
}
