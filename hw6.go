package main

import (
	"fmt"
)

//Домашнее задание 6
func main() {

	fmt.Println("========================")
	fmt.Println("Задача 06.01.")
	fmt.Println("========================")
	task0601()
	fmt.Println("========================")
	fmt.Println("Задача 06.02.")
	fmt.Println("========================")
	task0602()
	fmt.Println("========================")
	fmt.Println("Задача 06.03.")
	fmt.Println("========================")
	task0603()
	fmt.Println("========================")
	fmt.Println("Задача 06.04.")
	fmt.Println("========================")
	task0604()
	fmt.Println("========================")
	fmt.Println("Задача 06.05.")
	fmt.Println("========================")
	task0605()
	fmt.Println("========================")
	fmt.Println("Задача 06.06.")
	fmt.Println("========================")
	task0606()

}
func task0601() {
	/*Задание 1. Написание простого цикла
	Напишите программу, которая будет принимать от пользователя произвольное число и в цикле выводить на экран все значения от нуля до указанного числа.
	*/
	fmt.Println("Написание простого цикла")
	var val int
	fmt.Print("Введите число:")
	fmt.Scan(&val)
	for i := 0; i <= val; i++ {
		fmt.Println(i)
	}
}
func task0602() {
	/*Задание 2. Сумма двух чисел по единице
	Напишите программу, которая запрашивает у пользователя два числа и складывает их в цикле
	следующим образом: берёт первое число и прибавляет к нему по единице, пока не достигнет суммы двух чисел.*/
	fmt.Println("Сумма двух чисел по единице")
	var val1, val2 int
	fmt.Print("Введите число 1:")
	fmt.Scan(&val1)
	fmt.Print("Введите число 2:")
	fmt.Scan(&val2)
	for i := val1; i <= val1+val2; i++ {
		fmt.Println(i)
	}

}
func task0603() {
	/*Задание 3. Расчёт суммы скидки
	Напишите программу, которая принимает на вход цену товара и скидку. Посчитайте и верните на экран сумму скидки.
	Дополнительное условие: скидка не превышает 30% от цены товара и не больше 2 000 рублей.*/
	fmt.Println("Расчёт суммы скидки")
	var price int         // цена товара
	var discount int      // процент скидки
	var discountMoney int // скидка в рублях
	fmt.Print("Введите цену:")
	fmt.Scan(&price)
	fmt.Print("Введите процент скидки:")
	fmt.Scan(&discount)
	if discount < 30 {
		discountMoney = price * discount / 100
		if discountMoney <= 2000 {
			fmt.Print("Сумма скидки:", discountMoney, "\n")
		} else {
			fmt.Println("Сумма скидки превышает 2000 рублей")
		}
	} else {
		fmt.Println("Размер скидки превышает 30%")
	}

}
func task0604() {
	/*Задание 4. Предыдущий урок на if
	Есть три переменные со значениями 0. Первую нужно наполнить до десяти, вторую до 100 и третью до 1 000.
	Напишите цикл, в котором эти переменные будут увеличиваться на один.
	Используйте условия для пропуска переменных, которые уже достигли своих лимитов.*/
	fmt.Println("Предыдущий урок на if")
	var val1, val2, val3 int
	for val3 = 0; val3 < 1000; val3++ {
		if val2 < 100 {
			val2++
			if val1 < 10 {
				val1++
			}
		}
	}
	fmt.Println(val1, val2, val3)

}
func task0605() {
	/*Задание 5*. Задача на постепенное наполнение корзин разной ёмкости (if и continue и break)
	Представьте, что у вас есть три корзины разной ёмкости. Пользователю предлагается ввести, какое количество яблок помещается в каждую корзину.
	После этого программа должна заполнить все корзины по очереди, учитывая, какие корзины уже заполнены,
	строго соблюдая очерёдность заполнения и добавляя по одному яблоку в каждой итерации.
	Пример: пользователь решил, что у корзин будет ёмкость на шесть, четыре и девять яблок.
	Программа должна заполнить корзину 1 и в следующей итерации перейти к корзине 2, далее в следующей итерации перейти к корзине 3.
	Если очередная корзина уже заполнена, программа должна переходить к следующей по очереди, и так по кругу, пока не заполнит все.
	*/
	fmt.Println("Задача на постепенное наполнение корзин разной ёмкости (if и continue и break)")
	var basketCapacity1, basketCapacity2, basketCapacity3 int // вместимость корзин
	var basketAmount1, basketAmount2, basketAmount3 int       // количество яблок в корзине
	var applesAmount int                                      // количество яблок
	fmt.Print("Введите вместимость корзины 1:")
	fmt.Scan(&basketCapacity1)
	fmt.Print("Введите вместимость корзины 2:")
	fmt.Scan(&basketCapacity2)
	fmt.Print("Введите вместимость корзины 3:")
	fmt.Scan(&basketCapacity3)
	fmt.Print("Введите количество яблок:")
	fmt.Scan(&applesAmount)

	for applesAmount > 0 && basketAmount1 < basketCapacity1 {
		applesAmount--
		basketAmount1++
	}
	for applesAmount > 0 && basketAmount2 < basketCapacity2 {
		applesAmount--
		basketAmount2++
	}
	for applesAmount > 0 && basketAmount3 < basketCapacity3 {
		applesAmount--
		basketAmount3++
	}
	fmt.Print("Яблок в корзине 1:", basketAmount1, "\n")
	fmt.Print("Яблок в корзине 2:", basketAmount2, "\n")
	fmt.Print("Яблок в корзине 3:", basketAmount3, "\n")
	fmt.Print("Яблок осталось:", applesAmount, "\n")
	//сокрее всего не верно понял задачу т.к. не используются continue и break
}
func task0606() {
	/*Задание 6*. Движение лифта
	В доме 24 этажа. Лифт должен ходить по кругу, пока не доставит всех пассажиров на первый этаж.
	Три пассажира ждут на четвёртом, седьмом и десятом этажах. При движении вверх лифт не должен останавливаться, при движении вниз — собирать всех,
	но не более двух человек в лифте. При этом лифт каждый раз доезжает до самого верхнего этажа и только после этого начинает движение вниз.
	Напишите программу, которая доставит всех пассажиров на первый этаж.
	*/
	fmt.Println("Движение лифта")

	var floorMax int = 24        // максимальный этаж
	var floorMin int = 1         // минимальный этаж
	var floorPassenger1 int = 4  // этаж пассажира 1
	var floorPassenger2 int = 7  // этаж пассажира 2
	var floorPassenger3 int = 10 // этаж пассажира 3
	var currentFloor int = 1     // текущий этаж лифта
	var direction int = -1       // направление движения лифта
	var amountPeople int         // количество людей в лифте

	for floorPassenger1 > 1 || floorPassenger2 > 1 || floorPassenger3 > 1 || amountPeople > 0 {
		fmt.Println("Этаж:", currentFloor, "пассажиров в лифте:", amountPeople)
		if currentFloor == floorMax {
			direction *= -1
			fmt.Println("Этаж:", currentFloor, "меняем направление движения", direction)
		} else if currentFloor == floorMin {
			direction *= -1
			amountPeople = 0
			fmt.Println("Этаж:", currentFloor, "меняем направление движения", direction)
		}

		currentFloor += direction

		if direction == 1 {
			continue
		}

		if amountPeople < 2 && currentFloor != 1 {
			if currentFloor == floorPassenger1 {
				fmt.Println("Этаж:", currentFloor, "забираем пассажира")
				amountPeople++
				floorPassenger1 = 1
			} else if currentFloor == floorPassenger2 {
				fmt.Println("Этаж:", currentFloor, "забираем пассажира")
				amountPeople++
				floorPassenger2 = 1
			} else if currentFloor == floorPassenger3 {
				fmt.Println("Этаж:", currentFloor, "забираем пассажира")
				amountPeople++
				floorPassenger3 = 1
			}
		}
	}
}
