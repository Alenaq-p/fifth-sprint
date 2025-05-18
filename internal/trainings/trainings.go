package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	activityRun  = "Бег"    // Активность типа бег
	activityWalk = "Ходьба" // Активность типа ходьба
)

type Training struct {
	Steps                 int           //количество шагов, проделанных за тренировку.
	TrainingType          string        //тип тренировки(бег или ходьба).
	Duration              time.Duration // длительность тренировки
	personaldata.Personal               //   встроенная структура Personal. имя вес и рост
}

func (t *Training) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",") //Разделение строки на слайс строк
	if len(parts) != 3 {                    //проверка длины слайса
		return fmt.Errorf("invalid format")
	}

	// первая часть. количество шагов
	steps, err := strconv.Atoi(parts[0]) // преобразование в int
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("the number of steps must be greater than 0")
	}
	t.Steps = steps

	// 2 часть. активность
	t.TrainingType = parts[1]

	// 3 часть строки. продолжительность прогулки
	duration, err := time.ParseDuration(parts[2]) // строка в time.Duration
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("the duration of the walk should be more than 0")
	}
	t.Duration = duration

	return nil

}

func (t Training) ActionInfo() (string, error) {

	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	calories := 0.0
	var err error
	switch t.TrainingType {
	case activityWalk:
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case activityRun:
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("unknown training type")
	}
	if err != nil {
		return "", fmt.Errorf("data error: %w", err)
	}

	sample := "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n"
	line := fmt.Sprintf(sample, t.TrainingType, t.Duration.Hours(), distance, speed, calories)

	return line, nil

}
