# GoBasics
Basics of Go Programming Language


# GO Cheatsheet 

## Variables
### declare variable
var name string = "John"

### map 
var m map[string]string

a := make(map[string]string)
a := map[string]string{"key1":"value1", "key2":"value2"}

### array
var a [5]int // array of 5 integers
a := [5]int{1,2,3,4,5} // array of 5 integers initialized with values

### slice
var s []int // slice of integers
a := make([]int, 5) // slice of integers initialized with 5 zeros
a := []int{1,2,3,4,5} // slice of integers initialized with values


### multiple variable declaration
var a, b, c int = 1, 2, 3

var (
    a string = "initial"
    b int 
)

### struct
type person struct {
    name string
    age int
}

a := person{name: "John", age: 30}
a := person{"John", 30}

a.name = "John"
a.age = 30

## Types
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32 // represents a Unicode code point
float32 float64
complex64 complex128

## Conditions

### comparison operators
== equal to
!= not equal to
<  less than
<= less than or equal to
>  greater than
>= greater than or equal to

### Logical operators
&& and
|| or
!  not

### If else & switch 
if x > 10 {
    // do something
} else if x < 5 {
    // do something else
} else {
    // do yet another thing
}

switch x {
    case 1:
        // do something
    case 2:
        // do something else
    default:
        // do yet another thing //optional
}

### Arithmetic operators
+  addition
-  subtraction
*  multiplication
/  division
%  modulus
++ increment by 1
-- decrement by 1

## Slices
a := make([]int, 5) // slice of integers initialized with 5 zeros
a := []int{} 
a := []int{1,2,3,4,5}
