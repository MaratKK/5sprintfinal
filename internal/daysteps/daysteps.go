package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("incorrect data")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("failed operation: " + err.Error())
	}
	if steps <= 0 {
		return errors.New("incorrect value: steps <= 0")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return errors.New("failed operation: " + err.Error())
	}
	if duration <= 0 {
		return errors.New("incorrect value: durarion <= 0 ")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	dist := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	cals, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cals)
	return result, nil
}
