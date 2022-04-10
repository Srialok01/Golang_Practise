/* Create an array with the number 0 to 10*/

package main

import "fmt"

func main(){

	arr := [10]int {1,2,3,4,5,6,7,8,9}
	fmt.Println(arr)
	fmt.Printf("The type of arr is :%T",arr)
	fmt.Println("\n==============================\n")
	str_arr := [3]string {"Alok","kumar","MadhuSudan"}
	fmt.Println(str_arr)
	fmt.Printf("The type of arr is :%T",str_arr)

}