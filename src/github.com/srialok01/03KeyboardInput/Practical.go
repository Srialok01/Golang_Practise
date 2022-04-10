/*
Get a number from the console and check if it’s between 1 and 10.
*/
package main

import ("fmt"
		"bufio"
		"os"
		"strconv"
		"strings")

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number between 1 -10")
	entered_val, _ := reader.ReadString('\n')
	fmt.Println("You have entered :",entered_val)
	fmt.Printf("The type of variable is %T\n",entered_val)

	int_val, err := strconv.ParseInt(strings.TrimSpace(entered_val),32,32)
	fmt.Printf("The type of variable is %T ",int_val)

	if err!=nil{
		fmt.Println(err)
	}
	if int_val >=1 && int_val<10{
		fmt.Println("Yes the value is less than 10")
	}else{
		fmt.Println("No the value is not less than 10")

	}

}