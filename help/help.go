package help

import (
	"fmt"
	"strconv"
)

// Function to greet a name
func Greet(Name string) {
	fmt.Printf("Hello, %s!\n", Sad)
	fmt.Printf("Name: %s\n", Sad)
}

var Sad = "123"

const Type1 = 123

var Type2 = strconv.Itoa(Type1)

var Condi1 = 20

func Condition1() {
	// conditional statement
	if Condi1 > 10 {
		fmt.Println("Condi1 is greater than 10")
	} else if Condi1 < 10 {
		fmt.Println("Condi1 is less than 10")
	} else {
		fmt.Println("Condi1 is equal to 10")
	}
}
func Dothing() (string, error) {
	return "", fmt.Errorf("error found")
}

func Switch1() {
	//switch statement
	switch Condi1 {
	case 10:
		fmt.Println("Condi1 is 10")
	case 20:
		fmt.Println("Condi1 is 20")
	case 30:
		fmt.Println("Condi1 is 30")
	default:
		fmt.Println("Condi1 is not 10 or 20")
	}
}

func Loop1() {
	// for loop
	for i := 0; i < Condi1; i++ {
		fmt.Println(i)
	}
}

func Range1() {
	// for range loop
	Arr := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, Arr := range Arr {
		sum += Arr

	}
	fmt.Println(sum)

	nums := []int{2, 3, 4}
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// while loop
	i := 0
	for i <= sum {
		fmt.Println(i)
		i++
		fmt.Println("hello")
	}

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		fmt.Println("for")
	}
	// break statement

}

// arrow function
var Foo = func(rr int) (int, string) {
	return rr, "arrow"
}
 
// multiple arguments
var Multi1 = func(A, B, C int) (int, string) {
	return A + B + C, "multi"
}
