package structs

type intList []int

//NewIntList ... A counstructor for intList type
func NewIntList(list ...int) intList {
	var l intList = list
	return l
}

/*
Iterative ... A iterative implementation of binary search on a sorted (ascending) slice of integers
Output: Index of the first occurence of the element if exists or -1 otherwise
*/
func (il intList) Iterative(x int) int {
	firstpoint, lastpoint := 0, (len(il) - 1)
	for firstpoint != lastpoint {
		midpoint := (firstpoint + lastpoint) / 2
		if x <= il[midpoint] {
			lastpoint = midpoint
		} else {
			firstpoint = midpoint + 1
		}
	}
	if x != il[firstpoint] {
		return -1
	}
	return firstpoint
}

/*
Recursive ... A recursive implementation of binary search on a sorted (ascending) slice of integers
Output: Index of the first occurence of the element if exists or -1 otherwise
*/
func (il intList) Recursive(x int) int {
	firstpoint, lastpoint := 0, (len(il) - 1)
	return il.recursive(x, firstpoint, lastpoint)
}

func (il intList) recursive(x, firstpoint, lastpoint int) int {
	if firstpoint == lastpoint {
		if x == il[firstpoint] {
			return firstpoint
		}
		return -1
	}
	midpoint := (firstpoint + lastpoint) / 2
	if x <= il[midpoint] {
		return il.recursive(x, firstpoint, midpoint)
	}
	return il.recursive(x, midpoint+1, lastpoint)
}
