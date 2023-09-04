package renderer

import (
	"github.com/octohelm/ebiten-compose/pkg/compose/internal"
)

func Portal(target internal.Element) internal.VNode {
	n := internal.H(internal.Portal{})
	n.(internal.VNodeAccessor).Mount(target)
	return n
}
