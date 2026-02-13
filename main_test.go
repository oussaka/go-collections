package gocollections

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

var P1 = Person{Name: "Haris", Age: 18}
var P2 = Person{Name: "Aslam", Age: 30}
var P3 = Person{Name: "Rashid", Age: 21}
var P4 = Person{Name: "Zaid", Age: 12}
var P5 = Person{Name: "Zahid", Age: 13}

func TestFind(t *testing.T) {
	tests := []struct {
		name     string
		persons  []Person
		want     string
		expected Person
		withErr  bool
	}{
		{"Empty slice", []Person{}, "Haris", Person{}, true},
		{"Single element", []Person{P1}, "Haris", P1, false},
		{"Multiple elements", []Person{P3, P1, P4, P1, P5, P2}, "Haris", P1, false},
		{"No elements found", []Person{P3, P4, P3, P5, P2}, "Haris", Person{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Find(tt.persons, func(p Person) bool { return p.Name == tt.want })
			if got != tt.expected {
				t.Errorf("Find(%v) = %v, want %v", tt.persons, got, tt.expected)
			}
			if tt.withErr && err == nil {
				t.Errorf("Expected error, but got none")
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		persons  []Person
		want     int
		expected []Person
		withErr  bool
	}{
		{"Empty slice", []Person{}, 17, []Person(nil), true},
		{"Single element", []Person{P1}, 17, []Person{P1}, false},
		{"Multiple elements", []Person{P1, P2, P3, P4, P5}, 17, []Person{P1, P2, P3}, false},
		{"Multiple same elements", []Person{P3, P1, P4, P1, P5, P2}, 17, []Person{P3, P1, P1, P2}, false},
		{"No elements found", []Person{P4, P5}, 17, []Person(nil), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Filter(tt.persons, func(p Person) bool { return p.Age > tt.want })
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Filter(%v) = %v, want %v", tt.persons, got, tt.expected)
			}
			if tt.withErr && err == nil {
				t.Errorf("Expected error, but got none")
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name    string
		persons []Person
		want    int
	}{
		{"Empty slice", []Person{}, 32},
		{"Single element", []Person{P1}, 32},
		{"Multiple elements", []Person{P1, P2, P3, P4, P5}, 32},
		{"Multiple same elements", []Person{P3, P1, P4, P1, P5, P2}, 18},
		{"No elements found", []Person{P4, P5}, 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gots := Map(tt.persons, func(p Person) Person {
				person := p
				person.Age = tt.want
				return person
			})
			for _, got := range gots {
				if got.Age != tt.want {
					t.Errorf("Map(%v) = %v, want Person.Age = (%v)", tt.persons, got, tt.want)
				}
			}
		})
	}
}

//
//func TestFilterEven(t *testing.T) {
//	tests := []struct {
//		name    string
//		numbers []int
//		want    []int
//	}{
//		{"Empty slice", []int{}, []int{}},
//		{"No even numbers", []int{1, 3, 5, 7, 9}, []int{}},
//		{"Only even numbers", []int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}},
//		{"Mixed numbers", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
//		{"Negative numbers", []int{-1, -2, -3, -4}, []int{-2, -4}},
//		{"Zero included", []int{0, 1, 2, 3}, []int{0, 2}},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got := FilterEven(tt.numbers)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("FilterEven(%v) = %v, want %v", tt.numbers, got, tt.want)
//			}
//		})
//	}
//}
