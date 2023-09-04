package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/layout"
	"github.com/octohelm/ebiten-compose/pkg/layout/position"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

func PaddingAll(dp unit.Dp) compose.Modifier {
	return &edgeInsetModifier{
		dp: dp,
		positions: []position.Position{
			position.Left, position.Right, position.Top, position.Bottom,
		},
	}
}

func PaddingLeft(dp unit.Dp) compose.Modifier {
	return &edgeInsetModifier{
		dp: dp,
		positions: []position.Position{
			position.Left,
		},
	}
}

func PaddingRight(dp unit.Dp) compose.Modifier {
	return &edgeInsetModifier{
		dp: dp,
		positions: []position.Position{
			position.Right,
		},
	}
}

func PaddingBottom(dp unit.Dp) compose.Modifier {
	return &edgeInsetModifier{
		dp: dp,
		positions: []position.Position{
			position.Bottom,
		},
	}
}

func PaddingTop(dp unit.Dp) compose.Modifier {
	return &edgeInsetModifier{
		dp: dp,
		positions: []position.Position{
			position.Top,
		},
	}
}

func PaddingVertical(dp unit.Dp) compose.Modifier {
	return &edgeInsetModifier{
		dp: dp,
		positions: []position.Position{
			position.Top,
			position.Bottom,
		},
	}
}

func PaddingHorizontal(dp unit.Dp) compose.Modifier {
	return &edgeInsetModifier{
		dp: dp,
		positions: []position.Position{
			position.Left,
			position.Right,
		},
	}
}

type edgeInsetModifier struct {
	dp        unit.Dp
	positions []position.Position
}

func (e *edgeInsetModifier) Modify(ctx context.Context, w compose.Element) bool {
	if setter, ok := w.(layout.EdgeInsetSetter); ok {
		return setter.SetEdgeInset(e.dp, e.positions...)
	}
	return false
}
