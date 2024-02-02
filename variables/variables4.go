package main

import "fmt"

func main(){
	name := new(string)
	*name = "Tama"

	fmt.Println(name)
	fmt.Println(*name)
}
