package main

import "fmt"

func main(){
	var names = [4]string {"apple", "gray", "white", "grape"}
	fmt.Println(names)

	var anotherArray [4]string
	anotherArray[0] = "Lapar"
	anotherArray[1] = "Mau"
	anotherArray[2] = "Makan"
	anotherArray[3] = "Nasi"

	fmt.Println(anotherArray[0], anotherArray[1], anotherArray[2], anotherArray[3])
}
