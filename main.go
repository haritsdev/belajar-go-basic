package main

import (
	"fmt"
)

func main(){
	//* OPERATOR ARITMATIKA
	var value uint= (((2+6)%3)*4-2)/3
	fmt.Print("nilai: ",value,"!\n")

	//* OPERATOR PERBANDINGAN
	var isEqual =(value==2)
	fmt.Printf("nilai %d (%t) \n",value,isEqual)

	//* Operator Logika
	//digunakan untuk mencari benar tidaknya kombinasi tipe data bertipe bool

	var left=false
	var right =true
	var leftAndRight= left && right
	fmt.Printf("left && right \t(%t) \n",leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n",leftOrRight)

	var leftReverse= !left
	fmt.Printf("!left \t\t(%t) \n",leftReverse)

}
