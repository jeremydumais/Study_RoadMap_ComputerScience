package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList_Make_ReturnNilHead(t *testing.T) {
    var actual = MakeSingleLinkedList()
    assert.Nil(t, actual.Head())
}

func TestSinglyLinkedList_GetNode_WithIndexMinus1OnEmptyList_ReturnError(t *testing.T) {
    var actual = MakeSingleLinkedList()
    res, err := actual.GetNode(-1)
    assert.Nil(t, res)
    assert.Equal(t, "Index out of range", err.Error())
}

//TODO to continue

func TestSinglyLinkedList_AddNode_ReturnNonNilHeadAndOneNode(t *testing.T) {
    var actual = MakeSingleLinkedList()
    actual.AddNode("Test")
    assert.NotNil(t, actual.Head())
    assert.Equal(t, "Test", actual.Head().value)
    assert.Nil(t, actual.Head().next)
}

func TestSinglyLinkedList_AddTwoNodes_ReturnSuccess(t *testing.T) {
    var actual = MakeSingleLinkedList()
    actual.AddNode("Test")
    actual.AddNode("Test2")
    assert.NotNil(t, actual.Head())
    assert.NotNil(t, actual.Head().next)
    assert.Equal(t, "Test", actual.Head().value)
    assert.Equal(t, "Test2", actual.Head().next.value)
    assert.Nil(t, actual.Head().next.next)
}

func TestSinglyLinkedList_AddThreeNodes_ReturnSuccess(t *testing.T) {
    var actual = MakeSingleLinkedList()
    actual.AddNode("Test")
    actual.AddNode("Test2")
    actual.AddNode("Test3")
    assert.NotNil(t, actual.Head())
    assert.NotNil(t, actual.Head().next)
    assert.NotNil(t, actual.Head().next.next)
    assert.Equal(t, "Test", actual.Head().value)
    assert.Equal(t, "Test2", actual.Head().next.value)
    assert.Equal(t, "Test3", actual.Head().next.next.value)
    assert.Nil(t, actual.Head().next.next.next)
}

