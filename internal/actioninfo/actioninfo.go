// Package actioninfo предоставляет унифицированный интерфейс для обработки и вывода информации
// о различных типах физической активности (тренировках, дневной активности).
package actioninfo

import "log"

// DataParser определяет интерфейс для парсинга данных активности и генерации отчетов
type DataParser interface {
	// Parse разбирает строку с данными активности
	Parse(dataString string) error

	// ActionInfo генерирует отчет на основе распарсенных данных
	ActionInfo() (string, error)
}

// Info обрабатывает набор данных активности и выводит результаты
//
// Параметры:
//   - dataset: слайс строк с данными активностей в формате "параметры1,параметры2,..."
//   - dp: экземпляр типа, реализующего DataParser (Training или DaySteps)
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}

		if info, err := dp.ActionInfo(); err != nil {
			log.Println(err)
		} else {
			log.Println(info)
		}
	}
}
