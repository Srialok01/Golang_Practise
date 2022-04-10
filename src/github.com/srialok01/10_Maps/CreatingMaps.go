/* Creating a map using make method*/
package main

import "fmt"

func main(){
	fmt.Println("Creating a map using make function")

	city := make(map[string]string)
		city["A"] = "America"
		city["B"] = "Boston"
		city["C"] = "china"	
	fmt.Println(city)

	fmt.Println("\nCreating a map directly")
	ct := map[int]string{
		1: "Apple",
		2: "Ball",
		3: "Cat",
	}
	fmt.Println(ct)

}