package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList_Make_ReturnNilHead(t *testing.T) {
    var actual = MakeSingleLinkedList()
    assert.Nil(t, actual.Head())
}

func TestSinglyLinkedList_AddNode_ReturnNonNilHeadAndOneNode(t *testing.T) {
    var actual = MakeSingleLinkedList()
    actual.AddItem("Test")
    assert.NotNil(t, actual.Head())
    assert.Equal(t, "Test", actual.Head().value)
    assert.Nil(t, actual.Head().next)
}

