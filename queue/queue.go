/*
	Queue in Go
		Written by Lucas Mendes <lucas AT lucasmendes.org>
*/

package main

import "fmt"

type Queue struct {
	Content []string
}

func (q *Queue) Dump() {
	for _, s := range q.Content {
		fmt.Println("Queue.Dump(): s =", s) 
	}
}

func (q *Queue) IsEmpty() bool {
	if len(q.Content) == 0 {
		return true
	}
	return false
}

func (q *Queue) PopFront() (string, bool) {
	if q.IsEmpty() == true {
		return "", false
	}
	value := q.Content[0]
	q.Content = q.Content[1:]
	return value, true
}

func (s *Queue) PushBack(value string) {
	s.Content = append(s.Content, value)
}

func NewQueue() *Queue {
	return &Queue{}
}

func main() {
	q := NewQueue()
	fmt.Println("s.IsEmpty() ==", q.IsEmpty())
	q.PushBack("One")
	q.PushBack("Two")
	q.PushBack("Three")
	q.PushBack("Four")
	q.PushBack("Five")
	q.Dump()
	fmt.Println("q.IsEmpty() ==", q.IsEmpty())
	t, _ := q.PopFront()
	fmt.Println("q.IsEmpty() ==", q.IsEmpty())
	fmt.Println("t =", t);
	q.Dump()
}
