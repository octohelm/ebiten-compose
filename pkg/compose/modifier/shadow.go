package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/paint"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

func Shadow(elevation unit.Dp) compose.Modifier {
	return &shadowModifier{
		elevation: elevation,
	}
}

type shadowModifier struct {
	elevation unit.Dp
}

func (m *shadowModifier) Modify(ctx context.Context, w compose.Element) bool {
	if setter, ok := w.(paint.ShadowSetter); ok {
		return setter.SetShadow(m.elevation)
	}
	return false
}
