/*Make a program that counts from 1 to 10.*/

package main

import "fmt"

func main(){
	fmt.Println("This is section loops")
	var sum =0

	for i:=1; i<=10;i++{
		sum = sum+i
	}	
	fmt.Println("The sum of numbers from 1 to 10 is:",sum)
}