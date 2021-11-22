package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10
const n2 = 12

func main() {

	fmt.Println("========================")
	fmt.Println("Задача 22.01.")
	fmt.Println("========================")
	task2201()
	fmt.Println("========================")
	fmt.Println("Задача 22.02.")
	fmt.Println("========================")
	task2202()
}

func task2201() {
	//Задание 1. Подсчёт чётных и нечётных чисел в массиве
	//Что нужно сделать
	//Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число.
	//Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого.
	//При отсутствии введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.
	//Что оценивается
	//Программа должна корректно считать числа в массиве после заданного.

	rand.Seed(time.Now().UnixNano())

	var m [n]int // массив значений
	for i := 0; i < n; i++ {
		m[i] = rand.Intn(10)
	}
	fmt.Println(m)
	var value int //искомое число

	fmt.Print("Введите число:")
	fmt.Scanln(&value)

	startCount := false // флаг начинать подсчет
	result := 0         // результат
	for i := 0; i < n; i++ {
		if startCount {
			result = n - i
			break
		}
		if m[i] == value {
			startCount = true
		}
	}

	fmt.Println(result)

}

func task2202() {
	//Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
	//Что нужно сделать
	//Заполните упорядоченный массив из 12 элементов и введите число.
	//Необходимо реализовать поиск первого вхождения заданного числа в массив.
	//Сложность алгоритма должна быть минимальная.
	//Что оценивается
	//Верность индекса.
	//При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1.

	rand.Seed(time.Now().UnixNano())

	var m [n2]int //массив чисел
	for i := 0; i < n2; i++ {
		m[i] = i*10 + rand.Intn(10)
	}
	fmt.Println(m)

	var value int //искомое число
	fmt.Print("Введите число:")
	fmt.Scanln(&value)

	fmt.Println(findIndexInMass(m, value))

}

func findIndexInMass(m [n2]int, desiredValue int) (result int) {
	result = -1   // результат
	min := 0      // минимальное значение индекса
	max := n2 - 1 // максимальное значение индекса
	for max >= min {
		middle := (max + min) / 2 // средний индекс
		if m[middle] == desiredValue {
			result = middle
			break
		} else if m[middle] > desiredValue {
			max = middle - 1
		} else {
			min = middle + 1
		}
		for result > 0 && m[result] == desiredValue {
			result--
		}
	}
	return
}
