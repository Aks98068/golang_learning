package main

import (
	"fmt"
	"reflect"
)

// init function
func init() {
	fmt.Println("my nae is ranjit kumar sah")
}

func init() {
	fmt.Println("hello friends ")
}
func init() {
	var p person = person{"ranjit", "sah", 22, 5.6, "sah", "kumar", "male"}
	fmt.Printf("%+v\n", p)
}

// data structure for person
type person struct {
	firstname  string
	lastname   string
	age        int
	height     float32
	fathername string
	mothername string
	gender     string
}

// naked return functions
func naked(sum int) (a, b int) {
	a = sum * 4
	b = sum * 5
	return
}

func main() {

	fmt.Print("hello friends\n ")
	var p person = person{"ranjit", "sah", 22, 5.6, "sah", "kumar", "male"}
	// reflection of struct
	val := reflect.ValueOf(p)

	// iterate over the fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Printf("%s: %v\n", val.Type().Field(i).Name, field)

	}
	// anonymus function returning two string and functions literals also
	hello, helo2 := func() (string, string) {
		x := []string{"hello", "world"}
		return x[0], x[1]
	}()

	fmt.Print(hello, "\n")
	fmt.Print(helo2)

	a, b := naked(8)
	fmt.Print(a, b, "\n")
	fmt.Print(sum(55, 55, 44), "\n")
	fmt.Print(minus(55, 55, 44))

}

// variadic function for taking variable number of arguments and to sum them
func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func minus(sub ...int) int {
	minus := 0
	for _, sub := range sub[1:] {
		minus -= sub
	}
	return minus
}


