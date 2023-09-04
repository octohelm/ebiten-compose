package modifier

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/paint"
	"github.com/octohelm/ebiten-compose/pkg/paint/size"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

func Height(dp unit.Dp) compose.Modifier {
	return &sizeModifier{v: dp, sizes: []size.Size{
		size.Height,
	}}
}

func Width(dp unit.Dp) compose.Modifier {
	return &sizeModifier{v: dp, sizes: []size.Size{
		size.Width,
	}}
}

func Size(dp unit.Dp) compose.Modifier {
	return &sizeModifier{v: dp, sizes: []size.Size{
		size.Width,
		size.Height,
	}}
}

type sizeModifier struct {
	v     unit.Dp
	sizes []size.Size
}

func FillMaxHeight() compose.Modifier {
	return &sizeModifier{v: -1, sizes: []size.Size{
		size.Height,
	}}
}

func FillMax() compose.Modifier {
	return &sizeModifier{v: -1, sizes: []size.Size{
		size.Width,
		size.Height,
	}}
}

func FillMaxWidth() compose.Modifier {
	return &sizeModifier{v: -1, sizes: []size.Size{
		size.Width,
	}}
}

func FillMaxSize() compose.Modifier {
	return &sizeModifier{v: -1, sizes: []size.Size{
		size.Width,
		size.Height,
	}}
}

func (r *sizeModifier) Modify(ctx context.Context, w compose.Element) bool {
	if r.v > -1 {
		if setter, ok := w.(paint.SizeSetter); ok {
			return setter.SetSize(r.v, r.sizes...)
		}
		return false
	}

	if setter, ok := w.(paint.FillMaxSetter); ok {
		return setter.SetFillMax(r.sizes...)
	}

	return false
}
