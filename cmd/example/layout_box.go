package main

import (
	"github.com/octohelm/ebiten-compose/pkg/layout/alignment"
	"golang.org/x/image/colornames"

	. "github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/compose/modifier"
)

type LayoutAbsolute struct {
}

func (LayoutAbsolute) Build(context BuildContext) VNode {
	return Box(
		modifier.BackgroundColor(colornames.Yellowgreen),
		modifier.FillMaxSize(),
		modifier.PaddingAll(20),
	).Children(
		Box(
			modifier.Size(100),
			modifier.BackgroundColor(colornames.Aliceblue),
			modifier.Align(alignment.TopStart),
		),
		Box(
			modifier.Size(100),
			modifier.BackgroundColor(colornames.Aliceblue),
			modifier.Align(alignment.TopEnd),
		),
		Box(
			modifier.Size(100),
			modifier.BackgroundColor(colornames.Aliceblue),
			modifier.Align(alignment.BottomStart)),

		Box(
			modifier.Size(100),
			modifier.BackgroundColor(colornames.Aliceblue),
			modifier.Align(alignment.BottomEnd)),

		Box(
			modifier.Size(100),
			modifier.BackgroundColor(colornames.Beige),
			modifier.Align(alignment.Center),
		),

		Box(
			modifier.Size(100),
			modifier.BackgroundColor(colornames.Beige),
			modifier.Align(alignment.Start),
		),
		Box(
			modifier.Size(100),
			modifier.BackgroundColor(colornames.Beige),
			modifier.Align(alignment.End),
		),

		Box(
			modifier.Size(100),
			modifier.BackgroundColor(colornames.Beige),
			modifier.Align(alignment.Top),
		),

		Box(
			modifier.Size(100),
			modifier.BackgroundColor(colornames.Beige),
			modifier.Align(alignment.Bottom),
		),

		Box(
			modifier.Size(100),
			modifier.BackgroundColor(colornames.Lightblue),
			modifier.Align(alignment.Center),
			modifier.Offset(50),
			modifier.BorderStroke(2, colornames.Black),
			modifier.PaddingAll(10),
			modifier.RoundedAll(50),
		).Children(
			Box(
				modifier.FillMaxSize(),
				modifier.BackgroundColor(colornames.Royalblue),
				modifier.RoundedAll(20),
			),
		),
	)
}
