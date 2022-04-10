/* Declaring and appending value into a slice */

package main

import "fmt"

func main(){
	fmt.Println("Welcome to Slice section")

	var str_slice[] string
	fmt.Printf("The value of slice is%d and type is %T",str_slice,str_slice)

	fmt.Println("\nAppending values into a slice")
	str_slice = append(str_slice,"Hello", "My","Name","is","Alok")
	fmt.Println("The slice is :",str_slice)
}