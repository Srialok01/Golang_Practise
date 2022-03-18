/* This is also known as shorthand operator
denoted as := 
* This is used to declare and assign value at once 
* While using walrus operator you don't need to provide var keyword explictly
* Walrus operator can not be used globally
*/

package main
import "fmt"

func main(){

	x := 10
	fmt.Println("\nThe value of x is ",x)

	x = 5
	fmt.Println("\nThe updated value of x is ",x)

}