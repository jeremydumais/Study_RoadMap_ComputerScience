//This is a comment
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
