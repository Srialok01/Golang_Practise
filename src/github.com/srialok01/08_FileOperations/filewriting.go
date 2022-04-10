package main
import ("fmt"
		"os"
		"bufio")

func main(){
	fmt.Println("Taking input from user!!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the details that needs to be appended in the file")
	str_input, _ := reader.ReadString('\n')
	fmt.Println("The text entered is", str_input)

	//  Checking if file exists 
	_,err := os.Stat("Demo.txt")
	if err==nil{
		fmt.Println("File Present")
		err = os.Remove("Demo.txt")
		if err !=nil{
			panic(err)
		}
	}else{
		fmt.Println("File Not Present")
		// Creating file
	}
	f,_ := os.Create("Demo.txt")
	defer f.Close()
	writer := bufio.NewWriter(f)
	stat,_ := writer.WriteString(str_input)
	fmt.Println("The text wrote in file is",stat)
	writer.Flush()
	
}