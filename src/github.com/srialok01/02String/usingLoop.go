package main

import "fmt"

func main(){
	str := "This section is for string"
	for key, val := range(str){
		fmt.Printf("For key %v The character is '%c' and value is '%v'\n",key,val,val)
	}
}