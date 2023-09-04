package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/layout"
)

func Arrangement(arrangement layout.Arrangement) compose.Modifier {
	return &arrangementModifier{
		arrangement: arrangement,
	}
}

type arrangementModifier struct {
	arrangement layout.Arrangement
}

func (m *arrangementModifier) Modify(ctx context.Context, w compose.Element) bool {
	if setter, ok := w.(layout.ArrangementSetter); ok {
		return setter.SetArrangement(m.arrangement)
	}
	return false
}
