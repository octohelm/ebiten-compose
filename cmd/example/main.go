package main

import (
	"context"

	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/compose/renderer"
)

func main() {
	ebiten.SetWindowSizeLimits(1200, 720, -1, -1)

	r := renderer.CreateRoot()

	// could inject singletons in to context
	ctx := context.Background()

	r.Render(ctx, H(BoxShadow{}))

	if err := r.Wait(); err != nil {
		panic(err)
	}
}
