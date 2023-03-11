package main

import "fmt"

func main() {
	var name string

	fmt.Print("Masukan nama anda : ")
	fmt.Scanln(&name)

	fmt.Printf("Your name is %s", name)
}