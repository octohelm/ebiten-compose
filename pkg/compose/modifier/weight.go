package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/layout"
)

func Weight(weight float32) compose.Modifier {
	return &weightModifier{weight: weight}
}

type weightModifier struct {
	weight float32
}

func (m weightModifier) Modify(ctx context.Context, w compose.Element) (changed bool) {
	if s, ok := w.(layout.WeightSetter); ok {
		return s.SetWeight(m.weight)
	}
	return false
}
