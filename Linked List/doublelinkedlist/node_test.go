package doublelinkedlist

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestDoubleLinkedListNode_Prev_WithUnattachedNode_ReturnNil(t *testing.T) {
    nodeItem := Node(&node { nil, nil, "Test" })
    assert.Nil(t, nodeItem.Prev())
}

func TestDoubleLinkedListNode_Prev_WithAttachedNode_ReturnPreviousNode(t *testing.T) {
    node1 := node { nil, nil, "Test1"}
    node2 := node { &node1, nil, "Test2" }
    node1.next = &node2
    nodeItem := Node(&node2)
    assert.NotNil(t, nodeItem.Prev())
    assert.Equal(t, node1.Value(), nodeItem.Prev().Value())
}

func TestDoubleLinkedListNode_Next_WithUnattachedNode_ReturnNil(t *testing.T) {
    nodeItem := Node(&node { nil, nil, "Test" })
    assert.Nil(t, nodeItem.Next())
}

func TestDoubleLinkedListNode_Next_WithAttachedNode_ReturnPreviousNode(t *testing.T) {
    node1 := node { nil, nil, "Test1"}
    node2 := node { &node1, nil, "Test2" }
    node1.next = &node2
    nodeItem := Node(&node1)
    assert.NotNil(t, nodeItem.Next())
    assert.Equal(t, node2.Value(), nodeItem.Next().Value())
}

func TestDoubleLinkedListNode_Value_WithUnattachedNode_ReturnTest(t *testing.T) {
    nodeItem := Node(&node { nil, nil, "Test" })
    assert.Equal(t, "Test", nodeItem.Value())
}

func TestDoubleLinkedListNode_SetValue_WithEmpty_ReturnSuccess(t *testing.T) {
    nodeItem := Node(&node { nil, nil, "Test" })
    nodeItem.SetValue("")
    assert.Equal(t, "", nodeItem.Value())
}

func TestDoubleLinkedListNode_SetValue_WithHi_ReturnSuccess(t *testing.T) {
    nodeItem := Node(&node { nil, nil, "Test" })
    nodeItem.SetValue("Hi")
    assert.Equal(t, "Hi", nodeItem.Value())
}

