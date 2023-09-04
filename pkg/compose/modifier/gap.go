package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose"

	"github.com/octohelm/ebiten-compose/pkg/layout"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

func Gap(spacing unit.Dp) compose.Modifier {
	return &gapModifier{spacing: spacing}
}

type gapModifier struct {
	spacing unit.Dp
}

func (m *gapModifier) Modify(ctx context.Context, w compose.Element) (changed bool) {
	if s, ok := w.(layout.SpacingSetter); ok {
		return s.SetSpacing(m.spacing)
	}
	return false
}
