package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var num_generated  = rand.Intn(100)

	if num_generated > 50{ 
		if num_generated % 2 == 0 {
			fmt.Println("It's closer to 100, and it's even")
		} else{
			fmt.Println("It's closer to 100")
		}
	} else if num_generated < 50 {

		fmt.Println("It's closer to 0")	

	} else if num_generated == 50{

		fmt.Println("It's 50!")

	} else {
		fmt.Println(num_generated)
		
	}
}
