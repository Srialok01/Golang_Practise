/* Create a struct house with variables noRooms, price and city */
package main

import "fmt"

func main(){

	studio := house{1,1500.00,"Ghazipur"}
	fmt.Println(studio)

	villa := house{5, 13000.45,"Varanasi"}

	fmt.Println(villa)
	fmt.Println(studio.noRooms, studio.price, studio.city )
}

type house struct{
	noRooms int
	price float64
	city string
}