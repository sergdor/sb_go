package main

import (
	"fmt"
)

func main() {
	fmt.Println("========================")
	fmt.Println("Задача 21.01.")
	fmt.Println("========================")
	task2101()
	fmt.Println("========================")
	fmt.Println("Задача 21.02.")
	fmt.Println("========================")
	fmt.Println(task2102())
}

func task2101() {
	//Задание 1. Расчёт по формуле
	//Что нужно сделать
	//Напишите функцию, производящую следующие вычисления.
	//S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.
	//Тип S должен быть во float32.
	//Советы и рекомендации
	//Обратите внимание, к какому типу надо привести конечный результат.
	//В качестве среды разработки может помочь Goland или VScode.

	f := func(x int16, y uint8, z float32) float32 { return float32(2)*float32(x) + float32(y^2) - 3/z }
	fmt.Println(f(1, 2, 3))
}

func task2102() (res int) {
	//Задание 2. Анонимные функции
	//Что нужно сделать
	//Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает и вызывает её при выходе (через defer).
	//Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное, чтобы все три выполняли разное действие.
	//Рекомендация
	//В качестве среды разработки может помочь Goland или VScode.

	defer func() {
		res = task2102_1(task2102_4)
	}()
	task2102_1(task2102_2)
	task2102_1(task2102_3)
	return
}

func task2102_1(A func(int, int) int) int {
	res := A(1, 2)
	fmt.Println(res)
	return res
}
func task2102_2(a int, b int) int {
	return a + b + 1
}
func task2102_3(a int, b int) int {
	return a + b + 2
}
func task2102_4(a int, b int) int {
	return a + b + 3
}
