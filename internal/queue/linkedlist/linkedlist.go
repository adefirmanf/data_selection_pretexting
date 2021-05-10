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

// PushFront .
func (l *LinkedList) PushFront(UserID string) {
	l.list.PushFront(UserID)
}

// PeekFront .
func (l *LinkedList) PeekFront() string {
	val := l.list.Front().Value.(string)
	return val
}

// Count .
func (l *LinkedList) Count() int {
	return l.list.Len()
}

// NewLinkedList .
func NewLinkedList() *LinkedList {
	return &LinkedList{
		list: list.New().Init(),
	}
}
