package nodes

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestForwardNode_MakeForwardNode_WithNilAndTest_ReturnSuccess(t *testing.T) {
    nodeItem := MakeForwardNode(nil, "Test")
    assert.Nil(t, nodeItem.Next())
    assert.Equal(t, "Test", nodeItem.Value())
}

func TestForwardNode_MakeForwardNode_WithNode1Node2AndTest_ReturnSuccess(t *testing.T) {
    nodeItem1 := MakeForwardNode(nil, "Test1")
    nodeItem := MakeForwardNode(nodeItem1, "Test")
    assert.Same(t, nodeItem1, nodeItem.Next())
    assert.Equal(t, "Test", nodeItem.Value())
}

func TestForwardNode_Next_WithUnattachedNode_ReturnNil(t *testing.T) {
    nodeItem := ForwardNode(&forwardNode { nil, "Test" })
    assert.Nil(t, nodeItem.Next())
}

func TestForwardNode_Next_WithAttachedNode_ReturnNextNode(t *testing.T) {
    node1 := forwardNode { nil, "Test1"}
    node2 := forwardNode { nil, "Test2" }
    node1.next = &node2
    nodeItem := ForwardNode(&node1)
    assert.NotNil(t, nodeItem.Next())
    assert.Equal(t, node2.Value(), nodeItem.Next().Value())
}

func TestForwardNode_Value_WithUnattachedNode_ReturnTest(t *testing.T) {
    nodeItem := ForwardNode(&forwardNode { nil, "Test" })
    assert.Equal(t, "Test", nodeItem.Value())
}

func TestForwardNode_SetValue_WithEmpty_ReturnSuccess(t *testing.T) {
    nodeItem := ForwardNode(&forwardNode { nil, "Test" })
    nodeItem.SetValue("")
    assert.Equal(t, "", nodeItem.Value())
}

func TestForwardNode_SetValue_WithHi_ReturnSuccess(t *testing.T) {
    nodeItem := ForwardNode(&forwardNode { nil, "Test" })
    nodeItem.SetValue("Hi")
    assert.Equal(t, "Hi", nodeItem.Value())
}

func TestForwardNode_SetNext_WithNode1_ReturnSuccess(t *testing.T) {
    nodeItem := MakeForwardNode(nil, "Test")
    nodeNext := MakeForwardNode(nil, "Test")
    nodeItem.SetNext(nodeNext)
    assert.Same(t, nodeNext, nodeItem.Next())
}

func TestForwardNode_SetNext_WithNil_ReturnSuccess(t *testing.T) {
    nodeNext := MakeForwardNode(nil, "Test")
    nodeItem := MakeForwardNode(nodeNext, "Test")
    nodeItem.SetNext(nil)
    assert.Nil(t, nodeItem.Next())
}

