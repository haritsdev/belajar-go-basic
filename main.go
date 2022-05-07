package main

import "fmt"

func main(){
	//* 1. TIPE DATA NUMERIK NON-DESIMAL
	var positiveNumber uint8= 89
	var negativeNumber int = -123456789
	fmt.Printf("bilangan positif: %d\n",positiveNumber)
	fmt.Printf("bilangan Negative: %d\n",negativeNumber)

	//*2 TIPE DATA NUMERIK DESIMAL
	// * ada 2 yait float32 dan float64
	var decimalNumber=2.62
	fmt.Printf("bilangan desimal : %f\n",decimalNumber)
	fmt.Printf("bilangan desimal 3 digit : %.3f\n",decimalNumber)
	fmt.Printf("bilangan desimal 2 digit: %.2f\n",decimalNumber)

	//*3 Tipe Data bool (BOOLEAN)
	var exist bool =true
	fmt.Printf("exist? %t\n",exist)

	//*4 Tipe Data String
	var message string ="Selamat Belajar golang"
	var biodata=`Nama Saya "TEGUH MUHAMMAD HARITS".
		     Salam Kenal.
		     Mari belajar "Golang"
		     .
	`
	fmt.Printf("message: %s \n",message)
	fmt.Printf("biodata: %s \n",biodata)
}
