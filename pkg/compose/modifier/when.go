package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose"
)

func When(when bool, modifiers ...compose.Modifier) compose.Modifier {
	return &whenModifier{
		when: when,
		then: compose.WithModifiers(modifiers...),
	}
}

type whenModifier struct {
	when bool
	then compose.Modifier
}

func (m *whenModifier) Modify(ctx context.Context, w compose.Element) bool {
	if m.when {
		return m.then.Modify(ctx, w)
	}
	return false
}
