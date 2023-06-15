package main

import (
	"fmt"
)

func main() {
	// simpleStringSlice()
	sliceUsingVar()
	// exploreInt()
}

func exploreInt() {
	parentSlice := [3]int{1, 2, 3}
	fmt.Println("len(num1):", len(parentSlice))
	fmt.Println("cap(num1):", cap(parentSlice))
	childSlice := parentSlice[:2] // will have {1,2}
	fmt.Println(childSlice)
	fmt.Println("len(num2):", len(childSlice)) // length=2
	fmt.Println("cap(num2):", cap(childSlice)) // capacity =3
	childSlice = append(childSlice, 10)        // will have {1,2,10} and parentSlice will have the same {1,2,10}
	fmt.Println("parentSlice", parentSlice)
	fmt.Println("len(parentSlice):", len(parentSlice))
	fmt.Println("cap(parentSlice):", cap(parentSlice))
	fmt.Println("childSlice", childSlice)
	fmt.Println("len(childSlice):", len(childSlice))
	fmt.Println("cap(childSlice):", cap(childSlice))
	fmt.Println("End of using parent's array for child slice")
	fmt.Println("--------------------------------------------")
	fmt.Println("Starting  of creating seperate array for child slice as its max capaciy is above the parents array.")
	childSlice = append(childSlice, 20)
	fmt.Println("--------------------------------------------")
	fmt.Println("parentSlice", parentSlice)
	fmt.Println("len(parentSlice):", len(parentSlice))
	fmt.Println("cap(parentSlice):", cap(parentSlice))
	fmt.Println("--------------------------------------------")
	childSlice = append(childSlice, 30)
	fmt.Println("childSlice", childSlice)
	fmt.Println("len(childSlice):", len(childSlice))
	fmt.Println("cap(childSlice):", cap(childSlice))
	fmt.Println("--------------------------------------------")

}

func simpleStringSlice() {
	sports := []string{"Volleyball", "Swimming", "Running"}
	fmt.Println("Total of sports:", len(sports))
	fmt.Println("Capacity of sports:", cap(sports))
	addressOfUnderlyingArray := &sports
	fmt.Println("Address of underlying Array :", addressOfUnderlyingArray)
}

func sliceUsingVar() []string {
	var sports []string
	sports = append(sports, "Volleyball")
	fmt.Println("Address:", &sports[0])
	sports = append(sports, "Swimming")
	fmt.Println("Address:", &sports[0])
	sports = append(sports, "Running")
	fmt.Println("Address:", &sports[0])
	sports = append(sports, "Badminton")
	fmt.Println("Address:", &sports[0])
	sports = append(sports, "Cricket")
	fmt.Println("Address:", &sports[0])
	sports = append(sports, "Football")
	fmt.Println("Address:", &sports[0])
	return sports
}
