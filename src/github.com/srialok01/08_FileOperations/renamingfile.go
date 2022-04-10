package main

import ("fmt"
		"os")

func main(){
	fmt.Println("Welcome to File Rename section !!")
	source := "Demo.txt"
	destination := "golang.log"
	os.Rename(source,destination) 

}