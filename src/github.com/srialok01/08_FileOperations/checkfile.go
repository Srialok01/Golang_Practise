package main

import ("fmt"
		"os")

func main(){
	fmt.Println("Welcome to File Section !!")
	_,err := os.Stat("demo1.log")
	if err ==nil{
		fmt.Println("File exists")

	}else{
		fmt.Println("File not Present")
	}
}