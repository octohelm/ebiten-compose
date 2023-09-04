package modifier

import (
	"context"
	"image/color"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/paint"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

func BorderStroke(width unit.Dp, c color.Color) compose.Modifier {
	return &borderStrokeModifier{
		width: width,
		color: c,
	}
}

type borderStrokeModifier struct {
	width unit.Dp
	color color.Color
}

func (f *borderStrokeModifier) Modify(ctx context.Context, w compose.Element) bool {
	if setter, ok := w.(paint.BorderStrokeSetter); ok {
		return setter.SetBorderStroke(f.width, f.color)
	}
	return false
}
