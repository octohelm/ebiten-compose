package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose"
)

type DisplaySetter interface {
	SetDisplayName(name string) bool
}

type DisplayName string

func (a DisplayName) Modify(ctx context.Context, w compose.Element) bool {
	if setter, ok := w.(DisplaySetter); ok {
		return setter.SetDisplayName(string(a))
	}
	return false
}
