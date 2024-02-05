package main

import "fmt"

func main(){
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	//With Index
	for i, fruit := range fruits {
    		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	//Without Index
	for _, fruit := range fruits {
    		fmt.Printf("nama buah : %s\n", fruit)
	}
}
