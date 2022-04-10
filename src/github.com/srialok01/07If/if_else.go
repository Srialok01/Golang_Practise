/* Make a program that divides x by 2 if it’s greater than 0 */

package main

import ("fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
		)

func main(){
	fmt.Println("Welcome to If-else Section")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the Number: ")
	input, _ := reader.ReadString('\n')
	int_input,err := strconv.ParseFloat(strings.TrimSpace(input),32)
	if err !=nil{
		fmt.Println("Please check the input provided")
		panic(err)
	}
	if int_input != 0{
		val := int_input/2
		fmt.Println("The value after divided by 2 is :",val)
	}else{
		fmt.Println("The value is 0 hence can't divide it")
	}
}