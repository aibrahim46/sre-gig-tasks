package main

import "fmt"

func main() {
	var Menu []string
	Menu = append(Menu, "Hamburger")
	Menu = append(Menu, "salad")
	fmt.Println("First Item in Menu:", Menu[0])
	fmt.Println("Second Item in Menu:", Menu[1])

	// an array of 5 items
	var Items = [5]string{"Pencil", "Marker", "Dry eraser", "Board", "Projector"}

	for index, item := range Items {
		fmt.Printf("This is %v and its index in the array is %v \n", item, index)
	}

}
