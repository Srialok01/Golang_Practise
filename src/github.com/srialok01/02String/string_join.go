package main

import ("fmt"
		"strings")

func main (){
	fmt.Println("This section is for joining string")
	var string1 = "Alok"
	var string2 = "srivastav"
	str_demo := strings.Join([]string{string1,string2},":")
	fmt.Println(str_demo)
}