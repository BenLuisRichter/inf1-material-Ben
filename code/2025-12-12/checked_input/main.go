package main

import "fmt"

func GetInput() int {
	fmt.Println("Gib eine Zahl zwischen 1 und 10 ein: ")
	var n int
	fmt.Scan(&n)

	if n <= 10 && n >= 0 {
		return n
	}

	fmt.Println("Ungültige Eingabe, versuch es noch einmal.")
	return GetInput() //mach die Funktion nochmal
}

func main() {
	n := GetInput()
	fmt.Printf("%v war eine gute Zahl.", n)
}
