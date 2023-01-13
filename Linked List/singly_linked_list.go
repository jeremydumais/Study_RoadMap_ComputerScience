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
    InsertNode(index int, value string) error
    Clear()
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

func (list *singleLinkedList) InsertNode(index int, value string) error {
    if index < 0 || index > list.nodeCount {
        return errors.New("Index out of range")
    }
    var nodeAtInsertPosition *Node = nil
    if index < list.nodeCount {
        nodeAtInsertPosition, _ = list.fetchNode(index)
    }
    nodeBeforeInsertPosition, _ := list.fetchNode(index - 1)
    if nodeAtInsertPosition == nil {
        newNode := &Node{nodeAtInsertPosition, value}
        if index == 0 {
            list.head = newNode
        } else {
            nodeBeforeInsertPosition.next = newNode
        }
        list.nodeCount++
        return nil
    } else {
        newNode := &Node{nodeAtInsertPosition, value}
        if nodeBeforeInsertPosition != nil {
            nodeBeforeInsertPosition.next = newNode
        } else {
            list.head = newNode
        }
        list.nodeCount++
        return nil
    }
}

func (list *singleLinkedList) Clear() {
    list.head = nil
    list.nodeCount = 0
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
