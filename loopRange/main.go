package main

import "log"

func main() {

	// for i := 0; i <= 5; i++ {
	// 	log.Println(i)
	// }

	// animals := []string{"cat", "dog", "elefant", "eule", "hund", "katze"}

	// for i, animal := range animals {
	// 	log.Println(i, animal)
	// }

	// // COMMENT: - to ignore the index use _ blank identifier.
	// for _, animal := range animals {
	// 	log.Println(animal)
	// }

	// animals := make(map[string]string)
	// animals["dog"] = "Fluffy"
	// animals["cat"] = "Suzzy"
	// for animalType, animal := range animals {
	// 	log.Println(animalType, animal)
	// }

	// // COMMENT: loop over the string 1
	// var firstLine = "Once upona midnight dreary"

	// for i, l := range firstLine {
	// 	log.Println(i, ":", l)
	// }

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@smith.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
