package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/AeddGynvael3110/sprint5test/internal/personaldata"
	"github.com/AeddGynvael3110/sprint5test/internal/spentenergy"
	"github.com/AeddGynvael3110/sprint5test/internal/trainings"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return errors.New("ошибка: неверная длина слайса")
	}
	ds.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return errors.New("ошибка парсинга числа шагов")
	}
	ds.Duration, err = time.ParseDuration(slice[1])
	if err != nil {
		return errors.New("ошибка парсинга продолжительности прогулки")
	}
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", errors.New("ошибка: продолжительность меньше или равна 0")
	}

	distance := spentenergy.Distance(ds.Steps)

	var dataTraining trainings.Training
	var spentCal float64
	var err error
	if dataTraining.TrainingType == "Бег" {
		spentCal, err = spentenergy.RunningSpentCalories(dataTraining.Steps, float64(dataTraining.Weight), dataTraining.Duration)
		if err != nil {
			return "", errors.New("ошибка получения потраченных калорий")
		}
	} else if dataTraining.TrainingType == "Ходьба" {
		spentCal, err = spentenergy.WalkingSpentCalories(dataTraining.Steps, float64(dataTraining.Weight), float64(dataTraining.Height), dataTraining.Duration)
		if err != nil {
			return "", errors.New("ошибка получения потраченных калорий")
		}
	} else {
		return "", errors.New("unknown training type")
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, distance, spentCal)
	return result, nil
}
