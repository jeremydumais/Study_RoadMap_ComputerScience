package main

// Node This represent a.
type Node struct {
	next  *Node
	value string
}

type SingleLinkedList interface {
    GetHead() *Node
    AddItem(value string)
}

// Singly linked list: a type of linked list in which each node has only one 
// link to the next node in the list.
type singleLinkedList struct {
	head *Node
}

func MakeSingleLinkedList() SingleLinkedList {
    return singleLinkedList { nil }
}

func (list singleLinkedList) GetHead() *Node {
    return list.head
}

func (list singleLinkedList) AddItem(value string) {
    
}
