package compose

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/canvas"

	"github.com/octohelm/ebiten-compose/pkg/compose/internal"
)

type VNode = internal.VNode
type BuildContext = internal.BuildContext
type Component = internal.Component
type ElementPainter = internal.ElementPainter
type Element = internal.Element

type Modifier = internal.Modifier

func ApplyModifiers(ctx context.Context, widget Element, modifiers ...Modifier) (changed bool) {
	for i := range modifiers {
		if modifiers[i].Modify(ctx, widget) {
			changed = true
		}
	}
	return
}

func WithModifiers(modifiers ...Modifier) Modifier {
	return &modifierComposer{
		modifiers: modifiers,
	}
}

type modifierComposer struct {
	modifiers []Modifier
}

func (m *modifierComposer) Modify(ctx context.Context, w Element) (changed bool) {
	return ApplyModifiers(ctx, w, m.modifiers...)
}

func ElementPainterFunc(layout func(cc canvas.Canvas) canvas.Dimensions) ElementPainter {
	return &widgetPainterFunc{
		layout: layout,
	}
}

type widgetPainterFunc struct {
	layout func(cc canvas.Canvas) canvas.Dimensions
}

func (w *widgetPainterFunc) Layout(cc canvas.Canvas) canvas.Dimensions {
	return w.layout(cc)
}
