package main

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type WeatherData struct {
	Observers   []Observer
	Temperature float64
	Humidity    float64
	Pressure    float64
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.Observers = append(w.Observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.Observers {
		observer.Update(w.Temperature, w.Humidity, w.Pressure)
	}
}

func (w *WeatherData) SetMeasurement(tempreture, humidity, pressure float64) {
	w.Temperature = tempreture
	w.Humidity = humidity
	w.Pressure = pressure
	w.MeasurementChange()
}

func (w *WeatherData) MeasurementChange() {
	w.NotifyObservers()
}
