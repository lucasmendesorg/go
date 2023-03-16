/*
	Stack in Go
		Written by Lucas Mendes <lucas AT lucasmendes.org>
*/

package main

import "fmt"

type Stack struct {
	Content []string
}

func (s *Stack) Dump() {
	for _, str := range s.Content {
		fmt.Println("Stack.Dump(): str =", str)
	}
}

func (s *Stack) IsEmpty() bool {
	if len(s.Content) == 0 {
		return true
	}
	return false
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() == true {
		return "", false
	}
	current := len(s.Content) - 1
	value := s.Content[current]
	s.Content = s.Content[:current]
	return value, true
}

func (s *Stack) Push(value string) {
	s.Content = append(s.Content, value)
}

func NewStack() *Stack {
	return &Stack{}
}
/*
func main() {
	s := NewStack()
	fmt.Println("s.IsEmpty() ==", s.IsEmpty())
	s.Push("One")
	s.Push("Two")
	s.Push("Three")
	s.Push("Four")
	s.Push("Five")
	s.Dump()
	fmt.Println("s.IsEmpty() ==", s.IsEmpty())
	t, _ := s.Pop()
	fmt.Println("s.IsEmpty() ==", s.IsEmpty())
	fmt.Println("t =", t);
	s.Dump()
}
*/
