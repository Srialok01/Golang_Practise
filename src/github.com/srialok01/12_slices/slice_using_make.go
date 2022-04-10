/* Creating and appending slice using make*/

package main

import "fmt"

func main(){
	val := make([]int, 4)
	fmt.Println("The initial val is :",val)
	val[0] = 2
	val[1] = 3
	val[2] = 4
	val[3] = 5
	fmt.Println("The updated value of val is",val)
	// val[4] = 6	//This will throw an error

	val = append(val,6,7,8,9)
	fmt.Println("This is final value ",val)

}