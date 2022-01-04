package main

import "fmt"

//Set class
type Set struct {
	integerMap map[int]bool
}

//create the map of integer and bool
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

//checks if element is in the set
func (set *Set) ContainsElement(element int) bool {
	var exists bool
	_, exists = set.integerMap[element]
	return exists
}

// adds the element to the set
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

//deletes the element from the set
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

//Intersect method returns the set which intersects with anotherSet
func (set *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()
	var value int
	for value = range set.integerMap {
		if anotherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

func (set3 *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	var value int
	for value = range set3.integerMap {
		unionSet.AddElement(value)
	}
	for value = range anotherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}

func main() {
	var set *Set
	set = &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println("initial set", set)
	fmt.Println(set.ContainsElement(1))

	var anotherSet *Set
	anotherSet = &Set{}
	anotherSet.New()
	anotherSet.AddElement(2)
	anotherSet.AddElement(4)
	anotherSet.AddElement(5)

	fmt.Println(set.Intersect(anotherSet))

	fmt.Println(set.Union(anotherSet))

}
