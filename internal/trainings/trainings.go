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

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("incorrect data")
	}

	stepsStr := strings.TrimSpace(parts[0])
	kindStr := strings.TrimSpace(parts[1])
	durStr := strings.TrimSpace(parts[2])

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return errors.New("failed operation " + err.Error())
	}
	if steps <= 0 {
		return errors.New("incorrect value: steps <= 0")
	}
	t.Steps = steps
	t.TrainingType = kindStr

	duration, err := time.ParseDuration(durStr)
	if err != nil {
		return errors.New("failed operation: " + err.Error())
	}
	if duration <= 0 {
		return errors.New("incorrect value: duration <= 0")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	height := t.Personal.Height
	weight := t.Personal.Weight
	steps := t.Steps
	trainingType := t.TrainingType
	var spentCalories float64
	var err error

	speed := spentenergy.MeanSpeed(steps, height, t.Duration)
	distance := spentenergy.Distance(steps, height)

	if strings.ToLower(t.TrainingType) == "ходьба" {
		spentCalories, err = spentenergy.WalkingSpentCalories(steps, weight, height, t.Duration)
	} else if strings.ToLower(t.TrainingType) == "бег" {
		spentCalories, err = spentenergy.RunningSpentCalories(steps, weight, height, t.Duration)
	} else {
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, t.Duration.Hours(), distance, speed, spentCalories)
	return result, nil
}
