package main

import "fmt"

func main(){
	city := make(map[int]string)
	city[1] = "One"
	city[2] = "Two"
	city[3] = "Dummy"
	city[3] = "Three"
	fmt.Println(city)

	fmt.Println("\nAdding values to map")
	city[4]= "Four"
	fmt.Println(city)

	fmt.Println("\nDeleting values to map")
	delete(city,3)
	fmt.Println(city)

	fmt.Println("\nGetting a specific value to map")sss
	val := city[1]
	fmt.Println(val)
}