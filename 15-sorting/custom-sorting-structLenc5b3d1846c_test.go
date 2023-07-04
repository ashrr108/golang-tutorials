package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestLenc5b3d1846c(t *testing.T) {
	// Test case 1: The slice is empty.
	u := UsersByAge{}
	if len(u) != 0 {
		t.Error("Expected len(u) to be 0, but got", len(u))
	}

	// Test case 2: The slice has one element.
	u = append(u, User{Age: 10})
	if len(u) != 1 {
		t.Error("Expected len(u) to be 1, but got", len(u))
	}

	// Test case 3: The slice has multiple elements.
	u = append(u, User{Age: 20}, User{Age: 30})
	if len(u) != 3 {
		t.Error("Expected len(u) to be 3, but got", len(u))
	}
}

type User struct {
	Age int
}

type UsersByAge []User

func (u UsersByAge) Len() int {
	return len(u)
}

func (u UsersByAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UsersByAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
