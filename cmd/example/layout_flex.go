package main

import (
	. "github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/compose/modifier"
	"github.com/octohelm/ebiten-compose/pkg/layout/alignment"
	"github.com/octohelm/ebiten-compose/pkg/layout/arrangement"
	"golang.org/x/image/colornames"
)

type LayoutFlex struct {
}

func (LayoutFlex) Build(context BuildContext) VNode {
	//list := MapIndexed(make([]VNode, 1000), func(e VNode, idx int) VNode {
	//	return Box(modifier.FillMaxWidth(), modifier.Height(40), modifier.BorderStroke(1, color.Black)).Children(
	//		//Text(fmt.Sprint(idx)),
	//	)
	//})

	return Row(modifier.FillMaxSize()).Children(
		Column(
			modifier.VerticalScroll(),
			modifier.FillMaxHeight(),
			modifier.Width(200),
		),
		Column(modifier.FillMaxHeight(), modifier.BackgroundColor(colornames.Beige)).Children(
			Box(modifier.FillMaxWidth(), modifier.BackgroundColor(colornames.Black), modifier.Height(40)),
			Row(modifier.FillMaxWidth(), modifier.Weight(0.5), modifier.PaddingAll(10), modifier.BackgroundColor(colornames.Royalblue), modifier.FillMaxSize()),
			Column(modifier.FillMaxWidth(), modifier.Weight(3), modifier.Gap(10), modifier.PaddingAll(10), modifier.BackgroundColor(colornames.Rosybrown), modifier.FillMaxSize()).Children(
				H(LayoutRow{Arrangement: arrangement.EqualWeight, Alignment: alignment.Start}),
				H(LayoutRow{Arrangement: arrangement.SpaceBetween, Alignment: alignment.Middle}),
				H(LayoutRow{Arrangement: arrangement.SpaceAround, Alignment: alignment.End}),
				H(LayoutRow{Arrangement: arrangement.SpaceEvenly, Alignment: alignment.Middle}),
				H(LayoutRow{Arrangement: arrangement.End, Alignment: alignment.Middle}),
				H(LayoutRow{Arrangement: arrangement.Center, Alignment: alignment.Middle}),
				H(LayoutRow{Arrangement: arrangement.Start, Alignment: alignment.Middle}),
			),
		),
	)
}

type LayoutRow struct {
	Arrangement arrangement.Arrangement
	Alignment   alignment.Alignment
}

func (f LayoutRow) Build(context BuildContext) VNode {
	return Row(
		modifier.Align(f.Alignment),
		modifier.Arrangement(f.Arrangement),
		modifier.Gap(10),
		modifier.PaddingAll(10),
		modifier.BackgroundColor(colornames.Aliceblue),
		modifier.FillMaxWidth(),
		modifier.FillMaxHeight(),
	).Children(
		Box(
			modifier.BackgroundColor(colornames.Lightpink),
			modifier.When(f.Arrangement != arrangement.EqualWeight, modifier.Width(80)),
			modifier.Height(10),
		),
		Box(
			modifier.BackgroundColor(colornames.Lightcoral),
			modifier.When(f.Arrangement != arrangement.EqualWeight, modifier.Width(80)),
			modifier.Height(10),
		),
		Box(
			modifier.BackgroundColor(colornames.Lightblue),
			modifier.When(f.Arrangement != arrangement.EqualWeight, modifier.Width(80)),
			modifier.Height(10),
		),
	)
}
