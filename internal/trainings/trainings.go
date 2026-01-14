// trainings пакет для работы с данными о тренировке.
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

// Training содержит данные о тренировке
type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

var (
	ErrValueLessZero       error = errors.New("should be more than 0") // значение элемента не может быть меньше либо равно нулю
	ErrUnknownTrainingType error = errors.New("неизвестный тип тренировки")
)

// Parse парсит строку с данными о тренировке, валидирует полученные данные.
func (t *Training) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if parsedCount := len(splitData); parsedCount != 3 {
		return fmt.Errorf("wrong parsedCount, got: %d wanted: %d", parsedCount, 3)
	}

	stepsCount, err := strconv.Atoi(splitData[0])
	if err != nil {
		return fmt.Errorf("did not parse stepsCount: %w", err)
	}

	if stepsCount <= 0 {
		return fmt.Errorf("stepsCount got: %d, %w", stepsCount, ErrValueLessZero)
	}

	activityTime, err := time.ParseDuration(splitData[2])
	if err != nil {
		return fmt.Errorf("did not parse activityTime: %w", err)
	}
	if activityTime <= 0 {
		return fmt.Errorf("wrong value of activityTime: %w", ErrValueLessZero)
	}

	t.Steps = stepsCount
	t.TrainingType = splitData[1]
	t.Duration = activityTime
	return nil
}

// ActionInfo формирует и возвращает строку с данными о тренировке, исходя из того, какой тип тренировки был передан.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var spentCalories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("TrainingInfo: %w", ErrUnknownTrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("TrainingInfo: %w", err)
	}

	return activityInfoText(t.TrainingType, t.Duration.Hours(), distance, avgSpeed, spentCalories), nil
}

// activityInfoText формирует текст для вывода пользователю информации о тренировке.
func activityInfoText(activity string, activityTimeHours, distanceKm, avgSpeed, spentCalories float64) string {
	output := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, activityTimeHours, distanceKm, avgSpeed, spentCalories)
	return output
}
