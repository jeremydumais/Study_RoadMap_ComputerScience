package main

import "errors"

// Node This represent a.
type Node struct {
	next  *Node
	value string
}

type SingleLinkedList interface {
    Head() *Node
    GetNode(index int) (*Node, error)
    AddNode(value string)
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

func (list *singleLinkedList) GetNode(index int) (*Node, error) {
    if index < 0 {
        return nil, errors.New("Index out of range") 
    }
    return nil, nil
}

func (list *singleLinkedList) AddNode(value string) {
    if list.head == nil {
        newNode := &Node{nil, value}
        list.head = newNode
    } else {
        // Go to last node
        var lastNode *Node = list.head
        for (lastNode.next != nil) {
            lastNode = lastNode.next
        }
        newNode := &Node{nil, value}
        lastNode.next = newNode
    }
}
