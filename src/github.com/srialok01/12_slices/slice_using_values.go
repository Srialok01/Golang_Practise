/* 2.	Initialize slice using values */

package main

import "fmt"

func main(){

	fmt.Println("This section is for slices")
	var int_slice = []int{1,2,3,4}
	fmt.Println("The initial slice is :",int_slice)
	fmt.Println("Now Appending Slice")
	int_slice = append(int_slice,5,6,7,8,9)
	fmt.Println("The final slice is :",int_slice)

}