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
	Steps                 int           // количество шагов
	Duration              time.Duration // длительность прогулки
	personaldata.Personal               // встроенная структура Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {

	parts := strings.Split(datastring, ",") // Разделение строки на слайс строк
	if len(parts) != 2 {                    // проверка длины слайса
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
	ds.Steps = steps

	// вторая часть строки. продолжительность прогулки
	duration, err := time.ParseDuration(parts[1]) // строка в time.Duration
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("the duration of the walk should be more than 0")
	}
	ds.Duration = duration

	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("error in calorie calculation: %w", err)
	}

	sample := "Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n"
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	line := fmt.Sprintf(sample, ds.Steps, distance, calories)

	return line, nil

}
