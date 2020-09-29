package structs

type stringList []string

//NewStringList ... A counstructor for stringList type
func NewStringList(list ...string) stringList {
	var l stringList = list
	return l
}

/*
Iterative ... A iterative implementation of binary search on a sorted (ascending) slice of strings
Output: Index of the first occurence of the element if exists or -1 otherwise
*/
func (sl stringList) Iterative(x string) int {
	firstpoint, lastpoint := 0, (len(sl) - 1)
	for firstpoint != lastpoint {
		midpoint := (firstpoint + lastpoint) / 2
		if x <= sl[midpoint] {
			lastpoint = midpoint
		} else {
			firstpoint = midpoint + 1
		}
	}
	if x != sl[firstpoint] {
		return -1
	}
	return firstpoint
}

/*
Recursive ... A recursive implementation of binary search on a sorted (ascending) slice of strings
Output: Index of the first occurence of the element if exists or -1 otherwise
*/
func (sl stringList) Recursive(x string) int {
	firstpoint, lastpoint := 0, (len(sl) - 1)
	return sl.recursive(x, firstpoint, lastpoint)
}

func (sl stringList) recursive(x string, firstpoint, lastpoint int) int {
	if firstpoint == lastpoint {
		if x == sl[firstpoint] {
			return firstpoint
		}
		return -1
	}
	midpoint := (firstpoint + lastpoint) / 2
	if x <= sl[midpoint] {
		return sl.recursive(x, firstpoint, midpoint)
	}
	return sl.recursive(x, midpoint+1, lastpoint)
}
