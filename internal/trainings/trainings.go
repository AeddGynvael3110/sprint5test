package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/AeddGynvael3110/sprint5test/internal/personaldata"
	"github.com/AeddGynvael3110/sprint5test/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return errors.New("неверная длина слайса")
	}

	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return errors.New("ошибка конвертации шагов в int")
	}
	t.Steps = steps
	if t.Steps < 0 {
		return errors.New("отрицательное число шагов")
	}

	active := slice[1]
	switch {
	case active == "Бег":
		t.TrainingType = "Бег"
	case active == "Ходьба":
		t.TrainingType = "Ходьба"
	default:
		return errors.New("неизвестный тип тренировки")
	}

	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return errors.New("ошибка парсинга продолжительности тренировки")
	}
	t.Duration = duration
	return
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)

	if t.Duration <= 0 {
		return "", errors.New("продолжительность меньше или равна 0")
	}

	meanSpeed, err := spentenergy.MeanSpeed(t.Steps, t.Duration)
	if err != nil {
		return "", errors.New("ошибка получения средней скорости")
	}

	var spentCal float64

	switch {
	case t.TrainingType == "Бег":
		spentCal, err = spentenergy.RunningSpentCalories(t.Steps, float64(t.Weight), t.Duration)
		if err != nil {
			return "", errors.New("ошибка получения потраченных калорий")
		}
	case t.TrainingType == "Ходьба":
		spentCal, err = spentenergy.WalkingSpentCalories(t.Steps, float64(t.Weight), float64(t.Height), t.Duration)
		if err != nil {
			return "", errors.New("ошибка получения потраченных калорий")
		}
	default:
		return "", errors.New("unknown training type")
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, spentCal)
	return result, nil
}
