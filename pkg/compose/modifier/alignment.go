package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/layout"
)

func Align(alignment layout.Alignment) compose.Modifier {
	return &alignmentModifier{
		alignment: alignment,
	}
}

type alignmentModifier struct {
	alignment layout.Alignment
}

func (m *alignmentModifier) Modify(ctx context.Context, w compose.Element) bool {
	if setter, ok := w.(layout.AlignSetter); ok {
		return setter.SetAlign(m.alignment)
	}
	return false
}
