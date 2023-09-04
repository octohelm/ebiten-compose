package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/layout/direction"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/layout"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

func Offset(dp unit.Dp) compose.Modifier {
	return &offsetModifier{
		dp: dp,
		directions: []direction.Direction{
			direction.X,
			direction.Y,
		},
	}
}

func OffsetX(dp unit.Dp) compose.Modifier {
	return &offsetModifier{
		dp: dp,
		directions: []direction.Direction{
			direction.X,
		},
	}
}

func OffsetY(dp unit.Dp) compose.Modifier {
	return &offsetModifier{
		dp: dp,
		directions: []direction.Direction{
			direction.Y,
		},
	}
}

type offsetModifier struct {
	dp         unit.Dp
	directions []direction.Direction
}

func (e *offsetModifier) Modify(ctx context.Context, w compose.Element) bool {
	if setter, ok := w.(layout.OffsetSetter); ok {
		return setter.SetOffset(e.dp, e.directions...)
	}
	return false
}
