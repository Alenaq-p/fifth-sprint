package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if duration <= 0 {
		return 0, fmt.Errorf("the duration of the walk should be more than 0")
	}
	if steps <= 0 {
		return 0, fmt.Errorf("the number of steps must be greater than 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be greater than 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be greater than 0")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH * walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if duration <= 0 {
		return 0, fmt.Errorf("the duration of the walk should be more than 0")
	}
	if steps <= 0 {
		return 0, fmt.Errorf("the number of steps must be greater than 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be greater than 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be greater than 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {

	if duration <= 0 { // продолжительность прогулки
		return 0
	}
	distance := Distance(steps, height) // вытаскиываем дистанцию
	averageSpeed := distance / duration.Hours()
	if duration == 0 {
		return 0
	}

	return averageSpeed

}

func Distance(steps int, height float64) float64 {

	stepLength := height * float64(stepLengthCoefficient) // длина 1 шага
	distanceInM := float64(steps) * stepLength
	distanceInKm := distanceInM / mInKm

	return distanceInKm
}
