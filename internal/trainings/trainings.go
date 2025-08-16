// Package trainings предоставляет функции для расчета параметров тренировки:
// пройденной дистанции, средней скорости и потраченных калорий.
package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training представляет информацию о тренировке пользователя
type Training struct {
	Steps                 int           // количество шагов, проделанных за тренировку
	TrainingType          string        // тип тренировки(бег или ходьба)
	Duration              time.Duration // длительность тренировки
	personaldata.Personal               // встроенная структура Personal из пакета personaldata
}

var (
	// ErrInvalidFormat возвращается при неверном формате входных данных
	ErrInvalidFormat = errors.New("input should be in format 'count,training_type,duration'")
	// ErrInvalidCount возвращается при невалидном количестве шагов
	ErrInvalidCount = errors.New("count must be a positive integer")
	// ErrInvalidDuration возвращается при невалидной продолжительности
	ErrInvalidDuration = errors.New("duration must be a valid time.Duration (e.g., '3h50m', '0h50m')")
	// ErrInvalidTraining возвращается при неизвестном типе тренировки
	ErrInvalidTraining = errors.New("неизвестный тип тренировки")
)

// Parse разбирает строку с данными о тренировке и заполняет структуру Training
//
// Параметры:
//   - dataString: строка формата "количество_шагов,тип_тренировки,продолжительность"
//
// Возвращаемые значения:
//   - error: ошибка если данные некорректны
func (t *Training) Parse(dataString string) (err error) {
	parsedString := strings.Split(dataString, ",")
	if len(parsedString) != 3 {
		return fmt.Errorf("%w: got %d parts", ErrInvalidFormat, len(parsedString))
	}

	parsedSteps, err := strconv.Atoi(parsedString[0])
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidCount, err)
	}
	if parsedSteps <= 0 {
		return fmt.Errorf("%w: got %d", ErrInvalidCount, parsedSteps)
	}

	parsedDuration, err := time.ParseDuration(parsedString[2])
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidDuration, err)
	}
	if parsedDuration.Seconds() <= 0 {
		return fmt.Errorf("%w: %v", ErrInvalidDuration, err)
	}

	t.Duration = parsedDuration
	t.TrainingType = parsedString[1]
	t.Steps = parsedSteps

	return nil
}

// ActionInfo возвращает форматированную строку с информацией о тренировке
//
// Возвращаемые значения:
//   - string: форматированная строка с результатами тренировки
//   - error: ошибка если данные некорректны или тип тренировки неизвестен
//
// Пример результата:
// Тип тренировки: ходьба
// Длительность: 1.50 ч.
// Дистанция: 3.25 км.
// Скорость: 2.17 км/ч
// Сожгли калорий: 215.50
func (t *Training) ActionInfo() (string, error) {
	var calories float64
	var calcErr error

	trainingDistance := spentenergy.Distance(t.Steps, t.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Ходьба":
		calories, calcErr = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		calories, calcErr = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", ErrInvalidTraining
	}

	if calcErr != nil {
		return "", calcErr
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		trainingDistance,
		avgSpeed,
		calories), nil
}
