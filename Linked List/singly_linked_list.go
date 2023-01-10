package main

import "errors"

// Node This represent a.
type Node struct {
	next  *Node
	value string
}

type SingleLinkedList interface {
    Head() *Node
    Count() int
    GetNode(index int) (*Node, error)
    AddNode(value string)
}

// Singly linked list: a type of linked list in which each node has only one 
// link to the next node in the list.
type singleLinkedList struct {
	head *Node
    nodeCount int
}

func MakeSingleLinkedList() SingleLinkedList {
    retval := &singleLinkedList{ nil, 0 }
    return retval 
}

func (list *singleLinkedList) Head() *Node {
    return list.head
}

func (list *singleLinkedList) Count() int {
    return list.nodeCount
}

func (list *singleLinkedList) GetNode(index int) (*Node, error) {
    return list.fetchNode(index);
}

func (list *singleLinkedList) AddNode(value string) {
    newNode := &Node{nil, value}
    if list.head == nil {
        list.head = newNode
    } else {
        // Go to last node
        var lastNode *Node = list.head
        for (lastNode.next != nil) {
            lastNode = lastNode.next
        }
        lastNode.next = newNode
    }
    list.nodeCount++
}

func (list *singleLinkedList) insertNode(index int, value string) error {
    return errors.New("To implement")
}

func (list *singleLinkedList) fetchNode(index int) (*Node, error) {
    if index < 0 || index > list.nodeCount - 1 {
        return nil, errors.New("Index out of range") 
    }
    currentNode := list.head
    for i := 0; i < index; i++ {
        currentNode = currentNode.next
    }
    return currentNode, nil
}
