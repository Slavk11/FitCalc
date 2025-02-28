package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	for {
		fmt.Println("___ Калькулятор индекса массы тела___")
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)

	switch {
	case imt < 16:
		fmt.Println(" У вас сильный недостаток веса")
	case imt < 18.5:
		fmt.Println(" У вас недостаток веса")
	case imt < 25:
		fmt.Println(" У вас нормальный вес")
	case imt < 30:
		fmt.Println(" У вас избыточный вес")
	default:
		fmt.Println(" У вас степень ожирения")
	}
}

func calculateIMT(userKG float64, userHeight float64) (float64, error) {
	if userKG <= 0 || userHeight <= 0 {
		return 0, errors.New("NO PARAMS ERROR")

	}
	const IMTPower = 2
	IMT := userKG / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите сделать еще расчёт (y/n): ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
