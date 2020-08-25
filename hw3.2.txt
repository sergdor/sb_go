package main

import "fmt"

func main() {
  Task2()
}
func Task2() {
//Задача 2. Симулятор маршрутки с циклами
  fmt.Println("Симулятор маршрутки")
    
  var busStopCount int  //Количество остановок
  var busStop [] string //Массив остановок 
  
  fmt.Print("Введите количество остановок:")
  fmt.Scan(&busStopCount)

  for busStopNumber := 0; busStopNumber < busStopCount; busStopNumber++ {
    busStopName := "" 
    fmt.Print("Введите название остановки ", busStopNumber + 1, ":")
    fmt.Scanln(&busStopName)
    if busStopName == "" {
      fmt.Println("Введено не корректное название остановки!")
      busStopNumber --
    } else {
      busStop = append(busStop, busStopName);
    }
  }

  var passengersOut int //Пассажиров вошло
  var passengersIn int  //Пассажиров вышло
  passengersTotal := 0  //Всего пассажиров в маршрутке
  allPassengersIn :=0   //Все вошедшие пассажиры

  for busStopNumber := 0; busStopNumber < busStopCount; busStopNumber++ {
    
    fmt.Println("\n", "-----------Едем---------", "\n")
    fmt.Print("Прибываем на остановку ", busStop [busStopNumber], ". В салоне пассажиров:", passengersTotal, "\n")
    
    if passengersTotal > 0 {
      fmt.Print("Сколько пассажиров вышло на остановке?:")
      fmt.Scan(&passengersOut)
      passengersTotal -=passengersOut
    }
    
    fmt.Print("Сколько пассажиров зашло на остановке?:")
    fmt.Scan(&passengersIn)
    
    passengersTotal += passengersIn
    allPassengersIn += passengersIn

    fmt.Print("Отправляемся с остановки ", busStop [busStopNumber], ". В салоне пассажиров:", passengersTotal, "\n")

  }
  
  allMoney  := allPassengersIn * 20 //Всего денег
  
  salary    := allMoney / 4 //Зарплата водителя
  fuel      := allMoney / 5 //Расходы на топливо
  taxes     := allMoney / 5 //Налоги
  repair    := allMoney / 5 //Расходы на ремонт машины
  
  fmt.Println("\n", "==========================", "\n")
  fmt.Println("\n", "Всего заработали:", allMoney ,"руб.")
  fmt.Println("\n", "Зарплата водителя:", salary ,"руб.")
  fmt.Println("\n", "Расходы на топливо:", fuel ,"руб.")
  fmt.Println("\n", "Налоги:", taxes ,"руб.")
  fmt.Println("\n", "Расходы на ремонт машины:", repair ,"руб.")
  fmt.Println("\n", "Итого доход:", allMoney - salary - fuel - taxes - repair ,"руб.")

}
