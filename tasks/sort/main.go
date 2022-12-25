package main

import "fmt"

func main() {
	fmt.Println("insert sort")
	InsertSort(getArr())
	fmt.Println("bubble sort")
	BubbleSort(getArr())
	//
	fmt.Println("sort structs")
	users := Users{
		NewUser("Tommy", 30, 2000),
		NewUser("Jack", 30, 600),
		NewUser("Mark", 26, 1200),
		NewUser("Ruslan", 37, 2000),
		NewUser("Kamil", 27, 800),
	}

	SortByName(users)
	fmt.Printf("sort by name: %v\n", users)
	SortByAgeLess(users)
	fmt.Printf("sort by age less: %v\n", users)
	SortBySalaryLess(users)
	fmt.Printf("sort by salary: %v\n", users)
}
