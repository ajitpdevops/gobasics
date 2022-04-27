package main

import "fmt"

// COMMENT: iota starts from 0 and increments
// COMMENT: << bit shifting

// const (
// 	_ = iota  // ignore first value by assigning to blank identifier
// 	catSpecialist
// 	dogSpecialist
// 	parrotSpecialist
// )

// const (
// 	_  = iota
// 	KB = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// )

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// COMMENT: First set of Constants
	//	var specialistType int = 1
	//	fmt.Printf("%v\n", specialistType == catSpecialist)

	// COMMENT: Second set of constants
	// fileSize := 4000000000.
	// fmt.Printf("%.2fGB\n", fileSize/GB)

	// COMMENT: Third set of constants
	// println(isAdmin, isHeadquarters, canSeeFinancials, canSeeAfrica, canSeeAsia, canSeeEurope, canSeeNorthAmerica, canSeeSouthAmerica)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
}
