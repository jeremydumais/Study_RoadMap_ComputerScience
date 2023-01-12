package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
    "fmt"
)

func getListOfElements(numberOfNodes int) SingleLinkedList {
    list := MakeSingleLinkedList() 
    for i := 0; i < numberOfNodes; i++ {
        nodeValue := fmt.Sprintf("Test%d", i+1)
        list.AddNode(nodeValue) 
    }
    return list
}

func TestSinglyLinkedList_Make_ReturnNilHead(t *testing.T) {
    actual := MakeSingleLinkedList()
    assert.Nil(t, actual.Head())
}

func TestSinglyLinkedList_GetNode_WithIndexMinus1EmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    res, err := actual.GetNode(-1)
    assert.Nil(t, res)
    assert.Equal(t, "Index out of range", err.Error())
}

func TestSinglyLinkedList_GetNode_WithIndex0EmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    res, err := actual.GetNode(0)
    assert.Nil(t, res)
    assert.Equal(t, "Index out of range", err.Error())
}

func TestSinglyLinkedList_GetNode_WithIndex0OneElementList_ReturnNode(t *testing.T) {
    actual := getListOfElements(1) 
    res, _ := actual.GetNode(0)
    assert.NotNil(t, res)
    assert.Equal(t, "Test1", res.value)
}

func TestSinglyLinkedList_GetNode_WithIndex1TwoElementsList_ReturnNode(t *testing.T) {
    actual := getListOfElements(2)
    res, _ := actual.GetNode(1)
    assert.NotNil(t, res)
    assert.Equal(t, "Test2", res.value)
}

func TestSinglyLinkedList_GetNode_WithIndex2TwoElementsList_ReturnError(t *testing.T) {
    actual := getListOfElements(2)
    res, err := actual.GetNode(2)
    assert.Nil(t, res)
    assert.Equal(t, "Index out of range", err.Error())
}

func TestSinglyLinkedList_Count_With2TwoElementsList_Return2(t *testing.T) {
    assert.Equal(t, 2, getListOfElements(2).Count())
}

func TestSinglyLinkedList_AddNode_ReturnNonNilHeadAndOneNode(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test")
    assert.Equal(t, 1, actual.Count())
    assert.NotNil(t, actual.Head())
    assert.Equal(t, "Test", actual.Head().value)
    assert.Nil(t, actual.Head().next)
}

func TestSinglyLinkedList_AddTwoNodes_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test")
    actual.AddNode("Test2")
    assert.Equal(t, 2, actual.Count())
    assert.NotNil(t, actual.Head())
    assert.NotNil(t, actual.Head().next)
    assert.Equal(t, "Test", actual.Head().value)
    assert.Equal(t, "Test2", actual.Head().next.value)
    assert.Nil(t, actual.Head().next.next)
}

func TestSinglyLinkedList_AddThreeNodes_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test")
    actual.AddNode("Test2")
    actual.AddNode("Test3")
    assert.Equal(t, 3, actual.Count())
    assert.NotNil(t, actual.Head())
    assert.NotNil(t, actual.Head().next)
    assert.NotNil(t, actual.Head().next.next)
    assert.Equal(t, "Test", actual.Head().value)
    assert.Equal(t, "Test2", actual.Head().next.value)
    assert.Equal(t, "Test3", actual.Head().next.next.value)
    assert.Nil(t, actual.Head().next.next.next)
}

func TestSinglyLinkedList_InsertNodeAtIndexZeroEmptyList_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNode(0, "Test")
    assert.Nil(t, err)
}


func TestSinglyLinkedList_InsertNodeAtIndexOneEmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNode(1, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "Index out of range", err.Error())
}

func TestSinglyLinkedList_InsertNodeAtIndexZeroOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNode(0, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Another", actual.Head().value)
    assert.Equal(t, "Test1", actual.Head().next.value)
}

func TestSinglyLinkedList_InsertNodeAtIndexOneOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNode(1, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Test1", actual.Head().value)
    assert.Equal(t, "Another", actual.Head().next.value)
}

func TestSinglyLinkedList_InsertNodeAtIndexTwoOneList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNode(2, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "Index out of range", err.Error())
}

func TestSinglyLinkedList_InsertNodeAtIndexOneTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    err := actual.InsertNode(1, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 3, actual.Count())
    assert.Equal(t, "Test1", actual.Head().value)
    assert.Equal(t, "Another", actual.Head().next.value)
    assert.Equal(t, "Test2", actual.Head().next.next.value)
}

func TestSinglyLinkedList_ClearEmptyList_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.Clear()
    assert.Equal(t, 0, actual.Count())
    assert.Nil(t, actual.Head())
}

func TestSinglyLinkedList_ClearTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    actual.Clear()
    assert.Equal(t, 0, actual.Count())
    assert.Nil(t, actual.Head())
}

func TestSinglyLinkedList_ClearThreeItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(0)
    actual.Clear()
    assert.Equal(t, 0, actual.Count())
    assert.Nil(t, actual.Head())
}

