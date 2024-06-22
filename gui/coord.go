package gui

import (
	"image"
	"math"

	g "github.com/AllenDang/giu"
)

func pointInCircle(x float32, y float32, cx float32, cy float32, r float32) bool {
	d := math.Sqrt(math.Pow(float64(cx-x), 2) + math.Pow(float64(cy-y), 2))
	return d <= float64(r)
}

func mouseInCircle(cx float32, cy float32, r float32) bool {
	pos := g.GetMousePos()
	return pointInCircle(float32(pos.X), float32(pos.Y), cx, cy, r)
}

func mouseInCirclePt(pt image.Point, r float32) bool {
	return mouseInCircle(float32(pt.X), float32(pt.Y), r)
}
