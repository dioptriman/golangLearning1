package main

import "fmt"

func main(){

	/*
		Basic if-else
	*/

	var point = 8

	if point == 10 {
		fmt.Println("Nilai anda sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir saja lulus")
	} else {
		fmt.Println("Nilai anda kurang")
	}

}
