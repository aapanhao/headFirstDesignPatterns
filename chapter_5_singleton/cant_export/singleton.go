package main

import (
	"log"
	"sync"
)

var (
	s    *singleton
	once sync.Once
)

type singleton struct{}

func (s *singleton) GetName() {
	log.Println("singleton Name")
}

func NewSingleton() *singleton {
	once.Do(func() {
		s = &singleton{}
	})
	return s
}
