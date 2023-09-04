package main

import (
	. "github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/compose/modifier"
	"github.com/octohelm/ebiten-compose/pkg/layout/alignment"
	"github.com/octohelm/ebiten-compose/pkg/layout/arrangement"
	"github.com/octohelm/ebiten-compose/pkg/unit"
	"golang.org/x/image/colornames"
)

type BoxShadow struct{}

func (f BoxShadow) Build(context BuildContext) VNode {
	return Row(
		modifier.Align(alignment.Center),
		modifier.Arrangement(arrangement.SpaceAround),
		modifier.FillMaxWidth(),
		modifier.FillMaxHeight(),
	).Children(
		MapIndexed([]unit.Dp{0, 1, 3, 6, 12, 18, 22}, func(e unit.Dp, i int) VNode {
			return Box(
				modifier.Size(40),
				modifier.RoundedAll(10),
				modifier.Shadow(e),
				modifier.BackgroundColor(colornames.White),
			)
		})...,
	)
}
