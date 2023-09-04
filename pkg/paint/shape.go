package paint

import (
	"image"

	"github.com/octohelm/ebiten-compose/pkg/canvas"

	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Shape interface {
	Rectangle(c canvas.Canvas) image.Rectangle
	Path(c canvas.Canvas) *vector.Path
}

type SizedShape interface {
	Shape

	SizeSetter
	SizedChecker
	FillMaxSetter
	FillMinSetter
}
