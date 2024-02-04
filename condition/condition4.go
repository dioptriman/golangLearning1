package main

import "fmt"

func main(){

	/*
		Switch - Case w/ multiple default result
	*/

	var point = 1

	switch point {
		case 8:
    			fmt.Println("perfect")
		case 7, 6, 5, 4:
    			fmt.Println("awesome")
		default:
    			{
        			fmt.Println("not bad")
        			fmt.Println("you can be better!")
    			}
	}
}
