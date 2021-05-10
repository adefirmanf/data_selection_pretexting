package queue_test

import (
	"testing"

	"github.com/adefirmanf/data_selection_pretexting/internal/queue"
	"github.com/adefirmanf/data_selection_pretexting/internal/queue/linkedlist"
	"github.com/google/go-cmp/cmp"
)

func TestQueue(t *testing.T) {
	service := linkedlist.NewLinkedList()

	q := queue.NewQueue(service)

	type test struct {
		inputValue        []string
		expectedPeekValue string
		expectedSizeValue int
	}

	tests := []test{
		{
			inputValue:        []string{"ABC", "DEF", "GHI"},
			expectedPeekValue: "GHI",
			expectedSizeValue: 1,
		},
	}

	for _, v := range tests {
		for _, i := range v.inputValue {
			q.PushBack(i)
		}
		if diff := cmp.Diff("ABC", q.PullFront()); diff != "" {
			t.Errorf(diff)
		}
		if diff := cmp.Diff("DEF", q.PullFront()); diff != "" {
			t.Errorf(diff)
		}
		if diff := cmp.Diff("GHI", q.PeekFront()); diff != "" {
			t.Errorf(diff)
		}
		if diff := cmp.Diff(v.expectedSizeValue, q.Size()); diff != "" {
			t.Errorf(diff)
		}

	}
}
