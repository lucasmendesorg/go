/*
	Linked List in Go
		Written by Lucas Mendes <lucas AT lucas@lucasmendes.org>
*/

package main

import "fmt"

type Node struct {
	Next *Node
	Value string
}

func (n *Node) PrintValue() {
	fmt.Println(n.Value)
}

func NewNode(value string) *Node {
	return &Node{ Value: value }
}

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) RemoveAll() {
	l.Head, l.Tail = nil, nil
}

func (l *List) RemoveByValue(value string) bool {
	n := NewNode(value)
	return l.Remove(n)
}

func (l *List) Remove(n *Node) bool {
	var previous *Node
	for current := l.Head; current != nil; current = current.Next {
		if n.Value == current.Value {
			fmt.Println("List.Remove(): Removing n.Value =", n.Value)
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

func (l *List) Append(n *Node) *Node {
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

func (l *List) Dump() {
	fmt.Println("List.Dump() {")
	fmt.Println("\tHead =", l.Head)
	fmt.Println("\tTail =", l.Tail)
	for current := l.Head; current != nil; current = current.Next {
		fmt.Println("\t>>>", current.Value)
	} 
	fmt.Println("}")
}

func NewList() *List {
	return &List{}
}
/*
func main() {
	l := NewList()
	l.Append(NewNode("One"))
	l.Append(NewNode("Two"))
	n3 := l.Append(NewNode("Three"))
	l.Append(NewNode("Four"))
	l.Append(NewNode("Five"))
	l.Dump()
	l.Remove(n3)
	l.RemoveByValue("One")
	l.RemoveByValue("Five")
	l.Dump()
}
*/
