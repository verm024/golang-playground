package main

import "fmt"

func main() {
	var str string = "hellow world!!!!!!!!"
	var strtwo = "wkwkland"
	var strthree string

	fmt.Println(str, strtwo, strthree)

	str = "11"
	strtwo = "22"
	strthree = "tiga tiga"
	fmt.Println(str, strtwo, strthree)

	strthree = "tiga tiga 333"

	strfour := "444"
	fmt.Println(strthree, strfour)

	// automatically int -> could be int64 tho, idk, but it is not unsigned type
	numone := 99999999999999999

	// I think the best way to declare int is by using the var
	var numtwo int8 = -12
	var numthree uint8 = 32

	// default type of float is float64
	var numfour float32 = 32.1
	var numfive float64 = 1231231231283128.321
	numsix := 12371829738912789312.21

	numfour = 32

	fmt.Println(numone, numtwo, numthree, numfour, numfive, numsix)
}
