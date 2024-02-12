package queue

import "math"

type Queue[T any] struct {
	Priority float64
	Value    T
}

type ByPriority[T any] []Queue[T]

func (a ByPriority[T]) Len() int      { return len(a) }
func (a ByPriority[T]) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPriority[T]) Less(i, j int) bool { return a[i].Priority < a[j].Priority || (math.IsNaN(a[i].Priority) && !math.IsNaN(a[j].Priority)) }
