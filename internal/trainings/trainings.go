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
	active := slice[1]
	if active == "Бег" {
		t.TrainingType = "Бег"
	}
	if active == "Ходьба" {
		t.TrainingType = "Ходьба"
	} else {
		return errors.New("неизвестный тип тренировки")
	}
	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return errors.New("ошибка парсинга продолжительности тренировки")
	}
	t.Duration = duration
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentspentenergy.Distance(steps)
	duration := t.Duration
	if duration <= 0 {
		return "", errors.New("ошибка: продолжительность меньше или равна 0")
	}
	meanSpeed := spentenergy.MeanSpeed(steps, duration)
	var trType string
	if t.TrainingType == "Бег" {
		trType = "Бег"
		spentCal, err := spentenergy.RunningSpentCalories(steps, weight, duration)
		if err != nil {
			return "", errors.New("ошибка получения потраченных калорий")
		}
	}
	if t.TrainingType == "Ходьба" {
		trType = "Ходьба"
		spentCal, err := spentenergy.WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", errors.New("ошибка получения потраченных калорий")
		}
	} else {
		return "unknown training type", err
	}
	result := fmt.Printf("Тип тренировки: %s\nДлительность: %v\nДистанция: %.2f\nСкорость: %.2f\nСожгли калорий: %.2f", trType, duration, distance, meanSpeed, spentCal)
	return result, nil
}
