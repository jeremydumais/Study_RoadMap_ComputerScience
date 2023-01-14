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

func TestSinglyLinkedList_Tail_EmptyList_ReturnNil(t *testing.T) {
    assert.Nil(t, MakeSingleLinkedList().Tail())
}

func TestSinglyLinkedList_Tail_OneItemList_ReturnHead(t *testing.T) {
    list := getListOfElements(1)
    tail := list.Tail()
    assert.NotNil(t, tail)
    assert.Equal(t, list.Head(), tail)
}

func TestSinglyLinkedList_Tail_TwoItemsList_ReturnLastNode(t *testing.T) {
    list := getListOfElements(2)
    tail := list.Tail()
    assert.NotNil(t, tail)
    assert.Equal(t, list.Head().next, tail)
}

func TestSinglyLinkedList_Count_With2TwoElementsList_Return2(t *testing.T) {
    assert.Equal(t, 2, getListOfElements(2).Count())
}

func TestSinglyLinkedList_GetNode_WithIndexMinus1EmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    res, err := actual.GetNode(-1)
    assert.Nil(t, res)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_GetNode_WithIndex0EmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    res, err := actual.GetNode(0)
    assert.Nil(t, res)
    assert.Equal(t, "index out of range", err.Error())
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
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_Empty_WithEmptyList_ReturnTrue(t *testing.T) {
    assert.True(t, getListOfElements(0).Empty())
}

func TestSinglyLinkedList_Empty_WithOneItemList_ReturnTrue(t *testing.T) {
    assert.False(t, getListOfElements(1).Empty())
}

func TestSinglyLinkedList_Empty_WithTwoItemList_ReturnTrue(t *testing.T) {
    assert.False(t, getListOfElements(2).Empty())
}

func TestSinglyLinkedList_AddNode_ReturnNonNilHeadAndOneNode(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test")
    assert.Equal(t, 1, actual.Count())
    assert.NotNil(t, actual.Head())
    assert.Equal(t, "Test", actual.Head().value)
    assert.Nil(t, actual.Head().next)
}

func TestSinglyLinkedList_AddNode_WithTwoReturnSuccess(t *testing.T) {
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

func TestSinglyLinkedList_AddNode_WithThreeReturnSuccess(t *testing.T) {
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

func TestSinglyLinkedList_InsertNode_AtIndexZeroEmptyList_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNode(0, "Test")
    assert.Nil(t, err)
}


func TestSinglyLinkedList_InsertNode_AtIndexOneEmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNode(1, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_InsertNode_AtIndexZeroOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNode(0, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Another", actual.Head().value)
    assert.Equal(t, "Test1", actual.Head().next.value)
}

func TestSinglyLinkedList_InsertNode_AtIndexOneOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNode(1, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Test1", actual.Head().value)
    assert.Equal(t, "Another", actual.Head().next.value)
}

func TestSinglyLinkedList_InsertNode_AtIndexTwoOneList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNode(2, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_InsertNode_AtIndexOneTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    err := actual.InsertNode(1, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 3, actual.Count())
    assert.Equal(t, "Test1", actual.Head().value)
    assert.Equal(t, "Another", actual.Head().next.value)
    assert.Equal(t, "Test2", actual.Head().next.next.value)
}

func TestSinglyLinkedList_RemoveNode_AtIndexZeroWithEmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.RemoveNode(0)
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_RemoveNode_AtIndexZeroWithOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.RemoveNode(0)
    assert.Nil(t, err)
    assert.True(t, actual.Empty())
}

func TestSinglyLinkedList_RemoveNode_AtIndexZeroWithTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    err := actual.RemoveNode(0)
    assert.Nil(t, err)
    assert.Equal(t, 1, actual.Count())
    assert.Equal(t, "Test2", actual.Head().value)
}

func TestSinglyLinkedList_RemoveNode_AtIndexOneWithTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    err := actual.RemoveNode(1)
    assert.Nil(t, err)
    assert.Equal(t, 1, actual.Count())
    assert.Equal(t, "Test1", actual.Head().value)
}

func TestSinglyLinkedList_RemoveNode_AtIndexOneWithThreeItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(3)
    err := actual.RemoveNode(1)
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Test1", actual.Head().value)
    assert.Equal(t, "Test3", actual.Head().next.value)
}

func TestSinglyLinkedList_RemoveNode_AtIndextTwoWithThreeItemsList_ReturnSuccess(t *testing.T) {

    actual := getListOfElements(3)
    err := actual.RemoveNode(2)
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Test1", actual.Head().value)
    assert.Equal(t, "Test2", actual.Head().next.value)
}

func TestSinglyLinkedList_RemoveNode_AtIndexOneWithOneItemList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.RemoveNode(1)
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_Clear_EmptyList_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.Clear()
    assert.Equal(t, 0, actual.Count())
    assert.Nil(t, actual.Head())
}

func TestSinglyLinkedList_Clear_TwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    actual.Clear()
    assert.Equal(t, 0, actual.Count())
    assert.Nil(t, actual.Head())
}

func TestSinglyLinkedList_Clear_ThreeItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(0)
    actual.Clear()
    assert.Equal(t, 0, actual.Count())
    assert.Nil(t, actual.Head())
}

