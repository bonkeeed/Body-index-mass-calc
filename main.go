package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор ИМТ __  ")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(IMT float64) {
	hood := IMT
	var r string
	if hood < 16 {
		r = "Выраженный дефицит"
	} else if hood > 16 && hood < 18 {
		r = "Недостаточная масса"
	} else if hood > 18 && hood < 25 {
		r = "Норма"
	} else if hood > 25 && hood < 30 {
		r = "Избыточная (предожирение)"
	} else if hood > 30 && hood < 35 {
		r = "Ожирение I степени"
	} else if hood > 35 && hood < 40 {
		r = "Ожирение II степени"
	} else if hood > 40 {
		r = "Ожирение III степени"
	}
	result := fmt.Sprintf("Ваш имт: %.0f", IMT)
	fmt.Println(result, r)
}

func calculateIMT(userKG float64, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKG / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
