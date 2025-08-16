// Package personaldata предоставляет структуру для хранения и обработки персональных данных пользователя.
package personaldata

import "fmt"

// Personal содержит основные данные пользователя.
type Personal struct {
	Name   string  // имя пользователя
	Weight float64 // вес пользователя
	Height float64 // рост пользователя
}

// Print выводит форматированные данные пользователя в стандартный вывод.
//
// Пример вывода:
// Имя: Иван Иванов
// Вес: 75.50 кг.
// Рост: 1.80 м.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n"+
		"Вес: %.2f кг.\n"+
		"Рост: %.2f м.\n",
		p.Name, p.Weight, p.Height)
}
