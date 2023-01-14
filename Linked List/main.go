// A demo of multiple forms of linked list
package main

import ( 
    "fmt"
    "os"
)

func main() {
    var list = MakeSingleLinkedList();
    if list.Head() == nil {
        fmt.Printf("Erreur\n")
        os.Exit(1)
    }
}
