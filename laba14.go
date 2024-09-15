package main

import "fmt"

func main() {
 // Объявление двух целых чисел
 x1 := 10
 x2 := 5

 // Выполнение арифметических операций
 sum := x1 + x2
 difference := x1 - x2
 product := x1 * x2
 quotient := x1 / x2
 remainder := x1 % x2

 // Вывод результатов на экран
 fmt.Println("x1:", x1)
 fmt.Println("x2:", x2)
 fmt.Println("+:", sum)
 fmt.Println("-:", difference)
 fmt.Println("*:", product)
 fmt.Println("/:", quotient)
 fmt.Println("%:", remainder)
}