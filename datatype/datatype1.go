package main

import "fmt"

func main(){

	/*
		Integer (non-decimal) value
	*/

	//Unsigned Integer
	var positiveNumber8byte uint8 = 10 //0-255
	var positiveNumber16byte uint16 = 1000 //0-65535
	var positiveNumber32byte uint32 = 70000 //0-4294967295
	var positiveNumber64byte uint64 = 45000000 //0-18446744073709551615
	var positiveNumber uint = 20 //uint32 or uint64 tergantung dari nilai
	var positiveNumberbyte byte = 19 //sama dengan uint8

	//Signed Integer
	var number int8 = -100 //-128 - 127
	var number2 int16 = -1000 //-32768 - 32767
	var number3 int32 = 70000 //-2147483648 - 2147483647
	var number4 int64 = -922337203685477 //-9223372036854775808 - 9223372036854775807
	var number5 int = 90 //int32 or int64 tergantung dari nilai
	var number6 rune = 100 //sama dengan int32


	fmt.Println(positiveNumber8byte)
	fmt.Println(positiveNumber16byte)
	fmt.Println(positiveNumber32byte)
	fmt.Println(positiveNumber64byte)
	fmt.Println(positiveNumber)
	fmt.Println(positiveNumberbyte)

	fmt.Println(number)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(number4)
	fmt.Println(number5)
	fmt.Println(number6)
}
