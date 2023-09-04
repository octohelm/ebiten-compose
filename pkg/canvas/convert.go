package canvas

import (
	"image"

	"golang.org/x/exp/constraints"
)

func XY(p image.Point) (int, int) {
	return p.X, p.Y
}

func XYAs[T constraints.Ordered](p image.Point) (T, T) {
	return T(p.X), T(p.Y)
}
