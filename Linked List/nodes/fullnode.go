// Package nodes contains all the objects for the different node types.
package nodes

// FullNode is the interface for the node structure of a double linked list
// or a circular double linked list. This represent a node item in the list and 
// has a next and a previous method for navigate.
type FullNode interface {
    // Return the previous node. If there is no previous node the value returned
    // is nil.
    Prev() FullNode    
    
    // Return the next node. If there is no next node the value returned
    // is nil.
    Next() FullNode

    // Return the value of the node as a string.
    Value() string

    // Change the value of the node with the value passed in argument.
    SetValue(value string)
    
    // Change the previous node link.
    SetPrev(node FullNode)

    // Change the next node link.
    SetNext(node FullNode)
}

type fullNode struct {
    prev *fullNode
    next *fullNode
    value string
}

// MakeFullNode is the constructor of the fullNode struct.
func MakeFullNode(nodeBefore FullNode,
                  nodeAfter FullNode, 
                  value string) FullNode {
    var before *fullNode
    if nodeBefore != nil {
        before = nodeBefore.(*fullNode)
    }
    var after *fullNode
    if nodeAfter != nil {
        after = nodeAfter.(*fullNode)
    }
    retval := &fullNode{ before, 
                         after,
                         value }
    return retval 
}

func (nodeitem *fullNode) Prev() FullNode {
    return nodeitem.prev
}

func (nodeitem *fullNode) Next() FullNode {
    return nodeitem.next
}

func (nodeitem *fullNode) Value() string {
    return nodeitem.value
}

func (nodeitem *fullNode) SetValue(value string) {
    nodeitem.value = value
}
    
func (nodeitem *fullNode) SetPrev(node FullNode) {
    if node != nil {
        nodeitem.prev = node.(*fullNode)
    } else {
        nodeitem.prev = nil
    }
}

func (nodeitem *fullNode) SetNext(node FullNode) {
    if node != nil {
        nodeitem.next = node.(*fullNode)
    } else {
        nodeitem.next = nil
    }
}

