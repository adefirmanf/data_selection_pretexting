package queue

// Queue .
type Queue struct {
	repository I
}

// I .
type I interface {
	PushBack(UserID string)
	PullFront(UserID string)
	PeekFront() string
	Count() int
}

// PushBack .
func (q *Queue) PushBack(UserID string) {
	q.repository.PushBack(UserID)
}

// PullFront .
func (q *Queue) PullFront(UserID string) {
	q.repository.PullFront(UserID)
}

// PeekFront .
func (q *Queue) PeekFront() string {
	return q.repository.PeekFront()
}

// Count .
func (q *Queue) Count() int {
	return q.repository.Count()
}

// NewQueue .
func NewQueue(repository I) *Queue {
	return &Queue{
		repository: repository,
	}
}
