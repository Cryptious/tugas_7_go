package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	bacatipe1(10)
	go bacatipe2("Pesan")

	var input string
	fmt.Scanln(&input)

}

func bacatipe1(x int) {
	var reflectnumber = reflect.ValueOf(x)
	fmt.Println("Tipe Variabel :", reflectnumber.Type())

	if reflectnumber.Kind() == reflect.Int {
		fmt.Println("Nilai Variabel :", reflectnumber.Int())
	}
}

func bacatipe2(y string) {
	var reflectpesan = reflect.ValueOf(y)
	fmt.Println("Tipe Variabel :", reflectpesan.Type())

	if reflectpesan.Kind() == reflect.String {
		fmt.Println("Nilai Variabel :", reflectpesan.String())
	}
}
