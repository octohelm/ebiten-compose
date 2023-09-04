package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/layout"
)

func VerticalScroll() compose.Modifier {
	return &scrollModifier{
		axis: layout.Vertical,
	}
}

func HorizontalScroll() compose.Modifier {
	return &scrollModifier{
		axis: layout.Horizontal,
	}
}

type scrollModifier struct {
	axis layout.Axis
}

func (m *scrollModifier) Modify(ctx context.Context, w compose.Element) bool {
	if setter, ok := w.(layout.ScrollableSetter); ok {
		return setter.SetScrollable(m.axis, true)
	}
	return false
}
