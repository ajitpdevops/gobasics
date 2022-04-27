package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// myMap := make(map[string]string)
	// myMap["dog"] = "Samson"
	// myMap["other-dog"] = "Cassie"
	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	// myMap := make(map[string]int)
	// myMap["First"] = 1
	// myMap["Second"] = 2
	// log.Println(myMap["First"], myMap["Second"])

	myMap := make(map[string]User)
	me := User{
		FirstName: "Ajit",
		LastName:  "Patil",
	}

	myMap["me"] = me
	log.Println(myMap["me"].FirstName, myMap["me"].LastName)

	// myVar := helpers.SomeType{
	// 	TypeName:   "ajit",
	// 	TypeNumber: 3,
	// }

	// log.Println(myVar.TypeName, myVar.TypeNumber)

	// statePopulation := map[string]int{
	// 	"California":    39250017,
	// 	"Texas":         27862596,
	// 	"Florida":       2062439,
	// 	"New York":      19745289,
	// 	"Pennysylvania": 12802503,
	// 	"Illinois":      12801539,
	// 	"ohio":          11614373,
	// }
	// fmt.Println(statePopulation)
	// fmt.Println(statePopulation["California"])

}
