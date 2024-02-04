package main

import "fmt"

func main(){

	/*
		Temporary Variable in if-else
	*/

	var point = 8840.0


	//Temporary Variables hanya bisa dengan inference type
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}
