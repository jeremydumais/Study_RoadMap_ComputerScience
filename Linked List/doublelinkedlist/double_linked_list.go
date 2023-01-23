// Package doublelinkedlist contains all the objects of a doubly linked list.
package doublelinkedlist

import (
    "errors"
	"study/linkedlistdemo/nodes"
	"study/linkedlistdemo/helpers"
)

// DoubleLinkedList is the interface to the double linked list structure.
// A double linked list is a type of linked list in which each node has a
// previous and a next link to the corresponding node in the list.
type DoubleLinkedList interface {
    // Return the first node of the list. If the list is empty nil will be returned.
    Head() nodes.FullNode

    // Return the last node of the list. If the list is empty nil will be returned.
    Tail() nodes.FullNode

    // Return the number of nodes in the list.
    Len() int

    // Return a node at a specific index. The first element would be the index
    // zero.
    GetNode(index int) (nodes.FullNode, error)

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

    // Remove the last element of the list.
    PopBack() error

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

type doublelinkedlist struct {
	head nodes.FullNode
    tail nodes.FullNode
    nodeCount int
}

// MakeDoubleLinkedList is the constructor of the DoubleLinkedList struct.
func MakeDoubleLinkedList() DoubleLinkedList {
    retval := &doublelinkedlist{ nil, nil, 0 }
    return retval 
}

func (list *doublelinkedlist) Head() nodes.FullNode {
    return list.head
}

func (list *doublelinkedlist) Tail() nodes.FullNode {
    return list.tail
}

func (list *doublelinkedlist) Len() int {
    return list.nodeCount
}
    
func (list *doublelinkedlist) GetNode(index int) (nodes.FullNode, error) {
    return list.fetchNode(index);
}

func (list *doublelinkedlist) Empty() bool {
    return list.nodeCount == 0 && 
           helpers.IsNil(list.head) &&
           helpers.IsNil(list.tail)
}

func (list *doublelinkedlist) AddNode(value string) {
    newNode := nodes.MakeFullNode(nil, nil, value)
    if helpers.IsNil(list.head) {
        list.head = newNode
        list.tail = list.head
    } else {
        lastNode := list.tail
        lastNode.SetNext(newNode)
        newNode.SetPrev(lastNode)
        list.tail = newNode
    }
    list.nodeCount++
}

func (list *doublelinkedlist) InsertNodeBefore(index int, value string) error {
    if index < 0 || index > list.nodeCount {
        return errors.New("index out of range")
    }
    var nodeAtInsertPosition nodes.FullNode
    if index < list.nodeCount {
        nodeAtInsertPosition, _ = list.fetchNode(index)
    }
    nodeBeforeInsertPosition, _ := list.fetchNode(index - 1)
    newNode := nodes.MakeFullNode(nodeBeforeInsertPosition, 
                                  nodeAtInsertPosition, 
                                  value)
    list.bindNewNode(nodeBeforeInsertPosition, newNode, nodeAtInsertPosition)
    return nil
}
    
func (list *doublelinkedlist) InsertNodeAfter(index int, value string) error {
    nodeAtInsertPosition, err := list.fetchNode(index)
    if err != nil {
        return err
    }
    newNode := nodes.MakeFullNode(nil, nil, value)
    list.bindNewNode(nodeAtInsertPosition, newNode, nodeAtInsertPosition.Next())
    return nil 

}
    
func (list *doublelinkedlist) RemoveNode(index int) error {
    nodeAtIndex , err := list.fetchNode(index) 
    nodeBeforeIndex, _ := list.fetchNode(index - 1)
    nodeAfterIndex, _ := list.fetchNode(index + 1) 
    if err != nil {
        return err
    }
    list.unbindNode(nodeBeforeIndex, nodeAtIndex, nodeAfterIndex)
    return nil
}
    
func (list *doublelinkedlist) Clear() {
    list.head = nil
    list.tail = nil
    list.nodeCount = 0
}

func (list *doublelinkedlist) PopFront() error {
    if list.Empty() {
        return errors.New("the list is empty")
    }
    nodeAfterIndex, _ := list.fetchNode(1) 
    list.unbindNode(nil, list.head, nodeAfterIndex)
    return nil
}

func (list *doublelinkedlist) PopBack() error {
    if list.Empty() {
        return errors.New("the list is empty")
    }
    nodeBeforeIndex, _ := list.fetchNode(list.Len() - 2) 
    list.unbindNode(nodeBeforeIndex, list.tail, nil)
    return nil
}

func (list *doublelinkedlist) Unique() {
    nodeBefore := list.head
    for i := 1; i < list.nodeCount; i++ {
        currentNode, _ := list.fetchNode(i)
        nodeAfter, _ := list.fetchNode(i + 1)
        if currentNode.Value() == nodeBefore.Value() {
            list.unbindNode(nodeBefore, currentNode, nodeAfter)
            i--
        } else {
            nodeBefore = currentNode
        }
    }
}

func (list *doublelinkedlist) Less(i, j int) bool {
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

func (list *doublelinkedlist) Swap(i, j int) {
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

func (list *doublelinkedlist) bindNewNode(nodeBefore nodes.FullNode, 
                                          newNode nodes.FullNode, 
                                          nodeAfter nodes.FullNode) {
    if !helpers.IsNil(nodeBefore) {
        nodeBefore.SetNext(newNode)
        newNode.SetPrev(nodeBefore)
    } else {
        list.head = newNode
    }
    if !helpers.IsNil(nodeAfter) {
        newNode.SetNext(nodeAfter)
        nodeAfter.SetPrev(newNode)
    } else {
        list.tail = newNode
    }
    list.nodeCount++
}

func (list *doublelinkedlist) unbindNode(nodeBefore nodes.FullNode, 
                                         node nodes.FullNode, 
                                         nodeAfter nodes.FullNode) {
    if helpers.IsNil(nodeBefore) {
        list.head = node.Next()
    } else {
        nodeBefore.SetNext(node.Next())
    }
    if helpers.IsNil(nodeAfter) {
        list.tail = node.Prev()
    } else {
        nodeAfter.SetPrev(node.Prev())
    }
    list.nodeCount--
}

func (list *doublelinkedlist) fetchNode(index int) (nodes.FullNode, error) {
    if index < 0 || index > list.nodeCount - 1 {
        return nil, errors.New("index out of range") 
    }
    currentNode := list.head
    for i := 0; i < index; i++ {
        currentNode = currentNode.Next()
    }
    return currentNode, nil
}
