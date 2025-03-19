package main

import "fmt"

func main() {
	var gender string
	var age int
	var weight float32

	// Constants for dog size thresholds
	const smallDog = 10
	const mediumDog = 20
	const bigDog = 30

	// Input from the user
	fmt.Print("Enter the gender of the dog: ")
	fmt.Scanln(&gender)
	fmt.Print("Enter the age of the dog: ")
	fmt.Scanln(&age)
	fmt.Print("Enter the weight of the dog: ")
	fmt.Scanln(&weight)

	// Call the Gender function
	Gender(gender, age, weight, smallDog, mediumDog, bigDog)
}

func Gender(gender string, age int, weight float32, smallDog, mediumDog, bigDog int) {
	if gender == "male" {
		fmt.Println("The dog is a male")
		Male(age, weight, smallDog, mediumDog, bigDog)
	} else if gender == "female" {
		fmt.Println("The dog is a female")
		Female(age, weight)
	} else {
		fmt.Println("Invalid gender entered")
	}
}

func Male(age int, weight float32, smallDog, mediumDog, bigDog int) {
	if age < smallDog && weight > 20 {
		fmt.Println("The dog is a small dog")
		SmallDog(1, 2, 4)
	} else if age < mediumDog && age > smallDog && weight > 30 {
		fmt.Println("The dog is a medium dog")
		MediumDog(2, 4, 4)
	} else if age > mediumDog && age < bigDog && weight > 40 {
		fmt.Println("The dog is a big dog")
		BigDog(3, 4, 4)
	} else {
		fmt.Println("The dog's size could not be determined")
	}
}

func Female(age int, weight float32) {
	if age > 5 && weight > 20 {
		fmt.Println("The dog is a female")
	} else if age < 5 && weight < 20 {
		fmt.Println("The dog is a small female dog")
	} else {
		fmt.Println("The dog's size could not be determined")
	}
}

func SmallDog(tablet, cost, charge int) {
	total := cost + charge
	fmt.Println("The tablet for the small dog is: ", tablet)
	fmt.Println("The cost for the small dog is: ", cost)
	fmt.Println("The charge for the small dog is: ", charge)
	fmt.Println("The total for the small dog is: ", total)
}

func MediumDog(tablet, cost, charge int) {
	total := cost + charge
	fmt.Println("The tablet for the medium dog is: ", tablet)
	fmt.Println("The cost for the medium dog is: ", cost)
	fmt.Println("The charge for the medium dog is: ", charge)
	fmt.Println("The total for the medium dog is: ", total)
}

func BigDog(tablet, cost, charge int) {
	total := cost + charge
	fmt.Println("The tablet for the big dog is: ", tablet)
	fmt.Println("The cost for the big dog is: ", cost)
	fmt.Println("The charge for the big dog is: ", charge)
	fmt.Println("The total for the big dog is: ", total)
}
