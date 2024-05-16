package main

func main() {
	allMenu := NewMenu("All Menus", "all menu combines")
	pancakeMenu := NewMenu("Pancake Menu", "breakfast")
	dinnerMenu := NewMenu("Dinner Menu", "lunch")
	cafeMenu := NewMenu("CAFE Menu", "dinner")
	allMenu.Add(pancakeMenu, dinnerMenu, cafeMenu)

	dinnerMenu.Add(NewMenuItem("Pizza", "very much pizza", 18, false))
	dinnerMenu.Add(NewMenuItem("Tea", "hot tea", 5, true))

	dessertMenu := NewMenu("Dessert Menu", "dinner")
	dessertMenu.Add(NewMenuItem("ice cream", "cold ice cream", 6, false))
	dinnerMenu.Add(dessertMenu)

	w := NewWaiter(allMenu)
	w.Print()
}
