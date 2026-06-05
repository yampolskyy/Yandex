package main

import "fmt"

// const (
// 	pricePerKm     = 10.0
// 	pricePerMinute = 2.0
// )

// type TripParameters struct {
// 	Distance float64
// 	Duration float64
// }

// func CalculateBasePrice(trip TripParameters) float64 {
// 	return trip.Distance*pricePerKm + trip.Duration*pricePerMinute
// }

type WeatherCondition int

const (
	Clear WeatherCondition = iota
	Rain
	HeavyRain
	Snow
)

type WeatherData struct {
	Condition WeatherCondition
	WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	var kf = 1.0
	if weather.WindSpeed > 15 {
		kf += 0.1
	}

	switch weather.Condition {
	case 1:
		kf += 0.125
	case 2:
		kf += 0.2
	case 3:
		kf += 0.15
	}

	return kf
}

func main() {
	fmt.Println(GetWeatherMultiplier(WeatherData{HeavyRain, 10}))
	fmt.Println(GetWeatherMultiplier(WeatherData{HeavyRain, 20}))

	fmt.Println(GetWeatherMultiplier(WeatherData{Rain, 10}))
	fmt.Println(GetWeatherMultiplier(WeatherData{Rain, 20}))

	fmt.Println(GetWeatherMultiplier(WeatherData{Clear, 10}))
	fmt.Println(GetWeatherMultiplier(WeatherData{Clear, 20}))

	fmt.Println(GetWeatherMultiplier(WeatherData{Snow, 10}))
	fmt.Println(GetWeatherMultiplier(WeatherData{Snow, 20}))
}
