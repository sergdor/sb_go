package main

import (
	"fmt"
)

func main() {

	fmt.Println("========================")
	fmt.Println("Задача 15.01.")
	fmt.Println("========================")
	task1501()
	fmt.Println("========================")
	fmt.Println("Задача 15.02.")
	fmt.Println("========================")
	task1502()
}

func task1501() {
	//Задание 1. Подсчёт чётных и нечётных чисел в массиве
	//Что нужно сделать
	//Разработайте программу, позволяющую ввести 10 целых чисел, а затем вывести из них количество чётных и нечётных чисел.
	//Для ввода и подсчёта используйте разные циклы.
	//Что оценивается
	//Для введённых чисел 1, 1, 1, 2, 2, 2, 3, 3, 3, 4 программа должна вывести: чётных — 4, нечётных — 6.
	const masCount = 10
	mas := [masCount]int{}
	evenNumber := 0
	for i, _ := range mas {

		fmt.Printf("Введите значение %v:", i+1)
		fmt.Scanln(&mas[i])
		if mas[i]%2 == 0 {
			evenNumber++
		}
	}

	fmt.Println("чётных — ", evenNumber)
	fmt.Println("нечётных — ", masCount-evenNumber)
}
func task1502() {
	//Задание 2. Функция, реверсирующая массив
	//Что нужно сделать
	//Напишите функцию, принимающую на вход массив и возвращающую массив,
	//в котором элементы идут в обратном порядке по сравнению с исходным.
	//Напишите программу, демонстрирующую работу этого метода.
	//Что оценивается
	//При вводе 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 программа должна выводить при помощи дополнительной функции,
	//реверсировав массив: 10, 9, 8, 7, 6, 5, 4, 3, 2, 1.

	mas := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(masReverse(mas))
}
func masReverse(mas [10]int) [10]int {
	res := mas
	resLen := len(res) - 1
	for i, _ := range mas {
		res[resLen-i] = mas[i]
	}
	return res
}