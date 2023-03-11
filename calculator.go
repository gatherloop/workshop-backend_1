package main

import "fmt"

func main() {
    var operator string
    var num1, num2 float32

    fmt.Print("Masukkan operator (+, -, *, /): ")
    fmt.Scanln(&operator)

    fmt.Print("Masukkan angka pertama: ")
    fmt.Scanln(&num1)

    fmt.Print("Masukkan angka kedua: ")
    fmt.Scanln(&num2)

    var result float32

    if operator == "+" {
        result = num1 + num2
    } else if operator == "-" {
        result = num1 - num2
    } else if operator == "*" {
        result = num1 * num2
    } else if operator == "/" {
        result = num1 / num2
    } else {
        fmt.Println("Operator yang dimasukkan tidak valid")
        return
    }

    fmt.Printf("%.1f %s %.1f = %.1f", num1, operator, num2, result)
}