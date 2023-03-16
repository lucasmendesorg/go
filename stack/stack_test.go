/*
	Unit test for Stack
		Written by Lucas Mendes <lucas AT lucasmendes.org>
*/

package main

import "testing"

func TestStack(t *testing.T) {
	values := []string {
		"One", "Two", "Three", "Four", "Five",
	}
	s := NewStack()
	for _, v := range values {
		s.Push(v)
	}
	for i := len(values) - 1; i >= 0; i-- {
		if e, ok := s.Pop(); ok == true {
			if e != values[i] {
				t.Fatalf("e != values[i]. Expected to be equal (e = %s, values[i] = %s", e, values[i])
			}
		} else {
			t.Fatalf("s.Pop() has failed. Expected success")
		}
	}
}
