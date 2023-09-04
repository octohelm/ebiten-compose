package main

func MapIndexed[E any, T any](list []E, m func(e E, i int) T) []T {
	var out = make([]T, len(list))
	for i := range list {
		out[i] = m(list[i], i)
	}
	return out
}
