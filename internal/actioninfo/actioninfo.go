package actioninfo

import "fmt"

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, value := range dataset {
		if err := dp.Parse(value); err != nil {
			fmt.Println("ошибка парсинга:", err)
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
