package main

import "log"

type Observer interface {
	Update(temperature, humidity, pressure float64)
}

type Display interface {
	Display()
}

type CurrentConditionDisplay struct {
	Temperature float64
	Humidity    float64
	Subject     Subject
}

func NewCurrentCondition(s Subject) {
	c := CurrentConditionDisplay{Subject: s}
	s.RegisterObserver(&c)
}

func (c *CurrentConditionDisplay) Update(temperature, humidity, pressure float64) {
	c.Temperature = temperature
	c.Humidity = humidity
	c.Display()
}

func (c *CurrentConditionDisplay) Display() {
	log.Println("CurrentCondition: tempreature:", c.Temperature, "humidity:", c.Humidity)
}

type StatisticDisplay struct {
	Temperatures []float64
	Subject      Subject
}

func NewStatisticDisplay(subject Subject) {
	statistic := StatisticDisplay{Subject: subject}
	subject.RegisterObserver(&statistic)
}

func (s *StatisticDisplay) Update(temperature, humidity, pressure float64) {
	s.Temperatures = append(s.Temperatures, temperature)
	s.Display()
}

func (s *StatisticDisplay) Display() {
	var sumTempreature float64
	var number int
	for _, t := range s.Temperatures {
		sumTempreature += t
		number += 1
	}

	avgTempreture := sumTempreature / float64(number)
	log.Println("StatisticDisplay: avg tempreature:", avgTempreture, "sum tempreature:", sumTempreature)
}
