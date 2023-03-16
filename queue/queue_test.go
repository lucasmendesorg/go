/*
	Unit test for Queue
		Written by Lucas Mendes <lucas AT lucasmendes.org>
*/

package main

import "testing"

func BenchmarkQueuePopFront(b *testing.B) {
	q := NewQueue()
	for t := 0; t < b.N; t++ {
		q.PushBack("Something")
	}
	b.ResetTimer()
	for t := 0; t < b.N; t++ {
		_, _ = q.PopFront()
	}
}

func BenchmarkQueuePushBack(b *testing.B) {
	str := "Something"
	q := NewQueue()
	b.ResetTimer()
	for t := 0; t < b.N; t++ {
		q.PushBack(str)
	}
}

func BenchmarkNewQueue(b *testing.B) {
	for t := 0; t < b.N; t++ {
		NewQueue();
	}
}

func TestQueue(t *testing.T) {
	values := []string {
		"One", "Two", "Three", "Four", "Five",
	}
	q := NewQueue()
	for i := range values {
		q.PushBack(values[i])
	}
	for i := 0; i < len(values); i++ {
		if f, ok := q.PopFront(); ok == true {
			if f != values[i] {
				t.Fatalf("q.PopFront() should be %s. Got %s", values[i], f)
			}
		} else {
			t.Fatalf("q.PopFront() has failed. Expected success")
		}
	}
}
