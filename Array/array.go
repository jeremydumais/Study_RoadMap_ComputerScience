/*
   Print an array of 10 integer
*/

package main

import (
	"fmt"
)

func main() {
	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i, v := range arr1 {
		fmt.Printf("The value at index %d is %d\n", i, v)
	}
}
