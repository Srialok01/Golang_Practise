package main

import ("fmt"
		"os"
		"bufio"
)


func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Name: ")
	input,_ := reader.ReadString('\n')
	fmt.Println("The name Entered is: ",input)
}