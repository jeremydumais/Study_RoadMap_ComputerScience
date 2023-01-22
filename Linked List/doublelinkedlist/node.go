// Package doublelinkedlist contains all the objects of a doubly linked list.
package doublelinkedlist

// Node is the interface of the node structure. This represent a node item in
// the list.
type Node interface {
    // Return the previous node. If there is no previous node the value returned
    // is nil.
    Prev() Node    
    
    // Return the next node. If there is no next node the value returned
    // is nil.
    Next() Node

    // Return the value of the node as a string.
    Value() string

    // Change the value of the node with the value passed in argument.
    SetValue(value string)
}

type node struct {
    prev *node
    next *node
    value string
}

func (nodeitem *node) Prev() Node {
    return nodeitem.prev
}

func (nodeitem *node) Next() Node {
    return nodeitem.next
}

func (nodeitem *node) Value() string {
    return nodeitem.value
}

func (nodeitem *node) SetValue(value string) {
    nodeitem.value = value
}
