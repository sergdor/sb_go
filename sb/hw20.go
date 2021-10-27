package main

import (
	"fmt"
)

const rows1 = 3
const cols1 = 3

const rows2 = 3
const cols2 = 5
const cols3 = 4

func main() {
	fmt.Println("========================")
	fmt.Println("Задача 20.01.")
	fmt.Println("========================")
	task2001()
	fmt.Println("========================")
	fmt.Println("Задача 20.02.")
	fmt.Println("========================")
	task2002()
}

func task2001() {

	//Задание 1. Подсчёт определителя
	//Что нужно сделать
	//Напишите функцию, вычисляющую определитель матрицы размером 3 × 3.
	//Советы и рекомендации
	//Алгоритм вычисления представлен в «Википедии» (https://ru.wikipedia.org/wiki/%D0%9E%D0%BF%D1%80%D0%B5%D0%B4%D0%B5%D0%BB%D0%B8%D1%82%D0%B5%D0%BB%D1%8C)
	//В качестве среды разработки может помочь Goland или VScode.

	m := [rows1][cols1]int64{
		{1, 2, 3},
		{2, 6, 5},
		{4, 8, 7},
	}
	fmt.Println(m)
	fmt.Println(determinant(m))
}

func determinant(m [rows1][cols1]int64) int64 {
	x := m[1][1]*m[2][2] - m[2][1]*m[1][2]
	y := m[1][0]*m[2][2] - m[2][0]*m[1][2]
	z := m[1][0]*m[2][1] - m[2][0]*m[1][1]
	return m[0][0]*x - m[0][1]*y + m[0][2]*z
}

func task2002() {
	//Задание 2. Умножение матриц
	//Что нужно сделать
	//Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4.
	//Советы и рекомендации
	//Алгоритм умножения матриц проиллюстрирован в «Википедии» (https://ru.wikipedia.org/wiki/%D0%A3%D0%BC%D0%BD%D0%BE%D0%B6%D0%B5%D0%BD%D0%B8%D0%B5_%D0%BC%D0%B0%D1%82%D1%80%D0%B8%D1%86)
	//В качестве среды разработки может помочь Goland или VScode.

	m1 := [rows2][cols2]int64{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}
	m2 := [cols2][cols3]int64{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(multiply(m1, m2))
}

func multiply(m1 [rows2][cols2]int64, m2 [cols2][cols3]int64) [rows2][cols3]int64 {
	var m [rows2][cols3]int64
	for i := 0; i < rows2; i++ {
		for j := 0; j < cols3; j++ {
			for k := 0; k < cols2; k++ {
				m[i][j] = m[i][j] + m1[i][k]*m2[k][j]
			}
		}
	}
	return m
}
