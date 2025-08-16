// Package daysteps предоставляет функциональность для обработки данных о шагах,
// расчета пройденной дистанции и потраченных калорий.
package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps содержит информацию о дневной активности пользователя
type DaySteps struct {
	Steps                 int           // Количество шагов.
	Duration              time.Duration // Длительность прогулки.
	personaldata.Personal               // встроенная структура Personal из пакета personaldata
}

// Ошибки парсинга:
var (
	// ErrInvalidFormat возвращается при неверном формате входных данных
	ErrInvalidFormat = errors.New("input should be in format 'count,duration'")
	// ErrInvalidCount возвращается при невалидном количестве шагов
	ErrInvalidCount = errors.New("count must be a positive integer")
	// ErrInvalidDuration возвращается при невалидной продолжительности
	ErrInvalidDuration = errors.New("duration must be a valid time.Duration (e.g., '3h50m', '0h50m')")
)

// Parse парсит строку с данными о шагах и продолжительности активности.
//
// Параметры:
//   - dataString: строка формата "количество_шагов,продолжительность" (например "1000,1h30m")
//
// Возвращаемые значения:
//   - error: ошибка парсинга или nil при успехе
func (ds *DaySteps) Parse(dataString string) (err error) {
	parsedString := strings.Split(dataString, ",")

	if len(parsedString) != 2 {
		return fmt.Errorf("%w: got %d parts", ErrInvalidFormat, len(parsedString))
	}

	parsedSteps, err := strconv.Atoi(parsedString[0])
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidCount, err)
	}
	if parsedSteps <= 0 {
		return fmt.Errorf("%w: got %d", ErrInvalidCount, parsedSteps)
	}

	parsedDuration, err := time.ParseDuration(parsedString[1])
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidDuration, err)
	}
	if parsedDuration.Seconds() <= 0 {
		return fmt.Errorf("%w: %v", ErrInvalidDuration, err)
	}

	ds.Steps = parsedSteps
	ds.Duration = parsedDuration

	return nil
}

// ActionInfo возвращает форматированный отчет о дневной активности
//
// Возвращаемые значения:
//   - string: форматированная строка с результатами активности
//   - error: ошибка при расчете показателей или nil при успехе
//
// Пример результата:
// Количество шагов: 792.
// Дистанция составила 0.51 км.
// Вы сожгли 221.33 ккал.
func (ds *DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
