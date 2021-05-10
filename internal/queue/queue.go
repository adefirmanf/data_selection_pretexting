package queue

// Queue .
type Queue struct {
	repository I
}

// I .
type I interface {
	PushBack(UserID string)
	PullFront() string
	PeekFront() string
	Size() int
}

// PushBack .
func (q *Queue) PushBack(UserID string) {
	q.repository.PushBack(UserID)
}

// PullFront .
func (q *Queue) PullFront() string {
	return q.repository.PullFront()
}

// PeekFront .
func (q *Queue) PeekFront() string {
	return q.repository.PeekFront()
}

// Size .
func (q *Queue) Size() int {
	return q.repository.Size()
}

// NewQueue .
func NewQueue(repository I) *Queue {
	return &Queue{
		repository: repository,
	}
}
