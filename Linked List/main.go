// A demo of multiple forms of linked list
package main

import ( 
    "fmt"
    "os"
    "study/linkedlistdemo/singlylinkedlist"
)

func main() {
    var list = singlylinkedlist.MakeSingleLinkedList();
    if list.Head() == nil {
        fmt.Printf("The list is empty\n")
        os.Exit(0)
    }
}
