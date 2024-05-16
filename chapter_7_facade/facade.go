package main

type Facade struct {
	Dvd   Dvd
	Cd    Cd
	Light Light
}

func (f *Facade) WatchMovie() {
	f.Dvd.On()
	f.Cd.On()
	f.Light.Open()
	f.Light.SetLight(11)
}

func (f *Facade) EndMovie() {
	f.Dvd.Off()
	f.Cd.Off()
	f.Light.Close()
}

func NewMovieFacade(d Dvd, c Cd, l Light) *Facade {
	return &Facade{d, c, l}
}
