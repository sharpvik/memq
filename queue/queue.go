package queue

type Queue[T any] struct {
	ch     chan T
	peek   T
	peeked bool
}

func New[T any](cap uint) *Queue[T] {
	return &Queue[T]{
		ch: make(chan T, cap),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.ch <- item
}

func (q *Queue[T]) Peek() T {
	if q.peeked {
		return q.peek
	}

	q.peek = <-q.ch
	q.peeked = true
	return q.peek
}

func (q *Queue[T]) Flush() {
	q.peeked = false
}
