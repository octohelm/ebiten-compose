package compose

import (
	"context"

	"github.com/octohelm/ebiten-compose/pkg/compose/internal"
)

// H Create VNode from Component
func H(component Component, modifiers ...Modifier) VNode {
	return internal.H(component, modifiers...)
}

// Fragment children without Element wrapper
func Fragment(vnodes ...VNode) VNode {
	return H(internal.Fragment{}).Children(vnodes...)
}

// Provider inject something into context for child nodes
func Provider(c func(ctx context.Context) context.Context) VNode {
	return H(internal.Provider(c))
}
