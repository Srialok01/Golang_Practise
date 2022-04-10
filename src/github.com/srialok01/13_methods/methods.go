/* Create a method that sums two numbers */

package main

import ("fmt"
		"bufio"
		"strconv"
		"strings"
		"os"
)

func main(){
	fmt.Println("This section belongs to methods section !!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Two Numbers ")
	str_input1,_ := reader.ReadString('\n')
	str_input2,_ := reader.ReadString('\n')

	int_val1,err := strconv.ParseFloat(strings.TrimSpace(str_input1),32)
	if err !=nil{
		fmt.Println(err)
	}
	int_val2,err := strconv.ParseFloat(strings.TrimSpace(str_input2),32)
	if err !=nil{
		fmt.Println(err)
	}

	val := sum(int_val1,int_val2)
	fmt.Println("The value is :",val)

}

func sum(num1,num2 float64)float64{
	var result float64
	result = num1 + num2
	return result

}