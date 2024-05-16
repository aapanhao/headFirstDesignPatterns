package main

func main() {
	w := WeatherData{}
	NewCurrentCondition(&w)
	NewStatisticDisplay(&w)

	w.SetMeasurement(10, 20, 30)
	w.SetMeasurement(30, 1.21, 6.88)
	w.SetMeasurement(7.12, 4, 9)
}
