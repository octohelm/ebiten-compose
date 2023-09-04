package internal

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/canvas"

	"github.com/octohelm/ebiten-compose/pkg/node"
)

type ElementPainter interface {
	Layout(cc canvas.Canvas) canvas.Dimensions
}

type Element interface {
	New() Element

	node.Node

	ElementPainter
}

type Modifier interface {
	Modify(ctx context.Context, element Element) bool
}
