package main

import (
	"log"
	"math/rand"
)

type State interface {
	InputCoin()
	EjectCoin()
	TurnCrank()
	Dispense()
}

type NoCoinState struct {
	client *Client
}

func (s *NoCoinState) InputCoin() {
	log.Println("input a coin")
	s.client.SetState(s.client.HasCoinState)
}
func (s *NoCoinState) EjectCoin() {
	log.Println("no coin input, can't eject crank")
}
func (s *NoCoinState) TurnCrank() {
	log.Println("no coin input, can't turn crank")
}
func (s *NoCoinState) Dispense() {
	log.Println("no coin input, can't dispense")
}

type HasCoinState struct {
	client *Client
}

func (s *HasCoinState) InputCoin() {
	log.Println("alread has coin")
}
func (s *HasCoinState) EjectCoin() {
	log.Println("ok return a coin")
	s.client.SetState(s.client.NoCoinState)
}
func (s *HasCoinState) TurnCrank() {
	log.Println("ok wait a minute give gumball")
	number := rand.Intn(10)
	log.Println("random int", number)
	if number == 0 && s.client.Count > 1 {
		s.client.SetState(s.client.WinnerState)
	} else {
		s.client.SetState(s.client.SoldState)
	}

}
func (s *HasCoinState) Dispense() {
	log.Println("sold state, wait for minute Dispense")
}

type SoldState struct {
	client *Client
}

func (s *SoldState) InputCoin() {
	log.Println("sold state, wait for minute InputCoin")
}
func (s *SoldState) EjectCoin() {
	log.Println("sold state, wait for minute EjectCoin")
}
func (s *SoldState) TurnCrank() {
	log.Println("sold state, wait for minute TurnCrank")
}
func (s *SoldState) Dispense() {
	s.client.ReleaseGumball()
	if s.client.Count > 0 {
		s.client.SetState(s.client.NoCoinState)
	} else {
		s.client.SetState(s.client.SoldOutState)
	}
}

type SoldOutState struct {
	client *Client
}

func (s *SoldOutState) InputCoin() {
	log.Println("sold out state, can't input coin")
}
func (s *SoldOutState) EjectCoin() {
	log.Println("sold out state, no coin input")
}
func (s *SoldOutState) TurnCrank() {
	log.Println("sold out state, can't turn crank")
}
func (s *SoldOutState) Dispense() {
	log.Println("sold out state, no gumball")
}

type WinnerState struct {
	client *Client
}

func (s *WinnerState) InputCoin() {
	log.Println("sold state, wait for minute InputCoin")
}
func (s *WinnerState) EjectCoin() {
	log.Println("sold state, wait for minute EjectCoin")
}
func (s *WinnerState) TurnCrank() {
	log.Println("sold state, wait for minute TurnCrank")
}
func (s *WinnerState) Dispense() {
	log.Println("You're winner!!!, you get two gumball")
	s.client.ReleaseGumball()
	if s.client.Count == 0 {
		s.client.SetState(s.client.SoldOutState)
	} else {
		s.client.ReleaseGumball()
		if s.client.Count > 0 {
			s.client.SetState(s.client.NoCoinState)
		} else {
			s.client.SetState(s.client.SoldOutState)
		}
	}

}
