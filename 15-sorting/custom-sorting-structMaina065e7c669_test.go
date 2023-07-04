package main

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
	Age  int
}

// UsersByAge is a custom sorting function that sorts a slice of User
// by their age in ascending order.
func UsersByAge(users []User) func(i, j int) bool {
	return func(i, j int) bool {
		return users[i].Age < users[j].Age
	}
}

func TestMain(t *testing.T) {
	// Test case 1: Check if the users are sorted by age in ascending order
	users := []User{
		{
			Name: "Rajeev",
			Age:  28,
		},
		{
			Name: "Monica",
			Age:  31,
		},
		{
			Name: "John",
			Age:  56,
		},
		{
			Name: "Amanda",
			Age:  16,
		},
		{
			Name: "Steven",
			Age:  28,
		},
	}

	sort.Sort(UsersByAge(users))

	for i := 0; i < len(users)-1; i++ {
		if users[i].Age > users[i+1].Age {
			t.Error("Users are not sorted by age in ascending order")
		}
	}

	// Test case 2: Check if the users are sorted by age in descending order
	users = []User{
		{
			Name: "Rajeev",
			Age:  28,
		},
		{
			Name: "Monica",
			Age:  31,
		},
		{
			Name: "John",
			Age:  56,
		},
		{
			Name: "Amanda",
			Age:  16,
		},
		{
			Name: "Steven",
			Age:  28,
		},
	}

	sort.Stable(UsersByAge(users))

	for i := 0; i < len(users)-1; i++ {
		if users[i].Age < users[i+1].Age {
			t.Error("Users are not sorted by age in descending order")
		}
	}
}
