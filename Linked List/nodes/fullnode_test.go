package nodes

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFullNode_MakeFullNode_WithNilNilAndTest_ReturnSuccess(t *testing.T) {
    nodeItem := MakeFullNode(nil, nil, "Test")
    assert.Nil(t, nodeItem.Prev())
    assert.Nil(t, nodeItem.Next())
    assert.Equal(t, "Test", nodeItem.Value())
}

func TestFullNode_MakeFullNode_WithNode1Node2AndTest_ReturnSuccess(t *testing.T) {
    nodeItem1 := MakeFullNode(nil, nil, "Test1")
    nodeItem2 := MakeFullNode(nil, nil, "Test2")
    nodeItem := MakeFullNode(nodeItem1, nodeItem2, "Test")
    assert.Same(t, nodeItem1, nodeItem.Prev())
    assert.Same(t, nodeItem2, nodeItem.Next())
    assert.Equal(t, "Test", nodeItem.Value())
}

func TestFullNode_Prev_WithUnattachedNode_ReturnNil(t *testing.T) {
    nodeItem := FullNode(&fullNode { nil, nil, "Test" })
    assert.Nil(t, nodeItem.Prev())
}

func TestFullNode_Prev_WithAttachedNode_ReturnPreviousNode(t *testing.T) {
    node1 := fullNode { nil, nil, "Test1"}
    node2 := fullNode { &node1, nil, "Test2" }
    node1.next = &node2
    nodeItem := FullNode(&node2)
    assert.NotNil(t, nodeItem.Prev())
    assert.Equal(t, node1.Value(), nodeItem.Prev().Value())
}

func TestFullNode_Next_WithUnattachedNode_ReturnNil(t *testing.T) {
    nodeItem := FullNode(&fullNode { nil, nil, "Test" })
    assert.Nil(t, nodeItem.Next())
}

func TestFullNode_Next_WithAttachedNode_ReturnNextNode(t *testing.T) {
    node1 := fullNode { nil, nil, "Test1"}
    node2 := fullNode { &node1, nil, "Test2" }
    node1.next = &node2
    nodeItem := FullNode(&node1)
    assert.NotNil(t, nodeItem.Next())
    assert.Equal(t, node2.Value(), nodeItem.Next().Value())
}

func TestFullNode_Value_WithUnattachedNode_ReturnTest(t *testing.T) {
    nodeItem := FullNode(&fullNode { nil, nil, "Test" })
    assert.Equal(t, "Test", nodeItem.Value())
}

func TestFullNode_SetValue_WithEmpty_ReturnSuccess(t *testing.T) {
    nodeItem := FullNode(&fullNode { nil, nil, "Test" })
    nodeItem.SetValue("")
    assert.Equal(t, "", nodeItem.Value())
}

func TestFullNode_SetValue_WithHi_ReturnSuccess(t *testing.T) {
    nodeItem := FullNode(&fullNode { nil, nil, "Test" })
    nodeItem.SetValue("Hi")
    assert.Equal(t, "Hi", nodeItem.Value())
}

func TestFullNode_SetPrev_WithNode1_ReturnSuccess(t *testing.T) {
    nodeItem := MakeFullNode(nil, nil, "Test")
    nodePrev := MakeFullNode(nil, nil, "Test")
    nodeItem.SetPrev(nodePrev)
    assert.Same(t, nodePrev, nodeItem.Prev())
}

func TestFullNode_SetPrev_WithNil_ReturnSuccess(t *testing.T) {
    nodePrev := MakeFullNode(nil, nil, "Test")
    nodeItem := MakeFullNode(nodePrev, nil, "Test")
    nodeItem.SetPrev(nil)
    assert.Nil(t, nodeItem.Prev())
}

func TestFullNode_SetNext_WithNode1_ReturnSuccess(t *testing.T) {
    nodeItem := MakeFullNode(nil, nil, "Test")
    nodeNext := MakeFullNode(nil, nil, "Test")
    nodeItem.SetNext(nodeNext)
    assert.Same(t, nodeNext, nodeItem.Next())
}

func TestFullNode_SetNext_WithNil_ReturnSuccess(t *testing.T) {
    nodeNext := MakeFullNode(nil, nil, "Test")
    nodeItem := MakeFullNode(nil, nodeNext, "Test")
    nodeItem.SetNext(nil)
    assert.Nil(t, nodeItem.Next())
}

