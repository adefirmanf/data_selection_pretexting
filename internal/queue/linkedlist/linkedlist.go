package linkedlist

import "container/list"

// LinkedList .
type LinkedList struct {
	list *list.List
}

// PushBack .
func (l *LinkedList) PushBack(UserID string) {
	l.list.PushBack(UserID)
}

// PullFront .
func (l *LinkedList) PullFront() string {
	val := l.list.Front().Value.(string)
	l.list.Remove(l.list.Front())
	return val
}

// PeekFront .
func (l *LinkedList) PeekFront() string {
	val := l.list.Front().Value.(string)
	return val
}

// Size .
func (l *LinkedList) Size() int {
	return l.list.Len()
}

// NewLinkedList .
func NewLinkedList() *LinkedList {
	return &LinkedList{
		list: list.New().Init(),
	}
}
