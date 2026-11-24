package spentenergy

import (
	"fmt"
	"time"
)

const (
	stepLengthCoefficient      = 0.45 // коэффициент длины шага
	mInKm                      = 1000 // метров в километре
	minInH                     = 60   // минут в часе
	walkingCaloriesCoefficient = 0.5  // коэффициент калорий при ходьбе
)

// Distance вычисляет дистанцию в км по шагам и росту в метрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceInMeters := float64(steps) * stepLength
	return distanceInMeters / mInKm
}

// MeanSpeed вычисляет среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

// RunningSpentCalories вычисляет калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("некорректное количество шагов: %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("некорректный вес: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("некорректный рост: %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("некорректная длительность: %v", duration)
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

// WalkingSpentCalories вычисляет калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("некорректное количество шагов: %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("некорректный вес: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("некорректный рост: %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("некорректная длительность: %v", duration)
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories * walkingCaloriesCoefficient, nil
}
