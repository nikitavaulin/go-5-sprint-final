// daysteps реализует парсинг данных о прогулках и формирование вывода информации о них
package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps содержит данные о дневных прогулках
type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

// Parse парсит данные о прогулке, валидирует и заполняет структуру прогулки
func (ds *DaySteps) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if parsedCount := len(splitData); parsedCount != 2 {
		return fmt.Errorf("wrong parsedCount, got: %d wanted: %d", parsedCount, 2)
	}

	stepsCount, err := strconv.Atoi(splitData[0])
	if err != nil {
		return fmt.Errorf("did not parse stepsCount: %w", err)
	}

	if stepsCount <= 0 {
		return fmt.Errorf("stepsCount should be more than 0, got: %d", stepsCount)
	}

	walkTime, err := time.ParseDuration(splitData[1])
	if err != nil {
		return fmt.Errorf("did not parse walkTime: %w", err)
	}

	if walkTime <= 0 {
		return fmt.Errorf("walkTime should be more than 0, got: %d", walkTime)
	}

	ds.Duration = walkTime
	ds.Steps = stepsCount
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ActionInfo: %w", err)
	}
	return actionInfoText(ds.Steps, distance, spentCalories), nil
}

// dayActionInfoText возвращает текст вывода информации для пользователя
func actionInfoText(stepsCount int, distance, spentCalories float64) string {
	output := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		stepsCount,
		distance,
		spentCalories)
	return output
}
