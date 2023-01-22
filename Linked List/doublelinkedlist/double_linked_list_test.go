// Package doublelinkedlist contains all the objects of a doubly linked list.
package doublelinkedlist

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "fmt"
)

func getListOfElements(numberOfNodes int) DoubleLinkedList {
    list := MakeDoubleLinkedList() 
    for i := 0; i < numberOfNodes; i++ {
        nodeValue := fmt.Sprintf("Test%d", i+1)
        list.AddNode(nodeValue) 
    }
    return list
}

func checkElementsLinkage(t *testing.T, expectedElement []string, list DoubleLinkedList) { 
    assert.Equal(t, len(expectedElement), list.Len())
    currentNode := list.Head();
    for i, elem := range expectedElement {
        assert.NotNil(t, elem)
        assert.Equal(t, elem, currentNode.Value())
        // If it's the first element
        if i == 0 {
            assert.Nil(t, currentNode.Prev())
            assert.Same(t, list.Head(), currentNode)
        } else {
            assert.NotNil(t, currentNode.Prev())
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

func TestDoubleLinkedList_Make_ReturnNilHead(t *testing.T) {
    actual := MakeDoubleLinkedList()
    assert.Nil(t, actual.Head())
    assert.Nil(t, actual.Tail())
}

func TestDoubleLinkedList_Head_EmptyList_ReturnNil(t *testing.T) {
    assert.Nil(t, MakeDoubleLinkedList().Head())
}

func TestDoubleLinkedList_Tail_EmptyList_ReturnNil(t *testing.T) {
    assert.Nil(t, MakeDoubleLinkedList().Tail())
}

func TestDoubleLinkedList_Head_OneNodeList_ReturnNotNil(t *testing.T) {
    actual := getListOfElements(1)
    assert.NotNil(t, actual.Head())
    assert.Equal(t, "Test1", actual.Head().Value())
}

func TestDoubleLinkedList_Tail_OneNodeList_ReturnNotNil(t *testing.T) {
    actual := getListOfElements(1)
    assert.NotNil(t, actual.Tail())
    assert.Equal(t, "Test1", actual.Tail().Value())
}

func TestDoubleLinkedList_Len_EmptyList_Return0(t *testing.T) {
    assert.Equal(t, 0, MakeDoubleLinkedList().Len())
}

func TestDoubleLinkedList_Len_OneNodeList_Return1(t *testing.T) {
    assert.Equal(t, 1, getListOfElements(1).Len())
}

func TestDoubleLinkedList_Len_TwoNodeList_Return2(t *testing.T) {
    assert.Equal(t, 2, getListOfElements(2).Len())
}

func TestDoubleLinkedList_GetNode_WithIndexMinus1EmptyList_ReturnError(t *testing.T) {
    actual := MakeDoubleLinkedList()
    res, err := actual.GetNode(-1)
    assert.Nil(t, res)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_GetNode_WithIndex0EmptyList_ReturnError(t *testing.T) {
    actual := MakeDoubleLinkedList()
    res, err := actual.GetNode(0)
    assert.Nil(t, res)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_GetNode_WithIndex0OneElementList_ReturnNode(t *testing.T) {
    actual := getListOfElements(1) 
    res, _ := actual.GetNode(0)
    assert.NotNil(t, res)
    assert.Equal(t, "Test1", res.Value())
}

func TestDoubleLinkedList_GetNode_WithIndex1TwoElementsList_ReturnNode(t *testing.T) {
    actual := getListOfElements(2)
    res, _ := actual.GetNode(1)
    assert.NotNil(t, res)
    assert.Equal(t, "Test2", res.Value())
}

func TestDoubleLinkedList_GetNode_WithIndex2TwoElementsList_ReturnError(t *testing.T) {
    actual := getListOfElements(2)
    res, err := actual.GetNode(2)
    assert.Nil(t, res)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_Empty_WithEmptyList_ReturnTrue(t *testing.T) {
    assert.True(t, getListOfElements(0).Empty())
}

func TestDoubleLinkedList_Empty_WithOneItemList_ReturnTrue(t *testing.T) {
    assert.False(t, getListOfElements(1).Empty())
}

func TestDoubleLinkedList_Empty_WithTwoItemList_ReturnTrue(t *testing.T) {
    assert.False(t, getListOfElements(2).Empty())
}

func TestDoubleLinkedList_AddNode_EmptyList_ReturnSuccess(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test1")
    checkElementsLinkage(t, []string{"Test1"}, actual)
}

func TestDoubleLinkedList_AddNode_OneNodeList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    actual.AddNode("Test2")
    checkElementsLinkage(t, []string{"Test1", "Test2"}, actual)
}

func TestDoubleLinkedList_AddNode_TwoNodeList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    actual.AddNode("Test3")
    checkElementsLinkage(t, []string{"Test1", "Test2", "Test3"}, actual)
}

func TestDoubleLinkedList_InsertNodeBeforeAtIndexZeroEmptyList_ReturnSuccess(t *testing.T) {
    actual := MakeDoubleLinkedList()
    assert.Nil(t, actual.InsertNodeBefore(0, "Test"))
    checkElementsLinkage(t, []string{"Test"}, actual)
}

func TestDoubleLinkedList_InsertNodeBeforeAtIndexOneEmptyList_ReturnError(t *testing.T) {
    actual := MakeDoubleLinkedList()
    err := actual.InsertNodeBefore(1, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_InsertNodeBeforeAtIndexZeroOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.InsertNodeBefore(0, "Another"))
    checkElementsLinkage(t, []string{"Another", "Test1"}, actual)
}

func TestDoubleLinkedList_InsertNodeBeforeAtIndexOneOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.InsertNodeBefore(1, "Another"))
    checkElementsLinkage(t, []string{"Test1", "Another"}, actual)
}

func TestDoubleLinkedList_InsertNodeBeforeAtIndexTwoOneList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeBefore(2, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_InsertNodeBeforeAtIndexOneTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.InsertNodeBefore(1, "Another"))
    checkElementsLinkage(t, []string{"Test1", "Another", "Test2"}, actual)
}

func TestDoubleLinkedList_InsertNodeAfterAtIndexZeroEmptyList_ReturnError(t *testing.T) {
    actual := MakeDoubleLinkedList()
    err := actual.InsertNodeAfter(0, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_InsertNodeAfterAtIndexOneEmptyList_ReturnError(t *testing.T) {
    actual := MakeDoubleLinkedList()
    err := actual.InsertNodeAfter(1, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_InsertNodeAfterAtIndexZeroOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.InsertNodeAfter(0, "Another"))
    checkElementsLinkage(t, []string{"Test1", "Another"}, actual)
}

func TestDoubleLinkedList_InsertNodeAfterAtIndexOneOneItemList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeAfter(1, "Another")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_InsertNodeAfterAtIndexTwoOneList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.InsertNodeAfter(2, "Test")
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_InsertNodeAfterAtIndexZeroTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.InsertNodeAfter(0, "Another"))
    checkElementsLinkage(t, []string{"Test1", "Another", "Test2"}, actual)
}

func TestDoubleLinkedList_InsertNodeAfterAtIndexOneTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.InsertNodeAfter(1, "Another"))
    checkElementsLinkage(t, []string{"Test1", "Test2", "Another"}, actual)
}

func TestDoubleLinkedList_RemoveNode_AtIndexZeroWithEmptyList_ReturnError(t *testing.T) {
    actual := MakeDoubleLinkedList()
    err := actual.RemoveNode(0)
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_RemoveNode_AtIndexZeroWithOneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.RemoveNode(0))
    assert.True(t, actual.Empty())
}

func TestDoubleLinkedList_RemoveNode_AtIndexZeroWithTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.RemoveNode(0))
    checkElementsLinkage(t, []string{"Test2"}, actual)
}

func TestDoubleLinkedList_RemoveNode_AtIndexOneWithTwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.RemoveNode(1))
    checkElementsLinkage(t, []string{"Test1"}, actual)
}

func TestDoubleLinkedList_RemoveNode_AtIndexOneWithThreeItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(3)
    assert.Nil(t, actual.RemoveNode(1))
    checkElementsLinkage(t, []string{"Test1", "Test3"}, actual)
}

func TestDoubleLinkedList_RemoveNode_AtIndextTwoWithThreeItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(3)
    assert.Nil(t, actual.RemoveNode(2))
    checkElementsLinkage(t, []string{"Test1", "Test2"}, actual)
}

func TestDoubleLinkedList_RemoveNode_AtIndexOneWithOneItemList_ReturnError(t *testing.T) {
    actual := getListOfElements(1)
    err := actual.RemoveNode(1)
    assert.NotNil(t, err)
    assert.Equal(t, "index out of range", err.Error())
}

func TestDoubleLinkedList_Clear_EmptyList_ReturnSuccess(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.Clear()
    assert.Equal(t, 0, actual.Len())
    assert.Nil(t, actual.Head())
    assert.Nil(t, actual.Tail())
}

func TestDoubleLinkedList_Clear_TwoItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    actual.Clear()
    assert.Equal(t, 0, actual.Len())
    assert.Nil(t, actual.Head())
    assert.Nil(t, actual.Tail())
}

func TestDoubleLinkedList_Clear_ThreeItemsList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(0)
    actual.Clear()
    assert.Equal(t, 0, actual.Len())
    assert.Nil(t, actual.Head())
    assert.Nil(t, actual.Tail())
}

func TestDoubleLinkedList_PopFront_EmptyList_ReturnError(t *testing.T) {
    actual := MakeDoubleLinkedList()
    err := actual.PopFront()
    assert.NotNil(t, err)
    assert.Equal(t, "the list is empty", err.Error())
}

func TestDoubleLinkedList_PopFront_OneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.PopFront())
    assert.True(t, actual.Empty())
}

func TestDoubleLinkedList_PopFront_TwoItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.PopFront())
    assert.False(t, actual.Empty())
    assert.Equal(t, "Test2", actual.Head().Value())
}

func TestDoubleLinkedList_PopBack_EmptyList_ReturnError(t *testing.T) {
    actual := MakeDoubleLinkedList()
    err := actual.PopBack()
    assert.NotNil(t, err)
    assert.Equal(t, "the list is empty", err.Error())
}

func TestDoubleLinkedList_PopBack_OneItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(1)
    assert.Nil(t, actual.PopBack())
    assert.True(t, actual.Empty())
}

func TestDoubleLinkedList_PopBack_TwoItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(2)
    assert.Nil(t, actual.PopBack())
    assert.False(t, actual.Empty())
    assert.Equal(t, "Test1", actual.Head().Value())
    assert.Same(t, actual.Head(), actual.Tail())
}

func TestDoubleLinkedList_PopBack_ThreeItemList_ReturnSuccess(t *testing.T) {
    actual := getListOfElements(3)
    assert.Nil(t, actual.PopBack())
    assert.Equal(t, 2, actual.Len())
    assert.Equal(t, "Test1", actual.Head().Value())
    node2, _ := actual.GetNode(1)
    assert.Equal(t, "Test2", node2.Value())
}

func TestDoubleLinkedList_Unique_WithEmptyList_ReturnSuccess(t *testing.T) {
    list := MakeDoubleLinkedList()
    list.Unique()
    assert.True(t, list.Empty())
}

func TestDoubleLinkedList_Unique_WithOneItemList_ReturnSuccess(t *testing.T) {
    list := getListOfElements(1)
    list.Unique()
    assert.Equal(t, 1, list.Len())
}

func TestDoubleLinkedList_Unique_WithTwoDifferentItemsList_ReturnSuccess(t *testing.T) {
    list := getListOfElements(2)
    list.Unique()
    assert.Equal(t, 2, list.Len())
}

func TestDoubleLinkedList_Unique_WithTwoSameItemList_ReturnSuccess(t *testing.T) {
    list := MakeDoubleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test")
    list.Unique()
    checkElementsLinkage(t, []string{"Test"}, list)
}

func TestDoubleLinkedList_Unique_WithTwoSameItemTwoTimesList_ReturnSuccess(t *testing.T) {
    list := MakeDoubleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test")
    list.AddNode("Hello")
    list.AddNode("Hello")
    list.Unique()
    checkElementsLinkage(t, []string{"Test", "Hello"}, list)
}

func TestDoubleLinkedList_Unique_WithTwoDiffItemsList_ReturnSuccess(t *testing.T) {
    list := MakeDoubleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test2")
    list.Unique()
    checkElementsLinkage(t, []string{"Test", "Test2"}, list)
}

func TestDoubleLinkedList_Unique_WithTwoSameItemOneDiffAndThreeTimesList_ReturnSuccess(t *testing.T) {
    list := MakeDoubleLinkedList()
    list.AddNode("Test")
    list.AddNode("Test")
    list.AddNode("Hi")
    list.AddNode("Hello")
    list.AddNode("Hello")
    list.AddNode("Hello")
    list.Unique()
    checkElementsLinkage(t, []string{"Test", "Hi", "Hello"}, list)
}

func TestDoubleLinkedList_Less_WithZeroAndOneEmptyList_ReturnFalse(t *testing.T) {
    actual := MakeDoubleLinkedList()
    assert.False(t, actual.Less(0, 1))
}

func TestDoubleLinkedList_Less_WithZeroAndOneOneItemList_ReturnFalse(t *testing.T) {
    actual := getListOfElements(1)
    assert.False(t, actual.Less(0, 1))
}

func TestDoubleLinkedList_Less_WithZeroAndOneTwoSameItemList_ReturnFalse(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test")
    actual.AddNode("Test")
    assert.False(t, actual.Less(0, 1))
}

func TestDoubleLinkedList_Less_WithZeroAndOneTwoDiffItemList_ReturnTrue(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test0")
    actual.AddNode("Test1")
    assert.True(t, actual.Less(0, 1))
}

func TestDoubleLinkedList_Less_WithZeroAndOneTwoDiffItemReverseList_ReturnFalse(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test1")
    actual.AddNode("Test0")
    assert.False(t, actual.Less(0, 1))
}

func TestDoubleLinkedList_Less_WithZeroAndOneTwoDiffItemReverseIndexList_ReturnFalse(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test0")
    actual.AddNode("Test1")
    assert.False(t, actual.Less(1, 0))
}

func TestDoubleLinkedList_Less_WithIOutOfRange_ReturnFalse(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test0")
    assert.False(t, actual.Less(1, 0))
}

func TestDoubleLinkedList_Less_WithJOutOfRange_ReturnFalse(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test0")
    assert.False(t, actual.Less(0, 1))
}

func TestDoubleLinkedList_Swap_WithZeroAndOneIndex_ReturnSuccess(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test0")
    actual.AddNode("Test1")
    actual.Swap(0, 1)
    assert.Equal(t, "Test1", actual.Head().Value())
    assert.Equal(t, "Test0", actual.Head().Next().Value())
}

func TestDoubleLinkedList_Swap_WithOneAndZeroIndex_ReturnSuccess(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test0")
    actual.AddNode("Test1")
    actual.Swap(1, 0)
    assert.Equal(t, "Test1", actual.Head().Value())
    assert.Equal(t, "Test0", actual.Head().Next().Value())
}

func TestDoubleLinkedList_Swap_WithIOutOfRange_ReturnSuccess(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test0")
    actual.Swap(1, 0)
    assert.Equal(t, 1, actual.Len())
    assert.Equal(t, "Test0", actual.Head().Value())
}

func TestDoubleLinkedList_Swap_WithJOutOfRange_ReturnSuccess(t *testing.T) {
    actual := MakeDoubleLinkedList()
    actual.AddNode("Test0")
    actual.Swap(0, 1)
    assert.Equal(t, 1, actual.Len())
    assert.Equal(t, "Test0", actual.Head().Value())
}

