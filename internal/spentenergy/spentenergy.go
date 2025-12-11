// spentcalories обрабатывает переданную информацию и рассчитывает потраченные калории
// в зависимости от вида активности — бега или ходьбы.
// возвращает информацию обо всех тренировках.
package spentenergy

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	ErrValueLessZero error = errors.New("should be more than 0") // значение элемента не может быть меньше либо равно нулю
)

// WalkingSpentCalories вычисляет количество затраченных калорий при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateParamsSpentCalories(steps, weight, height, duration); err != nil {
		return 0, fmt.Errorf("error WalkingSpentCalories: \n%w", err)
	}
	avgSpeed := MeanSpeed(steps, height, duration)
	spentCalories := weight * avgSpeed * duration.Minutes() / float64(minInH) // (weight * meanSpeed * durationInMinutes) / minInH
	return spentCalories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories вычисляет количество затраченных калорий при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateParamsSpentCalories(steps, weight, height, duration); err != nil {
		return 0, fmt.Errorf("error RunningSpentCalories: \n%w", err)
	}
	avgSpeed := MeanSpeed(steps, height, duration)
	spentCalories := weight * avgSpeed * duration.Minutes() / float64(minInH) // (weight * meanSpeed * durationInMinutes) / minInH
	return spentCalories, nil
}

// MeanSpeed вычисляет среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 {
		log.Printf("meanSpeed: value of steps, %s", ErrValueLessZero)
		return 0
	}
	if duration <= 0 {
		log.Printf("meanSpeed: value of duration, %s", ErrValueLessZero)
		return 0
	}
	distanceKm := Distance(steps, height)
	return distanceKm / duration.Hours() // average speed
}

// Distance вычисляет дистанцию в километрах на основе количества шагов и роста пользователя
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := stepLength * float64(steps)
	return distance / mInKm
}

// Валидирует параметры для вычисления затраченных калорий.
func validateParamsSpentCalories(steps int, weight, height float64, duration time.Duration) error {
	errCollection := []error{}
	if steps <= 0 {
		errCollection = append(errCollection, fmt.Errorf("steps value wrong, %w", ErrValueLessZero))
	}
	if weight <= 0 {
		errCollection = append(errCollection, fmt.Errorf("weight value wrong, %w", ErrValueLessZero))
	}
	if height <= 0 {
		errCollection = append(errCollection, fmt.Errorf("height value wrong, %w", ErrValueLessZero))
	}
	if duration <= 0 {
		errCollection = append(errCollection, fmt.Errorf("duration value wrong, %w", ErrValueLessZero))
	}
	if len(errCollection) != 0 {
		return errors.Join(errCollection...)
	}
	return nil
}
