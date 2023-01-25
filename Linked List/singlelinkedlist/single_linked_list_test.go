package singlelinkedlist

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

func checkElementsLinkage(t *testing.T, expectedElement []string, list SingleLinkedList) { 
    assert.Equal(t, len(expectedElement), list.Len())
    currentNode := list.Head();
    for i, elem := range expectedElement {
        assert.NotNil(t, elem)
        assert.Equal(t, elem, currentNode.Value())
        // If it's the first element
        if i == 0 {
            assert.Same(t, list.Head(), currentNode)
        }
        // If it's the last element
        if i == len(expectedElement) - 1 {
            assert.Nil(t, currentNode.Next())
            assert.Same(t, list.Tail(), currentNode)
        } else {
            assert.NotNil(t, currentNode.Next())
        }
        currentNode = currentNode.Next()
    }
}


func TestSingleLinkedList_Make_ReturnNilHead(t *testing.T) {
    actual := MakeSingleLinkedList()
    assert.Nil(t, actual.Head())
}

func TestSingleLinkedList_Tail_EmptyList_ReturnNil(t *testing.T) {
    assert.Nil(t, MakeSingleLinkedList().Tail())
}

func TestSingleLinkedList_Tail_OneItemList_ReturnHead(t *testing.T) {
    list := getListOfElements(1)
    tail := list.Tail()
    assert.NotNil(t, tail)
    assert.Equal(t, list.Head(), tail)
}

func TestSingleLinkedList_Tail_TwoItemsList_ReturnLastNode(t *testing.T) {
    list := getListOfElements(2)
    tail := list.Tail()
    assert.NotNil(t, tail)
    assert.Equal(t, list.Head().Next(), tail)
}

func TestSingleLinkedList_Len_With2TwoElementsList_Return2(t *testing.T) {
    assert.Equal(t, 2, getListOfElements(2).Len())
}

func TestSingleLinkedList_GetNode_WithIndexMinus1EmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    res, err := actual.GetNode(-1)
    assert.Nil(t, res)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_GetNode_WithIndex0EmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    res, err := actual.GetNode(0)
    assert.Nil(t, res)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_GetNode_WithIndex0OneElementList_ReturnNode(t *testing.T) {
    actual := getListOfElements(1) 
    res, _ := actual.GetNode(0)
    assert.NotNil(t, res)
    assert.Equal(t, "Test1", res.Value())
}

func TestSingleLinkedList_GetNode_WithIndex1TwoElementsList_ReturnNode(t *testing.T) {
    actual := getListOfElements(2)
    res, _ := actual.GetNode(1)
    assert.NotNil(t, res)
    assert.Equal(t, "Test2", res.Value())
}

func TestSingleLinkedList_GetNode_WithIndex2TwoElementsList_ReturnError(t *testing.T) {
    actual := getListOfElements(2)
    res, err := actual.GetNode(2)
    assert.Nil(t, res)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_Empty_WithEmptyList_ReturnTrue(t *testing.T) {
    assert.True(t, getListOfElements(0).Empty())
}

func TestSingleLinkedList_Empty_WithOneItemList_ReturnTrue(t *testing.T) {
    assert.False(t, getListOfElements(1).Empty())
}

func TestSingleLinkedList_Empty_WithTwoItemList_ReturnTrue(t *testing.T) {
    assert.False(t, getListOfElements(2).Empty())
}

func TestSingleLinkedList_AddNode_ReturnNonNilHeadAndOneNode(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test")
    checkElementsLinkage(t, []string {"Test"}, actual)
}

func TestSingleLinkedList_AddNode_WithThreeReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test")
    actual.AddNode("Test2")
    actual.AddNode("Test3")
    checkElementsLinkage(t, []string {"Test", "Test2", "Test3"}, actual)
}

func TestSingleLinkedList_InsertNodeBeforeAtIndexZeroEmptyList_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    assert.Nil(t, actual.InsertNodeBefore(0, "Test"))
    checkElementsLinkage(t, []string {"Test"}, actual)
}

func TestSingleLinkedList_InsertNodeBeforeAtIndexOneEmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNodeBefore(1, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_InsertNodeBeforeAtIndexZeroOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.InsertNodeBefore(0, "Another"))
    checkElementsLinkage(t, []string {"Another", "Test1"}, actual)
}

func TestSingleLinkedList_InsertNodeBeforeAtIndexOneOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.InsertNodeBefore(1, "Another"))
    checkElementsLinkage(t, []string {"Test1", "Another"}, actual)
}

func TestSingleLinkedList_InsertNodeBeforeAtIndexTwoOneList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeBefore(2, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_InsertNodeBeforeAtIndexOneTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.InsertNodeBefore(1, "Another"))
    checkElementsLinkage(t, []string {"Test1", "Another", "Test2"}, actual)
}

func TestSingleLinkedList_InsertNodeAfterAtIndexZeroEmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNodeAfter(0, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_InsertNodeAfterAtIndexOneEmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.InsertNodeAfter(1, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_InsertNodeAfterAtIndexZeroOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.InsertNodeAfter(0, "Another"))
    checkElementsLinkage(t, []string {"Test1", "Another"}, actual)
}

func TestSingleLinkedList_InsertNodeAfterAtIndexOneOneItemList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeAfter(1, "Another")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_InsertNodeAfterAtIndexTwoOneList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeAfter(2, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_InsertNodeAfterAtIndexZeroTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.InsertNodeAfter(0, "Another"))
    checkElementsLinkage(t, []string {"Test1", "Another", "Test2"}, actual)
}

func TestSingleLinkedList_InsertNodeAfterAtIndexOneTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.InsertNodeAfter(1, "Another"))
    checkElementsLinkage(t, []string {"Test1", "Test2", "Another"}, actual)
}

func TestSingleLinkedList_RemoveNode_AtIndexZeroWithEmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.RemoveNode(0)
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_RemoveNode_AtIndexZeroWithOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.RemoveNode(0))
    assert.True(t, actual.Empty())
}

func TestSingleLinkedList_RemoveNode_AtIndexZeroWithTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.RemoveNode(0))
    checkElementsLinkage(t, []string {"Test2"}, actual)
}

func TestSingleLinkedList_RemoveNode_AtIndexOneWithTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.RemoveNode(1))
    checkElementsLinkage(t, []string {"Test1"}, actual)
}

func TestSingleLinkedList_RemoveNode_AtIndexOneWithThreeItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(3)
    assert.Nil(t, actual.RemoveNode(1))
    checkElementsLinkage(t, []string {"Test1", "Test3"}, actual)
}

func TestSingleLinkedList_RemoveNode_AtIndextTwoWithThreeItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(3)
    assert.Nil(t, actual.RemoveNode(2))
    checkElementsLinkage(t, []string {"Test1", "Test2"}, actual)
}

func TestSingleLinkedList_RemoveNode_AtIndexOneWithOneItemList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.RemoveNode(1)
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestSingleLinkedList_Clear_EmptyList_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.Clear()
    assert.True(t, actual.Empty())
}

func TestSingleLinkedList_Clear_TwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    actual.Clear()
    assert.True(t, actual.Empty())
}

func TestSingleLinkedList_Clear_ThreeItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(0)
    actual.Clear()
    assert.True(t, actual.Empty())
}

func TestSingleLinkedList_PopFront_EmptyList_ReturnError(t *testing.T) {
    actual := MakeSingleLinkedList()
    err := actual.PopFront()
    assert.NotNil(t, err)
    assert.Equal(t, "the list is empty", err.Error())
}

func TestSingleLinkedList_PopFront_OneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.PopFront())
    assert.True(t, actual.Empty())
}

func TestSingleLinkedList_PopFront_TwoItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.PopFront())
    checkElementsLinkage(t, []string {"Test2"}, actual)
}

func TestSingleLinkedList_Unique_WithEmptyList_ReturnSuccess(t *testing.T) {
    list := MakeSingleLinkedList()
    list.Unique()
    assert.True(t, list.Empty())
}

func TestSingleLinkedList_Unique_WithOneItemList_ReturnSuccess(t *testing.T) {
    list := getListOfElements(1)
    list.Unique()
    assert.Equal(t, 1, list.Len())
}

func TestSingleLinkedList_Unique_WithTwoDifferentItemsList_ReturnSuccess(t *testing.T) {
    list := getListOfElements(2)
    list.Unique()
    assert.Equal(t, 2, list.Len())
}

func TestSingleLinkedList_Unique_WithTwoSameItemList_ReturnSuccess(t *testing.T) {
    list := MakeSingleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test")
    list.Unique()
    checkElementsLinkage(t, []string {"Test"}, list)
}

func TestSingleLinkedList_Unique_WithTwoSameItemTwoTimesList_ReturnSuccess(t *testing.T) {
    list := MakeSingleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test")
    list.AddNode("Hello")
    list.AddNode("Hello")
    list.Unique()
    checkElementsLinkage(t, []string {"Test", "Hello"}, list)
}

func TestSingleLinkedList_Unique_WithTwoDiffItemsList_ReturnSuccess(t *testing.T) {
    list := MakeSingleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test2")
    list.Unique()
    checkElementsLinkage(t, []string {"Test", "Test2"}, list)
}

func TestSingleLinkedList_Unique_WithTwoSameItemOneDiffAndThreeTimesList_ReturnSuccess(t *testing.T) {
    list := MakeSingleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test")
    list.AddNode("Hi")
    list.AddNode("Hello")
    list.AddNode("Hello")
    list.AddNode("Hello")
    list.Unique()
    checkElementsLinkage(t, []string {"Test", "Hi", "Hello"}, list)
}

func TestSingleLinkedList_Less_WithZeroAndOneEmptyList_ReturnFalse(t *testing.T) {
    actual := MakeSingleLinkedList()
    assert.False(t, actual.Less(0, 1))
}

func TestSingleLinkedList_Less_WithZeroAndOneOneItemList_ReturnFalse(t *testing.T) {
    actual := getListOfElements(1)
    assert.False(t, actual.Less(0, 1))
}

func TestSingleLinkedList_Less_WithZeroAndOneTwoSameItemList_ReturnFalse(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test")
    actual.AddNode("Test")
    assert.False(t, actual.Less(0, 1))
}

func TestSingleLinkedList_Less_WithZeroAndOneTwoDiffItemList_ReturnTrue(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test0")
    actual.AddNode("Test1")
    assert.True(t, actual.Less(0, 1))
}

func TestSingleLinkedList_Less_WithZeroAndOneTwoDiffItemReverseList_ReturnFalse(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test1")
    actual.AddNode("Test0")
    assert.False(t, actual.Less(0, 1))
}

func TestSingleLinkedList_Less_WithZeroAndOneTwoDiffItemReverseIndexList_ReturnFalse(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test0")
    actual.AddNode("Test1")
    assert.False(t, actual.Less(1, 0))
}

func TestSingleLinkedList_Less_WithIOutOfRange_ReturnFalse(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test0")
    assert.False(t, actual.Less(1, 0))
}

func TestSingleLinkedList_Less_WithJOutOfRange_ReturnFalse(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test0")
    assert.False(t, actual.Less(0, 1))
}

func TestSingleLinkedList_Swap_WithZeroAndOneIndex_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test0")
    actual.AddNode("Test1")
    actual.Swap(0, 1)
    checkElementsLinkage(t, []string {"Test1", "Test0"}, actual)
}

func TestSingleLinkedList_Swap_WithOneAndZeroIndex_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test0")
    actual.AddNode("Test1")
    actual.Swap(1, 0)
    checkElementsLinkage(t, []string {"Test1", "Test0"}, actual)
}

func TestSingleLinkedList_Swap_WithIOutOfRange_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test0")
    actual.Swap(1, 0)
    checkElementsLinkage(t, []string {"Test0"}, actual)
}

func TestSingleLinkedList_Swap_WithJOutOfRange_ReturnSuccess(t *testing.T) {
    actual := MakeSingleLinkedList()
    actual.AddNode("Test0")
    actual.Swap(0, 1)
    checkElementsLinkage(t, []string {"Test0"}, actual)
}

