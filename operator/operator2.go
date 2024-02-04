package main
import "fmt"

func main(){
	var left = false
	var right = true

	var left_and_right = left && right
	fmt.Println(left_and_right)

	var left_or_right = left || right
	fmt.Println(left_or_right)

	var negationLeft = !left
	fmt.Println(negationLeft)
}
