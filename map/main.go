package main

import "fmt"

func main() {
	// myMap := make(map[string]string)
	// myMap["dog"] = "Moti"
	// myMap["other-dog"] = "Taggya"
	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	// myMap := make(map[string]int)
	// myMap["first"] = 1
	// myMap["second"] = 2
	// log.Println(myMap["first"], myMap["second"])

	// myVar := helpers.SomeType{
	// 	TypeName:   "ajit",
	// 	TypeNumber: 3,
	// }

	// log.Println(myVar.TypeName, myVar.TypeNumber)

	statePopulation := map[string]int{
		"California":    39250017,
		"Texas":         27862596,
		"Florida":       2062439,
		"New York":      19745289,
		"Pennysylvania": 12802503,
		"Illinois":      12801539,
		"ohio":          11614373,
	}
	fmt.Println(statePopulation)

}
