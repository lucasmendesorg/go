/*
	Hash Table supporting collisions using linked list
		Written by Lucas Mendes <lucas AT lucasmendes.org>
*/

package main

import (
	"errors"
	"fmt"
)

const HashTableArraySize = 100

//
// hasnTableNode
//

type hashTableNode struct {
	next  *hashTableNode
	key   string
	value int
}

func (htn *hashTableNode) FindByKey(key string) *hashTableNode {
	for current := htn; current != nil; current = current.next {
		if current.key == key {
			return current
		}
	}
	return nil
}

func (htn *hashTableNode) Insert(node hashTableNode) {
	if htn.next == nil {
		htn.next = &node
		return
	}
	current := htn.next
	for current.next != nil {
		current = current.next
	}
	current.next = &node
}

func NewHashTableNode(key string, value int) hashTableNode {
	return hashTableNode{
		next:  nil,
		key:   key,
		value: value,
	}
}

//
// Hash Table
//

type HashTable struct {
	array [HashTableArraySize]*hashTableNode
}

func (ht *HashTable) hash(key string) int {
	sum := 0
	for c := range key {
		sum += c
	}
	return (sum % HashTableArraySize)
}

func (ht *HashTable) Set(key string, value int) {
	hash := ht.hash(key)
	fmt.Println("HashTable.Set: Hash for", key, "is", hash)
	htn := NewHashTableNode(key, value)
	if ht.array[hash] == nil {
		ht.array[hash] = &htn
	} else {
		found := ht.array[hash].FindByKey(key)
		if found == nil {
			ht.array[hash].Insert(htn)
		} else {
			found.value = value
//			ht.array[hash].Insert(*found)
		}
	}
}

func (ht *HashTable) Get(key string) (int, error) {
	hash := ht.hash(key)
	current := ht.array[hash]
	for ; current != nil; current = current.next {
		fmt.Println("HashTable.Get: Trying for key =", key)
		if current.key == key {
			return current.value, nil
		}
	}
	return -1, errors.New("value not found for key")
}

func NewHashTable() HashTable {
	return HashTable{}
}

func main() {
	ht := NewHashTable()
	
	ht.Set("eric", 111)			// ht["eric"] = 111;
	ht.Set("erhd", 222)			// ht["erhd"] = 222;
	ht.Set("luciano", 333)
	ht.Set("lucas", 444)
	
	eric, err := ht.Get("eric")	// eric = ht["eric"];
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("eric =", eric)
	
	erhd, err := ht.Get("erhd")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("erhd =", erhd)
	
	luciano, err := ht.Get("luciano")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("luciano =", luciano)
	
	lucas, err := ht.Get("lucas")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("lucas =", lucas)
	
}
