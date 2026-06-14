package main

import (
	"fmt"
	"math"
)

const (
	pricePerKm = 3.0
)

type rateType float64

const (
	standart rateType = iota
	high
	extreme
)

type FlightData struct {
	coords_A []uint
	coords_B []uint
	products []string
	rate     rateType
}

var price float64
var kf = 1.0
var assortment = map[string]int{"Помидоры": 350,
	"Огурцы": 200, "Макароны": 79, "Чеснок": 69, "Мармелад": 150, "Конфеты": 350}

func CalculatePrice(flight FlightData, distance float64) (float64, error) {
	price = 0.0
	switch flight.rate {
	case 1:
		kf = 2.0
	case 2:
		kf = 3.5
	}

	if distance <= 3 {
		kf += 0.25
	} else if distance <= 10 {
		kf += 1.25
	} else if distance <= 20 {
		kf += 3.75
	} else if distance <= 50 {
		kf += 7.25
	} else if distance > 50 {
		return 0, fmt.Errorf("Ошибка расстояния!")
	}

	for _, i := range flight.products {
		for j, p := range assortment {
			if i == j {
				price += float64(p)
			}
		}
	}
	fmt.Println(kf)
	return price + distance*pricePerKm*kf, nil
}

func main() {
	var ax, ay uint
	var coords_A = []uint{ax, ay}
	fmt.Scanln(&ax, &ay)

	var bx, by uint
	var coords_B = []uint{bx, by}
	fmt.Scanln(&bx, &by)

	var n int
	fmt.Scanln(&n)
	products := make([]string, n)
	for i := range products {
		fmt.Scanln(&products[i])
	}

	distance := math.Sqrt(float64((ax-bx)*(ax-bx) + (ay-by)*(ay-by)))
	answer, _ := CalculatePrice(FlightData{coords_A, coords_B, products, extreme}, distance)
	fmt.Printf("Цена: %.0f", answer)
}
