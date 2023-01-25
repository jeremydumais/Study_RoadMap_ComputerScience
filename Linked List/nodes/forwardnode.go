// Package nodes contains all the objects for the different node types.
package nodes

// ForwardNode is the interface for the node structure of a single linked list
// or a circular linked list. This represent a node item in the list and has
// only a next method for navigate.
type ForwardNode interface {
    // Return the next node. If there is no next node the value returned
    // is nil.
    Next() ForwardNode

    // Return the value of the node as a string.
    Value() string

    // Change the value of the node with the value passed in argument.
    SetValue(value string)
    
    // Change the next node link.
    SetNext(node ForwardNode)
}

type forwardNode struct {
    next *forwardNode
    value string
}

// MakeForwardNode is the constructor of the forwardNode struct.
func MakeForwardNode(nodeAfter ForwardNode, 
                     value string) ForwardNode {
    var after *forwardNode
    if nodeAfter != nil {
        after = nodeAfter.(*forwardNode)
    }
    retval := &forwardNode{ after,
                            value }
    return retval 
}

func (nodeitem *forwardNode) Next() ForwardNode {
    return nodeitem.next
}

func (nodeitem *forwardNode) Value() string {
    return nodeitem.value
}

func (nodeitem *forwardNode) SetValue(value string) {
    nodeitem.value = value
}
    
func (nodeitem *forwardNode) SetNext(node ForwardNode) {
    if node != nil {
        nodeitem.next = node.(*forwardNode)
    } else {
        nodeitem.next = nil
    }
}

