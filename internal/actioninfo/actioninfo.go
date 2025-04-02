package actioninfo

//package main

import (
	"fmt"
	//"time"
	//"github.com/AeddGynvael3110/sprint5test/internal/personaldata"
	//"github.com/AeddGynvael3110/sprint5test/internal/trainings"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, value := range dataset {
		if err := dp.Parse(value); err != nil {
			fmt.Println("Ошибка парсинга:", err)
			continue
		}
		result, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("ошибка получения информации об активности:", err)
			continue
		}
		fmt.Println(result)
	}
}

//func main() {
//	dataset := []string{
//		"2555,Ходьба,2h30m",
//		"биба",
//		"",
//		"-300,Бег,2h30m",
//		"биба,Ходьба,2h30m",
//		"4000,Ходьба,400",
//	}
//
//	p := personaldata.Personal{Name: "Anna", Weight: 80, Height: 180}
//
//	training := trainings.Training{Steps: 8765, TrainingType: "Бег", Duration: time.Duration(90) * time.Minute, Personal: p}
//
//	Info(dataset, &training)
//}
