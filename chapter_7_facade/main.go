package main

func main() {
	dvd := Dvd{}
	cd := Cd{}
	light := Light{}
	movieFacade := NewMovieFacade(dvd, cd, light)
	movieFacade.WatchMovie()
	movieFacade.EndMovie()
}
