package main

import ("fmt"
		"math/rand"
		"time")


func random(min int, max int) int{
	return rand.Intn(max-min)+min
}		


func main(){
	fmt.Println("Welcome to random section")
	rand.Seed(time.Now().UnixNano())
	random := random(1,6)
	fmt.Println(random)
	
}
