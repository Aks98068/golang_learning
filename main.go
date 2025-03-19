package main

import (
	"fmt"
	"go1/help"
	"strings"
	// Importing the utils package
)

// entry point of the program
func main() {

	fmt.Println("hello abhishekh kumar sah")

	var boy string // string
	boy = "abhishekh kumar sah"

	land(boy)

	myname := "abhishekh kumar sah"
	fmt.Println(myname)
	// if statement
	if a == 2 {
		say(boy)
	}

	// concatinating the strings
	adding := "hello" + "world"
	fmt.Println(len(adding), adding[0:6])

	num := 345           //int
	numfloat := 3.       //float64
	numcomplex := 3 + 4i //complex128
	numbyte := byte(num) // byte (alias: unit8)

	var u uint = 7       //unit(unsigned)
	var p float32 = 22.7 //32 bit float

	fmt.Println(num, u, p, numfloat, numcomplex, numbyte)
	// increament
	num++
	fmt.Println(num)
	// decreament
	num--
	fmt.Println(num)

	// logical operatores of and or not
	istrue := true
	isfalse := false
	fmt.Println(istrue && isfalse)
	fmt.Println(istrue && !isfalse)
	fmt.Println(istrue || isfalse)
	fmt.Println(!istrue)
	fmt.Println(true && false)
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(!true)
	fmt.Println(!false)

	// arrays
	primes := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes, len(primes)) // Outputs: [2 3 5 7 11 13] 6
	_array := [3]string{"fdsnndf", "sdfdsf", "sdfdsf"}
	fmt.Println(_array, len(_array)) // Outputs: [fdsnndf sdfdsf sdfdsf] 3
	var ab [6]string
	ab[0] = "hello"
	ab[1] = "world"
	ab[2] = "abhishekh"
	ab[3] = "kumar"
	ab[4] = "sah"
	fmt.Println(ab, len(ab), ab[1], ab[2:3])

	two_array := [5][6]int{{1, 22, 34, 4546, 54}, {32443, 45435, 5657, 67, 45}}
	fmt.Println(two_array, len(two_array), two_array[1][2])

	var arr1 = []int{} //empty array
	fmt.Println(arr1, len(arr1))
	fmt.Println(cap(arr1))     // capacity of the array
	arr2 := make([]int, 5, 10) // creating array with 5 elements and capacity of 10
	fmt.Println(arr2, len(arr2), cap(arr2))
	// pointers
	getpointer(&num) // passing the memory address of the variable a to the function getpointer
	fmt.Println()
	val1 := &num
	fmt.Println(val1) // return the memory address of the variable a
	fmt.Println(*val1)
	ptr := *val1 + 40
	fmt.Println(ptr)

	// slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 6)
	fmt.Println(slice, len(slice), cap(slice))

	// maps
	// creating a map
	mymap := make(map[string]int)
	mymap["one"] = 1
	mymap["two"] = 2
	mymap["three"] = 3
	mymap["four"] = 4
	mymap["five"] = 5
	fmt.Println(mymap, len(mymap), mymap["one"])

	// string functions
	data := "hrllo ,abhishekh ,kumar ,sah"
	data2 := "      hello"
	split1 := strings.Split(data, ",") // spliting the string with comma
	fmt.Println(split1)

	count1 := strings.Count(data, "a") // counting the number of a in the string
	fmt.Println(count1)
	index1 := strings.Index(data, "k") // finding the index of the first occurance of a in the string
	fmt.Println(index1)

	join1 := strings.Join([]string{data, data2}, " ") // joining the two strings
	fmt.Println(join1)

	replace1 := strings.Replace(data, "hrllo", "hello", -1) // replacing the a with b in the string
	fmt.Println(replace1)
	repeat1 := strings.Repeat(data, 3) // repeating the string 3 times
	fmt.Println(repeat1)

	trim1 := strings.TrimSpace(data2) // removing the white spaces from the string
	fmt.Println(trim1)

	upper1 := strings.ToUpper(data)
	fmt.Println(upper1)

	lower1 := strings.ToLower(upper1)
	fmt.Println(lower1)

	special1 := strings.Contains(data, "kumar")
	fmt.Println(special1)

	compare1 := strings.Compare(data, data2)
	fmt.Println(compare1)

	//strcut
	type person struct {
		firstname  string
		lastname   string
		age        int
		height     float32
		fathername string
		mothername string
		gender     string
	}
	abhishekh := person{
		firstname:  "abhishekh",
		lastname:   "sah",
		age:        20,
		height:     5.5,
		fathername: "subodh sah",
		mothername: "devi",
		gender:     "male",
	}
	gyani := person{
		firstname:  "gyani",
		lastname:   "sah",
		age:        20,
		height:     5.5,
		fathername: "prem lal sah",
		mothername: "devi",
		gender:     "male",
	}
	fmt.Println(abhishekh)
	fmt.Println(gyani)
	fmt.Println(abhishekh.age)

	fmt.Printf("%v\n", abhishekh)  //printing the variable
	fmt.Printf("%+v\n", abhishekh) //printing the variable with field names
	fmt.Printf("%#v\n", abhishekh) //printing the variable with field names
	fmt.Printf("%T\n", abhishekh)  //pritning the type of vaiable
	fmt.Printf("%t\n", abhishekh.age > gyani.age)
	fmt.Printf("%d\n", abhishekh.age)                    //printing the integer
	fmt.Printf("%b\n", abhishekh.height)                 //printing the binary
	fmt.Printf("%c\n", abhishekh.lastname[0])            //printing the character
	fmt.Printf("%x\n", abhishekh.lastname)               //printing the hex
	fmt.Printf("%f\n", abhishekh.height)                 //printing the float
	fmt.Printf("%s\n", abhishekh.lastname)               //printing the string
	fmt.Printf("%q\n", abhishekh.lastname)               //printing the double quoted string
	fmt.Printf("%x\n", abhishekh.firstname)              //printing the hex
	fmt.Printf("%p\n", &abhishekh)                       //printing the pointer
	fmt.Printf("%e\n", 12340000000.0)                    //printing the scientific notation
	fmt.Printf("%E\n", 12340000000.0)                    //printing the scientific notation
	fmt.Printf("|%6d|%6d|\n", 12, 345)                   //printing the decimal
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)             //printing the float
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)           //printing the float
	fmt.Printf("|%6s|%6s|\n", "foo", "b")                //printing the string
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")              //printing the string
	fmt.Printf("|%6t|%6t|\n", true, false)               //printing the boolean
	fmt.Printf("|%-6t|%-6t|\n", true, false)             //printing the boolean
	fmt.Printf("|%6v|%6v|\n", abhishekh, gyani)          //printing the value
	fmt.Printf("|%-6v|%-6v|\n", abhishekh, gyani)        //printing the value
	fmt.Printf("|%6+|%-6+|\n", abhishekh.age, gyani.age) //printing the sign
	fmt.Printf("|%6#|%-6#|\n", abhishekh, gyani)         //printing the alternate form
	fmt.Printf("|%6T|%6T|\n", abhishekh, gyani)          //printing the type
	fmt.Printf("|%-6T|%-6T|\n", abhishekh, gyani)        //printing the type
	fmt.Printf("|%6t|%6t|\n", true, false)               //printing the boolean
	fmt.Printf("|%-6t|%-6t|\n", true, false)             //printing the boolean
	fmt.Printf("|%6d|%6d|\n", 12, 345)                   //printing the decimal
	fmt.Printf("|%-6d|%-6d|\n", 12, 345)                 //printing the decimal
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)             //printing the float
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)           //printing the float
	fmt.Printf("|%6s|%6s|\n", "foo", "b")                //printing the string
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")              //printing the string
	fmt.Printf("|%6x|%6x|\n", 12, 345)                   //printing the hex
	fmt.Printf("|%-6x|%-6x|\n", 12, 345)                 //printing the hex
	fmt.Printf("|%6p|%6p|\n", &abhishekh, &gyani)        //printing the pointer

	fmt.Printf("%x\n", 80) //printing the hex
	help.Greet("Abhishekh")
	ss := help.Sad // Replace 'SomeValidIdentifier' with an actual identifier from the 'help' package
	fmt.Println(ss)
	fmt.Println(d)

	fmt.Printf("%T\n", help.Type1)
	fmt.Printf("%T\n", help.Type2)

	help.Condition1()

	result, err := help.Dothing()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	help.Switch1()
	help.Loop1()
	fmt.Println("sum is:")
	help.Range1()
	var A = 1
	var B = 2
	var C = 3
	rr := 1

	//calling arrow function
	Outputs, int1 := help.Foo(rr)
	println(Outputs, int1)

	//calling multiple arguments
	results, int2 := help.Multi1(A, B, C)
	println(results, int2)

}

// decelearing mupltiple variables at once
var a, b int32 = 1, 2 //integer
var d = true          //bool

func say(message string) {
	var d = true
	fmt.Println("you said :", message, d)
}
func land(me string) {
	// shrot deceleration of variable
	f, e := 1, 3
	fmt.Println(a, b, me, f, e)
}

func getpointer(mypointer *int) {
	v := mypointer
	fmt.Println(v) // return the memory address of the variable b
}
