package main

import "fmt"

func main() {
	//grades := [3]int{97, 85, 93}
	//grades := [...]int{97, 85, 93}
	//fmt.Printf("Grades: %v, %T\n", grades, grades)

	// var students [3]string
	// fmt.Printf("Students: %v\n", students)
	// students[0] = "Lisa"
	// students[1] = "Ajit"
	// students[2] = "Anvi"
	// fmt.Printf("Students: %v\n", students)
	// fmt.Printf("Students #1: %v\n", students[0])
	// fmt.Printf("Number of students: %v\n", len(students))

	// var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// identityMatrix := [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// fmt.Println(identityMatrix)

	//	a := [...]int{1, 2, 3}
	// COMMENT: Making a literal Copy of the array
	//b := a
	//fmt.Println(a)
	//fmt.Println(b)
	//b[1] = 5
	//fmt.Println("This is a literal copy of the first array hence we see the reassigned values:", b)

	// COMMENT: Array pointer with & sign
	// b := &a
	// fmt.Println(a)
	// fmt.Println(b)
	// a[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)

	//COMMENT: Introduction to SLICES
	// a := []int{1, 2, 3}
	// b := a // This points to the same data or location inside memory unlike array
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Printf("Length : %v\n", len(a))
	// fmt.Printf("Capacity : %v\n", cap(a))

	// //COMMENT: SLICING Operations
	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[:]
	// c := a[3:]
	// d := a[:6]
	// e := a[3:6]

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// fmt.Printf("Length %v\n", len(a))
	// fmt.Printf("Capacity %v\n", cap(a))

	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("Length %v\n", len(a))
	fmt.Printf("Capacity %v\n", cap(a))
}
