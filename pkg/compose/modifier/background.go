package modifier

import (
	"context"
	"image/color"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/paint"
)

func BackgroundColor(c color.Color) compose.Modifier {
	return &backgroundColorModifier{c: c}
}

type backgroundColorModifier struct {
	c color.Color
}

func (f *backgroundColorModifier) Modify(ctx context.Context, w compose.Element) bool {
	if setter, ok := w.(paint.FillSetter); ok {
		return setter.SetFill(f.c)
	}
	return false
}
