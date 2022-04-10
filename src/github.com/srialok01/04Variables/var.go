package main

import ("fmt"
		"bufio"
		"os"
		"strings"
		"strconv"
	)

func main(){

	//  Getting  Age
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the Age :")
	Age,_ := reader.ReadString('\n')
	fmt.Println("Your Age is ",Age)
	int_age ,_ := strconv.ParseFloat(strings.TrimSpace(Age),10)
	fmt.Printf("The type of variable is %T and value is %v",int_age,int_age)

	//Getting date of birth
	fmt.Print("Enter the Date of Birth in format dd-mm :")
	dob,_ := reader.ReadString('\n')
	fmt.Println("Your DOB is ",dob)
	dd_mm := strings.TrimSpace(dob)
	fmt.Print(dd_mm)

	//split the string
	str :=strings.Split(dd_mm,"-")
	fmt.Println("\nYour str is ",str)
	date := str[0]
	month :=str[1]
	year := 2022 - int_age
	fmt.Println("Your date is ",date)
	fmt.Println("Your month is ",month)
	fmt.Println("Your year is ",year)

	fmt.Printf("Your date of Birth is %v-%v-%v",date,month,year)



}