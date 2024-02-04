package main

import "fmt"

func main(){

	/*
		For but like Do - While
	*/

	var i = 0

	for{
		fmt.Println("Angka", i)
		i++

		if i == 5 {
			break
		}
	}
}
