package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("========================")
	fmt.Println("Задача 11.01.")
	fmt.Println("========================")
	task1101()
	fmt.Println("========================")
	fmt.Println("Задача 11.02.")
	fmt.Println("========================")
	task1102()
}

func task1101() {
	/*
	Задание 1. Определение количества слов начинающихся с большой буквы
	Что нужно сделать:
	В начале программы объявите строку и присвойте ей:
	Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software
	Напишите программу, которая выведет количество слов начинающихся с большой буквы.
	Попробуйте ввести другие строки, например, пустую строку.
	*/
	count:=0
	s1:="Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
	s1=" go is an open source programming language that makes it easy to build simple, reliable, and efficient Software "
	s1="           "
	s1="Проверка Раз два Три"
	fmt.Println(s1)
	s1 = s1 + " "
	for len(s1) > 0 {
		s1 = strings.TrimLeft(s1, " ")
		firstSymbol:=string([]rune(s1)[:1])
		if firstSymbol != strings.ToLower(firstSymbol) {
			count++
		}

		//fmt.Println(s1, count, firstSymbol)

		spaceIndex:=strings.Index(s1, " ")
		if spaceIndex > 0 {
			s1 = s1[spaceIndex+1:]
		}
	}
	fmt.Println("Заглавных букв:", count)
}

func task1102() {
	/*
	Задание 2. Вывод чисел содержащихся в строке
	Что нужно сделать:
	В начале программы объявите строку и присвойте ей: a10 10 20b 20 30c30 30 dd
	Напишите программу, которая выведет все части строки, разделенные пробелами, которые можно привести к целому числу.
	Попробуйте ввести другие строки, например, пустую строку.

	Советы и рекомендации:
	Для проверки является ли строка числом, можно использовать функции приведения строк к числам.
	*/
	s1:="a10 10 20b 20 30c30 30 dd"
	s1=" фыв12 12 322 23в2 "
	//s1=" "

	var isStrings, isNubers string
	s1 = strings.Trim(s1, " ")
	s1Mass:= strings.Fields(s1)
	fmt.Println(s1Mass)
	for _, i:=range s1Mass{
		_, err := strconv.Atoi(i)
		if  err == nil{
			isNubers += i + " "
		} else {
			isStrings += i + " "
		}
	}
	isStrings = strings.Trim(isStrings, " ")
	isNubers = strings.Trim(isNubers," ")
	fmt.Println("Строка символов:", isStrings)
	fmt.Println("Строка чисел:", isNubers)

}