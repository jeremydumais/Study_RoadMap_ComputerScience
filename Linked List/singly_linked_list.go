package main

// Node This represent a.
type Node struct {
	next  *Node
	value string
}

type SingleLinkedList interface {
    Head() *Node
    AddItem(value string)
}

// Singly linked list: a type of linked list in which each node has only one 
// link to the next node in the list.
type singleLinkedList struct {
	head *Node
}

func MakeSingleLinkedList() SingleLinkedList {
    retval := &singleLinkedList{nil}
    return retval 
}

func (list *singleLinkedList) Head() *Node {
    return list.head
}

func (list *singleLinkedList) AddItem(value string) {
    if list.head == nil {
        newNode := &Node{nil, value}
        list.head = newNode
    }
}

