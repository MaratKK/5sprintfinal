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
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	result, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return result * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		err := errors.New("incorrect value of steps")
		return 0, err
	}
	if weight <= 0 {
		err := errors.New("incorrect value of weight")
		return 0, err
	}
	if height <= 0 {
		err := errors.New("incorrect value of height")
		return 0, err
	}
	if duration <= 0 {
		err := errors.New("incorrect value of duration")
		return 0, err
	}
	return MeanSpeed(steps, height, duration) * weight * duration.Minutes() / float64(minInH), nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps < 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (height * stepLengthCoefficient * float64(steps)) / float64(mInKm)
}
