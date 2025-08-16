// Package spentenergy предоставляет функции для расчета параметров тренировки:
// пройденной дистанции, средней скорости и потраченных калорий.
package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

var (
	// ErrInvalidParams возвращается при невалидных параметрах веса/роста
	ErrInvalidParams = errors.New("weight and height must be positive values")
)

// RunningSpentCalories рассчитывает количество потраченных калорий при беге.
//
// Параметры:
//   - steps: количество шагов
//   - weight: вес пользователя в кг
//   - height: рост пользователя в м
//   - duration: продолжительность тренировки
//
// Возвращаемые значения:
//   - float64: количество калорий
//   - error: ошибка если параметры некорректны
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration.Seconds() <= 0 {
		return 0, ErrInvalidParams
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * avgSpeed * duration.Minutes()) / minInH

	return calories, nil
}

// WalkingSpentCalories рассчитывает количество потраченных калорий при ходьбе.
//
// Параметры аналогичны RunningSpentCalories.
//
// Возвращаемые значения:
//   - float64: количество калорий (умноженное на walkingCaloriesCoefficient)
//   - error: ошибка если параметры некорректны
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return calories * walkingCaloriesCoefficient, nil
}

// Distance рассчитывает пройденную дистанцию в километрах.
//
// Параметры:
//   - steps: количество шагов
//   - height: рост пользователя в метрах
//
// Возвращаемое значение:
//   - float64: дистанция в км
func Distance(steps int, height float64) float64 {
	stepHeight := height * stepLengthCoefficient
	stepsDistance := float64(steps) * stepHeight

	return stepsDistance / mInKm
}

// MeanSpeed рассчитывает среднюю скорость в км/ч.
//
// Параметры:
//   - steps: количество шагов
//   - height: рост пользователя в метрах
//   - duration: продолжительность тренировки
//
// Возвращаемое значение:
//   - float64: скорость в км/ч
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}
