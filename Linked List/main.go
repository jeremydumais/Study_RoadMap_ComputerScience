// A demo of multiple forms of linked list
package main

import (
	"fmt"
	"sort"
	"study/linkedlistdemo/singlylinkedlist"
)

func main() {
    var list = singlylinkedlist.MakeSingleLinkedList();
    list.AddNode("ABC")
    list.AddNode("ABC")
    list.AddNode("Test")
    list.AddNode("Test")
    list.AddNode("123")
    sort.Sort(list)
    for i := 0; i < list.Len(); i++ {
        item, _ := list.GetNode(i)
        fmt.Println(item.Value)
    }
    fmt.Println()
    sort.Sort(sort.Reverse(list))
    list.Unique()
    for i := 0; i < list.Len(); i++ {
        item, _ := list.GetNode(i)
        fmt.Println(item.Value)
    }
}
