/* Deleting an item from a slice 
Input = ['Apple','Orange','Onion','Banana','Grapes']
Output = ['Apple','Orange','Banana','Grapes']
*/
package main

import "fmt"

func main(){
	var veg_slice = []string{}
	veg_slice = append(veg_slice,"Apple","Orange","Onion","Banana","Grapes")
	fmt.Println(veg_slice)
	// var index int =3
	veg_slice = append(veg_slice[:index] ...string)
	fmt.Println(veg_slice)
}