// Package singlylinkedlist contains all the objects of a singly linked list.
package singlylinkedlist

import "errors"

// Node This represent a node item in the list.
type Node struct {
	Next  *Node
	Value string
}

// SingleLinkedList is the interface to the singly linked list structure. 
// A single linked list is a type of linked list in which each node has only 
// one link to the next node in the list.
type SingleLinkedList interface {
    // Return the first node of the list. If the list is empty nil will be returned.
    Head() *Node

    // Return the last node of the list. If the list is empty nil will be returned.
    Tail() *Node

    // Return the number of nodes in the list.
    Count() int

    // Return a node at a specific index. The first element would be the index
    // zero.
    GetNode(index int) (*Node, error)

    // Return true if the list has no element and false if the list contains at
    // least one element.
    Empty() bool

    // Add a new node at the end of the list.
    AddNode(value string)

    // Insert a new node before the element at the specified index.
    InsertNodeBefore(index int, value string) error

    // Insert a new node after the element at the specified index.
    InsertNodeAfter(index int, value string) error

    // Remove the node the is located at the specified index.
    RemoveNode(index int) error

    // Remove all the elements of the list.
    Clear()

    // Remove the first element of the list.
    PopFront() error

    // Remove consecutive duplicate elements of the list.
    Unique()

    //TODO Sort 
    //TODO Reverse
}

// Singly linked list: a type of linked list in which each node has only one
// link to the next node in the list.
type singleLinkedList struct {
	head *Node
    nodeCount int
}

// MakeSingleLinkedList is the constructor of the SingleLinkedList struct.
func MakeSingleLinkedList() SingleLinkedList {
    retval := &singleLinkedList{ nil, 0 }
    return retval 
}

func (list *singleLinkedList) Head() *Node {
    return list.head
}

func (list *singleLinkedList) Tail() *Node {
    nodeAtTail, err := list.fetchNode(list.nodeCount-1)
    if err != nil {
        return nil
    }
    return nodeAtTail
}

func (list *singleLinkedList) Count() int {
    return list.nodeCount
}

func (list *singleLinkedList) GetNode(index int) (*Node, error) {
    return list.fetchNode(index);
}

func (list *singleLinkedList) Empty() bool {
    return list.nodeCount == 0 && list.head == nil
}

func (list *singleLinkedList) AddNode(value string) {
    newNode := &Node{nil, value}
    if list.head == nil {
        list.head = newNode
    } else {
        // Go to last node
        lastNode := list.head
        for (lastNode.Next != nil) {
            lastNode = lastNode.Next
        }
        lastNode.Next = newNode
    }
    list.nodeCount++
}

func (list *singleLinkedList) InsertNodeBefore(index int, value string) error {
    if index < 0 || index > list.nodeCount {
        return errors.New("index out of range")
    }
    var nodeAtInsertPosition *Node
    if index < list.nodeCount {
        nodeAtInsertPosition, _ = list.fetchNode(index)
    }
    nodeBeforeInsertPosition, _ := list.fetchNode(index - 1)
    newNode := &Node{nil, value}
    list.bindNewNode(nodeBeforeInsertPosition, newNode, nodeAtInsertPosition)
    return nil
}
    
func (list *singleLinkedList) InsertNodeAfter(index int, value string) error {
    nodeAtInsertPosition, err := list.fetchNode(index)
    if err != nil {
        return err
    }
    newNode := &Node{nil, value}
    list.bindNewNode(nodeAtInsertPosition, newNode, nodeAtInsertPosition.Next)
    return nil 
}

func (list *singleLinkedList) RemoveNode(index int) error {
    nodeAtIndex , err := list.fetchNode(index) 
    nodeBeforeIndex, _ := list.fetchNode(index - 1)
    if err != nil {
        return err
    }
    list.unbindNode(nodeBeforeIndex, nodeAtIndex)
    return nil
}

func (list *singleLinkedList) Clear() {
    list.head = nil
    list.nodeCount = 0
}

func (list *singleLinkedList) PopFront() error {
    if list.Empty() {
        return errors.New("the list is empty")
    }
    list.unbindNode(nil, list.head)
    return nil
}

func (list *singleLinkedList) Unique() {
    nodeBefore := list.head
    for i := 1; i < list.nodeCount; i++ {
        currentNode, _ := list.fetchNode(i)
        if currentNode.Value == nodeBefore.Value {
            list.unbindNode(nodeBefore, currentNode)
            i--
        } else {
            nodeBefore = currentNode
        }
    }
}

func (list *singleLinkedList) bindNewNode(nodeBefore *Node, newNode *Node, nodeAfter *Node) {
    if nodeBefore != nil {
        nodeBefore.Next = newNode
    } else {
        list.head = newNode
    }
    if (nodeAfter != nil) {
        newNode.Next = nodeAfter
    } else {
        newNode.Next = nil
    }
    list.nodeCount++
}

func (list *singleLinkedList) unbindNode(nodeBefore *Node, node *Node) {
    if nodeBefore == nil {
        list.head = node.Next
    } else {
        nodeBefore.Next = node.Next
    }
    list.nodeCount--

}

func (list *singleLinkedList) fetchNode(index int) (*Node, error) {
    if index < 0 || index > list.nodeCount - 1 {
        return nil, errors.New("index out of range") 
    }
    currentNode := list.head
    for i := 0; i < index; i++ {
        currentNode = currentNode.Next
    }
    return currentNode, nil
}
