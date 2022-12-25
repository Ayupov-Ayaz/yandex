package main

import "sort"

type User struct {
	Age    uint8
	Salary uint64
	Name   string
}

func NewUser(name string, age uint8, salary uint64) User {
	return User{
		Age:    age,
		Salary: salary,
		Name:   name,
	}
}

type Users []User

func sortBy(u []User, comparator func(i, j User) bool) {
	sort.Slice(u, func(i, j int) bool {
		return comparator(u[i], u[j])
	})
}

func SortByName(u []User) {
	sortBy(u, func(i, j User) bool {
		return i.Name < j.Name
	})
}

func SortByAgeLess(u []User) {
	sortBy(u, func(i, j User) bool {
		return i.Age < j.Age
	})
}

func SortBySalaryLess(u []User) {
	sortBy(u, func(i, j User) bool {
		return i.Salary < j.Salary
	})
}
