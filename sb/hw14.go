package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	task1404_a = 10
	task1404_b = 20
	task1404_c = 30
)

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("========================")
	fmt.Println("Задача 14.01.")
	fmt.Println("========================")
	task1401()
	fmt.Println("========================")
	fmt.Println("Задача 14.02.")
	fmt.Println("========================")
	task1402()
	fmt.Println("========================")
	fmt.Println("Задача 14.03.")
	fmt.Println("========================")
	task1403()
	fmt.Println("========================")
	fmt.Println("Задача 14.04.")
	fmt.Println("========================")
	task1404()
}

func task1401() {

	//Задание 1. Функция возвращающая результат
	//Что нужно сделать
	//Напишите функцию, которая на вход получает число и возвращает true, если число четное, и false, если нечетное.
	//Советы и рекомендации
	//Программа запрашивает у пользователя или генерирует случайное число,
	//передает в функцию в качестве аргумента и выводит в консоль результат ее работы.
	var msg string
	a := rand.Intn(10)
	if evenТumber(a) {
		msg = "Число " + strconv.Itoa(a) + " четное"
	}else {
		msg = "Число " + strconv.Itoa(a) + " нечетное"
	}
	fmt.Println(msg)
}

func evenТumber(a int) bool {
	return a%2 == 0
}

func task1402()  {
	//Задание 2. Функция возвращающая несколько значений
	//Что нужно сделать
	//Напишите программу, которая с помощью функции генерирует 3 случайные точки в двумерном пространстве (две координаты),
	//а затем с помощью другой функции преобразует эти координаты по формулам: x = 2 * x + 10, y = -3 * y - 5.
	a1, b1 := getXY()
	a2, b2 := getXY()
	a3, b3 := getXY()
	fmt.Println("a1 =", a1, "b1 =",b1)
	fmt.Println("a2 =", a2, "b2 =",b2)
	fmt.Println("a3 =", a3, "b3 =",b3)
	applyFormula(&a1, &b1)
	applyFormula(&a2, &b2)
	applyFormula(&a3, &b3)
	fmt.Println("a1 =", a1, "b1 =",b1)
	fmt.Println("a2 =", a2, "b2 =",b2)
	fmt.Println("a3 =", a3, "b3 =",b3)
}
func applyFormula(x,y *int)  {
	*x = 2 * *x + 10
	*y = -3 * *y - 5
}
func getXY()  (int, int) {
	return rand.Intn(10), rand.Intn(10)
}

func task1403()  {
	//Задание 3. Именованные возвращаемые значения
	//Что нужно сделать
	//Напишите программу, которая на вход получает число, затем с помощью двух функции преобразует его.
	//Первая умножает, а вторая прибавляет число, используя именованные возвращаемые значения
	i := rand.Intn(10)
	fmt.Println(i)
	fmt.Println(task1403_1(i))
}
func task1403_1(i int) int {

	return addendum(multiplication(i))
}
func multiplication(i int) (res int) {
	res = i * 2
	return
}
func addendum(i int) (res int)  {
	res = i + 2
	return
}

func task1404()  {
	//Задание 4. Область видимости переменных
	//Что нужно сделать
	//Напишите программу, в которой будет 3 функции, попарно использующие разные глобальные переменные.
	//Функции, должны прибавлять к поданном на вход числу глобальную переменную и возвращать результат.
	//Затем вызвать по очереди три функции, передавая результат из одной в другую.
	i := rand.Intn(10)
	fmt.Println("Исходное число", i)
	fmt.Println("Значение глобальных переменных", task1404_a, task1404_b, task1404_c)

	fmt.Println(task1404_3(task1404_2(task1404_1(i))))

}
func task1404_1(i int) int{
	return i + task1404_a + task1404_b
}
func task1404_2(i int) int{
	return i + task1404_b + task1404_c
}
func task1404_3(i int) int{
	return i + task1404_c + task1404_a
}
