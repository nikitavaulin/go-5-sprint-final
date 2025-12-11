// actioninfo реализует вывод общей информации обо всех тренировках и прогулках
package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(dataString string) error // Парсит строку данных и заполняет структуру активности данными
	ActionInfo() (string, error)   // Формирует вывод информации об активности
}

// Info выводит информацию о виде аткивности
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("error: %s", err)
		}
		output, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error: %s", err)
		}
		fmt.Println(output)
	}
}
