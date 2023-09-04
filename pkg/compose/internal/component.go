package internal

import "context"

type Component interface {
	Build(BuildContext) VNode
}

type BuildContext interface {
	context.Context

	ChildVNodes() []VNode
}

type BuildContextAccessor interface {
	VNode() VNodeAccessor
}
