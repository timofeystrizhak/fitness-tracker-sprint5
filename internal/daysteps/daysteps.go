package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return fmt.Errorf("неверный формат данных: %q", datastring)
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %w", err)
	}
	// Добавляем проверку шагов
	if steps <= 0 {
		return fmt.Errorf("некорректное количество шагов: %d", steps)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %w", err)
	}
	// Добавляем проверку длительности
	if duration <= 0 {
		return fmt.Errorf("некорректная длительность: %v", duration)
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}
