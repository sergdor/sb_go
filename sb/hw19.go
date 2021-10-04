package main

import (
	"fmt"
)

func main() {

	fmt.Println("========================")
	fmt.Println("Задача 19.01.")
	fmt.Println("========================")
	task1901()
	fmt.Println("========================")
	fmt.Println("Задача 19.02.")
	fmt.Println("========================")
	task1902()
}

func task1901() {
	//Задание 1. Слияние отсортированных массивов
	//Что нужно сделать
	//Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
	//Советы и рекомендации
	//Обратите внимание на размеры массивов.
	//В качестве среды разработки может помочь GoLand или VS Code.
	//Что оценивается
	//Правильность размеров.
	//Правильный порядок элементов в конечном массиве.

	mas1 := [...]int{1, 2, 3, 4}
	mas2 := [...]int{2, 3, 4, 5, 6}
	fmt.Println(mas1)
	fmt.Println(mas2)

	fmt.Println(combiningArrays(mas1, mas2))
}
func combiningArrays(mas1 [4]int, mas2 [5]int) (masRes [9]int) {
	lenMas1 := len(mas1)
	lenMasRes := len(masRes)
	var i1, i2 int
	for i := 0; i < lenMasRes; i++ {
		if i1 < lenMas1 && mas1[i1] < mas2[i2] {
			masRes[i] = mas1[i1]
			i1++
		} else {
			masRes[i] = mas2[i2]
			i2++
		}
	}
	return
}

func task1902() {
	//Задание 2. Сортировка пузырьком
	//Что нужно сделать
	//Отсортируйте массив длиной шесть пузырьком.
	//Советы и рекомендации
	//Принцип сортировки пузырьком можно посмотреть на «Википедии», там есть наглядная демонстрация, или на YouTube.
	//В качестве среды разработки может помочь GoLand или VS Code.
	//Что оценивается
	//Правильность результата сортировки и отсутствие ошибок.
	mas := [6]int{4, 2, 3, 6, 1, 45}
	bubbleSort(mas)

}

func bubbleSort(mas [6]int) {
	fmt.Println(mas)
	for i := 0; i < len(mas); i++ {
		for j := len(mas) - 1; j > i; j-- {
			if mas[j-1] > mas[j] {
				mas[j-1], mas[j] = mas[j], mas[j-1]
			}
		}
	}
	fmt.Println(mas)
}
