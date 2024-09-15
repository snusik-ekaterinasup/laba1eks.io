package main

import "fmt"

// Функция для вычисления среднего значения трех чисел с плавающей запятой
func okboss(a, b, c float64) float64 {
 return (a + b + c) / 3
}

func main() {
 // Объявление трех чисел с плавающей запятой
 num1 := 1.488
 num2 := 1.4
 num3 := 1.48

 // Вычисление среднего значения с использованием функции
 avg := okboss(num1, num2, num3)

 // Вывод результатов на экран
 fmt.Println("First number:", num1)
 fmt.Println("Second number:", num2)
 fmt.Println("Third number:", num3)
 fmt.Println("Average:", avg)
}