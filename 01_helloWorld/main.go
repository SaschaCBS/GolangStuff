package main

import "fmt"

func main() {
	// String definieren
	var hello string
	// String konkatenieren
	hello = "Hello " + "Moon!"
	//Ausgabe mittels Println
	fmt.Println("Hello World!")
	fmt.Println(hello)

	//Gleichzeiting Deklaration und Zuweiseung mit :=
	bye := "Goodbye"
	fmt.Println(bye)

	//Array für int definieren. indiziert von 0 bis 9
	var myArr [10]int
	myArr[6] = 6
	myArr[7] = 7
	res := myArr[6] * myArr[7]
	println(res)

	// Type mit zwei Attributen definieren
	type User struct {
		Id   string
		Name string
	}

	//Gleichzeiting Deklaration und Zuweiseung mit := und anschließender map an Werten
	pers := User{
		Id:   "AD42",
		Name: "Douglas Adams",
	}

	println(pers.Id + pers.Name)

	// Einfaches If-Statement
	if len(pers.Name) < 42 {
		println("Name not too long")
	} else {
		println("Name too long!")
	}

	// Switch Statement bei mehreren möglichen Verwzweigungen
	switch {
	case myArr[6] < 0:
		println("below zero")
	case myArr[6] == 0:
		println("is Zero")
	case myArr[6] > 0:
		println("above Zero")
	default:
		println("I don't know!")
	}

	i := 1
	// For Schleife
	for i < myArr[6] {
		res := i * myArr[7]
		println(res)
		i++
	}
}
