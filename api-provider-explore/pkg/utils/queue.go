package utils

type Queue[T any] []T

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	if len(*q) == 0 {
		panic("Queue is empty")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *Queue[T]) Remove(index int) {
	if len(*q) == 0 {
		panic("Queue is empty")
	}
	*q = append((*q)[:index], (*q)[index+1:]...)
}
