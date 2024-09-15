package main

import (
 "fmt"
 "time"
)

func main() {
 // Получаем текущее время
 currentTime := time.Now()

 // Выводим текущее время и дату в формате "YYYY-MM-DD HH:MM:SS"
 fmt.Println("Текущее время и дата:", currentTime.Format ("2006-01-02 15:04:05"))
}