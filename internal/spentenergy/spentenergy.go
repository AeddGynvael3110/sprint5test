package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
	speed     = 1.39  // средняя скорость в м/с
)

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// чек на отрицательный вес\рост менее
	if weight < 0 || height < 0 {
		return 0, fmt.Errorf("вес или рост менее 0")
	}
	// чек на отрицательную продолжительность
	if duration < 0 {
		return 0, fmt.Errorf("продолжительность не может быть отрицательной")
	}
	meanSpeed, err := MeanSpeed(steps, duration)
	if err != nil {
		return 0, fmt.Errorf("ошибка расчета средней скорости")
	}
	durationHours := duration.Hours()
	spentCal := ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * durationHours * minInH
	return spentCal, nil
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) (float64, error) {
	if weight < 0 {
		return 0, fmt.Errorf("вес не может быть отрицательным")
	}
	if duration < 0 {
		return 0, fmt.Errorf("продолжительность не может быть отрицательной")
	}
	meanSpeed, err := MeanSpeed(steps, duration)
	if err != nil {
		return 0, fmt.Errorf("ошибка расчета средней скорости")
	}
	spentCal := ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight
	return spentCal, nil
}

// МeanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func MeanSpeed(steps int, duration time.Duration) (float64, error) {
	// чек на продолжительность < 0
	if duration < 0 {
		return 0, fmt.Errorf("продолжительность не может быть отрицательной")
	}
	// получаем дистанцию из функции
	distance := Distance(steps)
	// конвертим время в часы
	durationHours := duration.Hours()
	meanSpeed := distance / durationHours
	return meanSpeed, nil
}

// Distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Для расчета дистанции нужно шаги умножить на длину шага lenStep и разделить на mInKm
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
//
// Создайте функцию ниже
func Distance(steps int) float64 {
	distance := (float64(steps) * lenStep) / mInKm
	return distance
}
