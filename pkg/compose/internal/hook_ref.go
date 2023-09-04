package internal

import (
	"fmt"
)

type RefHook[T any] struct {
	Ref[T]
}

func (RefHook[T]) UpdateHook(next Hook) {
}

func (s *RefHook[T]) String() string {
	return fmt.Sprintf("UseRef: %v", s.Current)
}

type Ref[T any] struct {
	Current T
}
