package main

import "fmt"

func main(){
	/*
	* Penggunaan Single Variabel
	*/
	var firstName string ="Teguh"
	var middleName string
	middleName="Muhammad"
	//* deklarasi variabel dengan metode **type inference
	//** tipe data tidak ditulis pada saat deklarasi
	lastName:="Harits"

	/*
	* Penggunaan multi variable
	*/

	//*1 Menuliskan Variabel dengan pembatas koma
	var first,second,third string
	first,second,third="satu","dua","tiga"
	//*2 pengisian nilai variabel pada saat inisiasi
	var fourth,fifth,sixth string="empat","lima","enam"
	//*3 kalau ingin lebih ringkas
	seventh,eight,ninth:="tujuh","delepan","sembilan"

	//*Variabel under score _
	name,_:="Muhammad","Aris"

	//* Deklarasi variabel dengan keyword new
	myName:= new(string)
	fmt.Println(myName)
	fmt.Println(*myName)

	fmt.Printf("hallo %s %s %s !\n",firstName,middleName,lastName)
	fmt.Printf("Angka 1-3 %s %s %s !\n",first,second,third)
	fmt.Printf("Angka 4-6 %s %s %s !\n",fourth,fifth,sixth)
	fmt.Printf("Angka 7-9 %s %s %s !\n",seventh,eight,ninth)
	fmt.Printf("Nama %s  !\n",name)
}
