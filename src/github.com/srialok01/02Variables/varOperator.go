/*
VAR keyword is used to declare a variable with type and can be used globally
*/
package main

import "fmt"

var is_working bool
var a int = 32
// var b int 
func main(){
	 var b string = "Hello"

	 fmt.Printf("\nThe global value of is_working is '%v' and type is :'%T'\n",is_working,is_working)
	 fmt.Printf("\nThe global value of a is '%v' and type is : '%T'\n ",a,a)
	 fmt.Printf("\nThe local value of b is '%v' and type is : '%T'\n ",b,b)
	 fmt.Println()

}