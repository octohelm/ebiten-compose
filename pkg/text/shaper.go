package text

import (
	"context"

	"gioui.org/font/gofont"
	"gioui.org/text"
)

var DefaultShaper = text.NewShaper(text.WithCollection(gofont.Collection()))

type shaperContext struct{}

func ContextWithShaper(ctx context.Context, s *text.Shaper) context.Context {
	return context.WithValue(ctx, shaperContext{}, s)
}

func ShaperFromContext(ctx context.Context) *text.Shaper {
	if s, ok := ctx.Value(shaperContext{}).(*text.Shaper); ok {
		return s
	}
	return DefaultShaper
}
