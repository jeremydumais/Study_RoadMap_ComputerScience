// Package singlelinkedlist contains all the objects of a singly linked list.
package singlelinkedlist

import (
    "errors"
	"study/linkedlistdemo/nodes"
    "study/linkedlistdemo/helpers"
)

// SingleLinkedList is the interface to the single linked list structure. 
// A single linked list is a type of linked list in which each node has only 
// one link to the next node in the list.
type SingleLinkedList interface {
    // Return the first node of the list. If the list is empty nil will be returned.
    Head() nodes.ForwardNode

    // Return the last node of the list. If the list is empty nil will be returned.
    Tail() nodes.ForwardNode

    // Return the number of nodes in the list.
    Len() int

    // Return a node at a specific index. The first element would be the index
    // zero.
    GetNode(index int) (nodes.ForwardNode, error)

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

    // Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
    Less(i, j int) bool

    // Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Single linked list: a type of linked list in which each node has only one
// link to the next node in the list.
type singleLinkedList struct {
	head nodes.ForwardNode
    nodeCount int
}

// MakeSingleLinkedList is the constructor of the SingleLinkedList struct.
func MakeSingleLinkedList() SingleLinkedList {
    retval := &singleLinkedList{ nil, 0 }
    return retval 
}

func (list *singleLinkedList) Head() nodes.ForwardNode {
    return list.head
}

func (list *singleLinkedList) Tail() nodes.ForwardNode {
    nodeAtTail, err := list.fetchNode(list.nodeCount-1)
    if err != nil {
        return nil
    }
    return nodeAtTail
}

func (list *singleLinkedList) Len() int {
    return list.nodeCount
}

func (list *singleLinkedList) GetNode(index int) (nodes.ForwardNode, error) {
    return list.fetchNode(index);
}

func (list *singleLinkedList) Empty() bool {
    return list.nodeCount == 0 && 
    helpers.IsInterfaceNil(list.head)
}

func (list *singleLinkedList) AddNode(value string) {
    newNode := nodes.MakeForwardNode(nil, value)
    if helpers.IsInterfaceNil(list.head) {
        list.head = newNode
    } else {
        // Go to last node
        lastNode := list.head
        for (!helpers.IsInterfaceNil(lastNode.Next())) {
            lastNode = lastNode.Next()
        }
        lastNode.SetNext(newNode)
    }
    list.nodeCount++
}

func (list *singleLinkedList) InsertNodeBefore(index int, value string) error {
    if index < 0 || index > list.nodeCount {
        return errors.New("index out of range")
    }
    var nodeAtInsertPosition nodes.ForwardNode
    if index < list.nodeCount {
        nodeAtInsertPosition, _ = list.fetchNode(index)
    }
    nodeBeforeInsertPosition, _ := list.fetchNode(index - 1)
    newNode := nodes.MakeForwardNode(nil, value)
    list.bindNewNode(nodeBeforeInsertPosition, newNode, nodeAtInsertPosition)
    return nil
}
    
func (list *singleLinkedList) InsertNodeAfter(index int, value string) error {
    nodeAtInsertPosition, err := list.fetchNode(index)
    if err != nil {
        return err
    }
    newNode := nodes.MakeForwardNode(nil, value)
    list.bindNewNode(nodeAtInsertPosition, newNode, nodeAtInsertPosition.Next())
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
        if currentNode.Value() == nodeBefore.Value() {
            list.unbindNode(nodeBefore, currentNode)
            i--
        } else {
            nodeBefore = currentNode
        }
    }
}

func (list *singleLinkedList) Less(i, j int) bool {
    nodeI, err := list.fetchNode(i)
    if err != nil {
        return false
    }
    nodeJ, err := list.fetchNode(j)
    if err != nil {
        return false
    }
    return nodeI.Value() < nodeJ.Value()
}

func (list *singleLinkedList) Swap(i, j int) {
    nodeI, err := list.fetchNode(i)
    if err != nil {
        return
    }
    nodeJ, err := list.fetchNode(j)
    if err != nil {
        return
    }
    temp := nodeJ.Value()
    nodeJ.SetValue(nodeI.Value())
    nodeI.SetValue(temp)
}

func (list *singleLinkedList) bindNewNode(nodeBefore nodes.ForwardNode, newNode nodes.ForwardNode, nodeAfter nodes.ForwardNode) {
    if !helpers.IsInterfaceNil(nodeBefore) {
        nodeBefore.SetNext(newNode)
    } else {
        list.head = newNode
    }
    if !helpers.IsInterfaceNil(nodeAfter) {
        newNode.SetNext(nodeAfter)
    } else {
        newNode.SetNext(nil)
    }
    list.nodeCount++
}

func (list *singleLinkedList) unbindNode(nodeBefore nodes.ForwardNode, node nodes.ForwardNode) {
    if helpers.IsInterfaceNil(nodeBefore) {
        list.head = node.Next()
    } else {
        nodeBefore.SetNext(node.Next())
    }
    list.nodeCount--
}

func (list *singleLinkedList) fetchNode(index int) (nodes.ForwardNode, error) {
    if index < 0 || index > list.nodeCount - 1 {
        return nil, errors.New("index out of range") 
    }
    currentNode := list.head
    for i := 0; i < index; i++ {
        currentNode = currentNode.Next()
    }
    return currentNode, nil
}
