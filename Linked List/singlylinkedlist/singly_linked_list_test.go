package singlylinkedlist

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
    assert.Equal(t, list.Head().Next, tail)
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
    assert.Equal(t, "Test1", res.Value)
}

func TestSinglyLinkedList_GetNode_WithIndex1TwoElementsList_ReturnNode(t *testing.T) {
    actual := getListOfElements(2)
    res, _ := actual.GetNode(1)
    assert.NotNil(t, res)
    assert.Equal(t, "Test2", res.Value)
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
    assert.Equal(t, "Test", actual.Head().Value)
    assert.Nil(t, actual.Head().Next)
}

func TestSinglyLinkedList_AddNode_WithTwoReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test")
    actual.AddNode("Test2")
    assert.Equal(t, 2, actual.Count())
    assert.NotNil(t, actual.Head())
    assert.NotNil(t, actual.Head().Next)
    assert.Equal(t, "Test", actual.Head().Value)
    assert.Equal(t, "Test2", actual.Head().Next.Value)
    assert.Nil(t, actual.Head().Next.Next)
}

func TestSinglyLinkedList_AddNode_WithThreeReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test")
    actual.AddNode("Test2")
    actual.AddNode("Test3")
    assert.Equal(t, 3, actual.Count())
    assert.NotNil(t, actual.Head())
    assert.NotNil(t, actual.Head().Next)
    assert.NotNil(t, actual.Head().Next.Next)
    assert.Equal(t, "Test", actual.Head().Value)
    assert.Equal(t, "Test2", actual.Head().Next.Value)
    assert.Equal(t, "Test3", actual.Head().Next.Next.Value)
    assert.Nil(t, actual.Head().Next.Next.Next)
}

func TestSinglyLinkedList_InsertNodeBeforeAtIndexZeroEmptyList_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNodeBefore(0, "Test")
    assert.Nil(t, err)
}


func TestSinglyLinkedList_InsertNodeBeforeAtIndexOneEmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNodeBefore(1, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_InsertNodeBeforeAtIndexZeroOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeBefore(0, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Another", actual.Head().Value)
    assert.Equal(t, "Test1", actual.Head().Next.Value)
}

func TestSinglyLinkedList_InsertNodeBeforeAtIndexOneOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeBefore(1, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Test1", actual.Head().Value)
    assert.Equal(t, "Another", actual.Head().Next.Value)
}

func TestSinglyLinkedList_InsertNodeBeforeAtIndexTwoOneList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeBefore(2, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_InsertNodeBeforeAtIndexOneTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    err := actual.InsertNodeBefore(1, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 3, actual.Count())
    assert.Equal(t, "Test1", actual.Head().Value)
    assert.Equal(t, "Another", actual.Head().Next.Value)
    assert.Equal(t, "Test2", actual.Head().Next.Next.Value)
}

func TestSinglyLinkedList_InsertNodeAfterAtIndexZeroEmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNodeAfter(0, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_InsertNodeAfterAtIndexOneEmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNodeAfter(1, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_InsertNodeAfterAtIndexZeroOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeAfter(0, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Test1", actual.Head().Value)
    assert.Equal(t, "Another", actual.Head().Next.Value)
}

func TestSinglyLinkedList_InsertNodeAfterAtIndexOneOneItemList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeAfter(1, "Another")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_InsertNodeAfterAtIndexTwoOneList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeAfter(2, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSinglyLinkedList_InsertNodeAfterAtIndexZeroTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    err := actual.InsertNodeAfter(0, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 3, actual.Count())
    assert.Equal(t, "Test1", actual.Head().Value)
    assert.Equal(t, "Another", actual.Head().Next.Value)
    assert.Equal(t, "Test2", actual.Head().Next.Next.Value)
}

func TestSinglyLinkedList_InsertNodeAfterAtIndexOneTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    err := actual.InsertNodeAfter(1, "Another")
    assert.Nil(t, err)
    assert.Equal(t, 3, actual.Count())
    assert.Equal(t, "Test1", actual.Head().Value)
    assert.Equal(t, "Test2", actual.Head().Next.Value)
    assert.Equal(t, "Another", actual.Head().Next.Next.Value)
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
    assert.Equal(t, "Test2", actual.Head().Value)
}

func TestSinglyLinkedList_RemoveNode_AtIndexOneWithTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    err := actual.RemoveNode(1)
    assert.Nil(t, err)
    assert.Equal(t, 1, actual.Count())
    assert.Equal(t, "Test1", actual.Head().Value)
}

func TestSinglyLinkedList_RemoveNode_AtIndexOneWithThreeItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(3)
    err := actual.RemoveNode(1)
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Test1", actual.Head().Value)
    assert.Equal(t, "Test3", actual.Head().Next.Value)
}

func TestSinglyLinkedList_RemoveNode_AtIndextTwoWithThreeItemsList_ReturnSuccess(t *testing.T) {

    actual := getListOfElements(3)
    err := actual.RemoveNode(2)
    assert.Nil(t, err)
    assert.Equal(t, 2, actual.Count())
    assert.Equal(t, "Test1", actual.Head().Value)
    assert.Equal(t, "Test2", actual.Head().Next.Value)
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

func TestSinglyLinkedList_PopFront_EmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.PopFront()
    assert.NotNil(t, err)
    assert.Equal(t, "the list is empty", err.Error())
}

func TestSinglyLinkedList_PopFront_OneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.PopFront()
    assert.Nil(t, err)
    assert.True(t, actual.Empty())
}

func TestSinglyLinkedList_PopFront_TwoItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    err := actual.PopFront()
    assert.Nil(t, err)
    assert.False(t, actual.Empty())
    assert.Equal(t, "Test2", actual.Head().Value)
}

func TestSinglyLinkedList_Unique_WithEmptyList_ReturnSuccess(t *testing.T) {
    list := MakeSingleLinkedList()
    list.Unique()
    assert.True(t, list.Empty())
}

func TestSinglyLinkedList_Unique_WithOneItemList_ReturnSuccess(t *testing.T) {
    list := getListOfElements(1)
    list.Unique()
    assert.Equal(t, 1, list.Count())
}

func TestSinglyLinkedList_Unique_WithTwoDifferentItemsList_ReturnSuccess(t *testing.T) {
    list := getListOfElements(2)
    list.Unique()
    assert.Equal(t, 2, list.Count())
}

func TestSinglyLinkedList_Unique_WithTwoSameItemList_ReturnSuccess(t *testing.T) {
    list := MakeSingleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test")
    list.Unique()
    assert.Equal(t, 1, list.Count())
    assert.Equal(t, "Test", list.Head().Value)
}

func TestSinglyLinkedList_Unique_WithTwoSameItemTwoTimesList_ReturnSuccess(t *testing.T) {
    list := MakeSingleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test")
    list.AddNode("Hello")
    list.AddNode("Hello")
    list.Unique()
    assert.Equal(t, 2, list.Count())
    assert.Equal(t, "Test", list.Head().Value)
    assert.Equal(t, "Hello", list.Head().Next.Value)
}

func TestSinglyLinkedList_Unique_WithTwoSameItemOneDiffAndThreeTimesList_ReturnSuccess(t *testing.T) {
    list := MakeSingleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test")
    list.AddNode("Hi")
    list.AddNode("Hello")
    list.AddNode("Hello")
    list.AddNode("Hello")
    list.Unique()
    assert.Equal(t, 3, list.Count())
    assert.Equal(t, "Test", list.Head().Value)
    assert.Equal(t, "Hi", list.Head().Next.Value)
    assert.Equal(t, "Hello", list.Head().Next.Next.Value)
}

