package main

import "testing"

func TestSinglyLinkedList_Make_ReturnNilHead(t *testing.T) {
    var actual = MakeSingleLinkedList()
    if actual.GetHead() != nil {
        t.Errorf("expected head to be nil")
    }
}
