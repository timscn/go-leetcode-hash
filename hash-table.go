package main

import "fmt"

const ArraySize = 7

// hastable structure - will hold array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket structure - will be linkedList
type bucket struct {
	head *bucketNode
}

// bucket Node structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// init will create a bucket into each arrayElement

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// the below three for array operations so Keys
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// the below for Bucket which is linkedList

// insert for bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Printf("The %s alreay exist", k)
	}
}

// delete  for bucket
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previosNode := b.head
	for previosNode.next != nil {
		if previosNode.next.key == k {
			// delete
			previosNode.next = previosNode.next.next
		}
		previosNode = previosNode.next
	}
}

// seach  for bucket
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func hash(key string) int {
	sum := 0
	for _, elementOfKey := range key {
		sum += int(elementOfKey)
	}
	return sum % ArraySize
}
