package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/layout/position"
	"github.com/octohelm/ebiten-compose/pkg/paint"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

func RoundedAll(dp unit.Dp) compose.Modifier {
	return &cornerRadiusModifier{
		dp: dp,
		positions: []position.Position{
			position.TopLeft,
			position.TopRight,
			position.BottomLeft,
			position.BottomRight,
		},
	}
}

func RoundedLeft(dp unit.Dp) compose.Modifier {
	return &cornerRadiusModifier{
		dp: dp,
		positions: []position.Position{
			position.TopLeft,
			position.BottomLeft,
		},
	}
}

func RoundedRight(dp unit.Dp) compose.Modifier {
	return &cornerRadiusModifier{
		dp: dp,
		positions: []position.Position{
			position.TopRight,
			position.BottomRight,
		},
	}
}

func RoundedBottom(dp unit.Dp) compose.Modifier {
	return &cornerRadiusModifier{
		dp: dp,
		positions: []position.Position{
			position.BottomLeft,
			position.BottomRight,
		},
	}
}

func RoundedTop(dp unit.Dp) compose.Modifier {
	return &cornerRadiusModifier{
		dp: dp,
		positions: []position.Position{
			position.TopLeft,
			position.TopRight,
		},
	}
}

func RoundedTopLeft(dp unit.Dp) compose.Modifier {
	return &cornerRadiusModifier{
		dp: dp,
		positions: []position.Position{
			position.TopLeft,
		},
	}
}

func RoundedBottomLeft(dp unit.Dp) compose.Modifier {
	return &cornerRadiusModifier{
		dp: dp,
		positions: []position.Position{
			position.BottomLeft,
		},
	}
}

func RoundedTopRight(dp unit.Dp) compose.Modifier {
	return &cornerRadiusModifier{
		dp: dp,
		positions: []position.Position{
			position.TopRight,
		},
	}
}

func RoundedBottomRight(dp unit.Dp) compose.Modifier {
	return &cornerRadiusModifier{
		dp: dp,
		positions: []position.Position{
			position.BottomRight,
		},
	}
}

type cornerRadiusModifier struct {
	dp        unit.Dp
	positions []position.Position
}

func (e *cornerRadiusModifier) Modify(ctx context.Context, widget compose.Element) bool {
	if setter, ok := widget.(paint.CornerRadiusSetter); ok {
		return setter.SetCornerRadius(e.dp, e.positions...)
	}
	return false
}
