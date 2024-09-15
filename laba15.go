package main

import "fmt"

// Функция для вычисления суммы двух чисел с плавающей запятой
func x (a, b float64) float64 {
 return a + b
}

// Функция для вычисления разности двух чисел с плавающей запятой
func y (a, b float64) float64 {
 return a - b
}

func main() {
 // Объявление двух чисел с плавающей запятой
 num1 := 14.88
 num2 := 1.88

 // Вычисление суммы и разности с использованием функций
 sum := x (num1, num2)
 difference := y (num1, num2)

 // Вывод результатов на экран
 fmt.Println("First number:", num1)
 fmt.Println("Second number:", num2)
 fmt.Println("Sum:", sum)
 fmt.Println("Difference:", difference)
}