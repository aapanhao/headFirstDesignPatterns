package main

import (
	"log"
	"sync"
)

var (
	s    *Singleton
	once sync.Once
)

type Singleton struct{}

func (s *Singleton) GetName() {
	log.Println("Singleton Name")
}

func NewSingleton() *Singleton {
	once.Do(func() {
		s = &Singleton{}
	})
	return s
}
