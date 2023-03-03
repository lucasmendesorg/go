/*
	Linked list with Get and Set in Go
		Written by Lucas Mendes <lucas AT lucas@lucasmendes.org>
*/

package main

import "fmt"

type node struct {
	Next *node
	Value interface{}
}

func (n *node) Set(i interface{}) interface{} {
	n.Value = i
	return i
}

func (n *node) Get() interface{} {
	return n.Value
}

func (n *node) PrintValue() {
	fmt.Println(n.Value)
}

func Newnode(value interface{}) *node {
	return &node{ Value: value }
}

type list struct {
	Head *node
	Tail *node
}

func (l *list) RemoveAll() {
	l.Head, l.Tail = nil, nil
}

func (l *list) RemoveByValue(value interface{}) bool {
	n := Newnode(value)
	return l.Remove(n)
}

func (l *list) Remove(n *node) bool {
	var previous *node
	for current := l.Head; current != nil; current = current.Next {
		if n.Value == current.Value {
			fmt.Println("list.Remove(): Removing n.Value =", n.Value)
			if previous == nil {
				l.Head = l.Head.Next
			} else {
				previous.Next = current.Next
			}
			return true
		}
		previous = current
	}
	return false
}

func (l *list) AppendByValue(value interface{}) *node {
	n := Newnode(value)
	return l.Append(n)
}

func (l *list) Append(n *node) *node {
	if l.Head == nil {
		l.Head = n
	}
	if l.Tail == nil {
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
	return n
}

func (l *list) Dump() {
	fmt.Println("list.Dump() {")
	fmt.Println("\tHead =", l.Head)
	fmt.Println("\tTail =", l.Tail)
	for current := l.Head; current != nil; current = current.Next {
		fmt.Println("\t>>>", current.Value)
	} 
	fmt.Println("}")
}

func Newlist() *list {
	return &list{}
}
/*
func main() {
	l := Newlist()
	l.AppendByValue("One")
	l.AppendByValue("Two")
	n3 := l.AppendByValue("Three")
	n3.Set("Three-Dot-Five")
	fmt.Println("n3.Get() =", n3.Get().(string))
	l.AppendByValue("Four")
	l.AppendByValue("Five")
	l.Dump()
	l.Remove(n3)
	l.RemoveByValue("One")
	l.RemoveByValue("Five")
	l.Dump()
}
*/
