// ex4.4 rotates a slice of ints by one position to the left.
package rotate

import (
	"fmt"
)

func rotate_ints(ints []int) {
	first := ints[0]
	copy(ints, ints[1:])
	ints[len(ints)-1] = first
}

// space is O(n)
// This function left rotate an user inut number of shifts with O(n) space complexity
func rotate(s []int,shift int)[]int{  
	var newdata []int
	shift = shift%len(s)
	first := s[:shift]
	newdata = append(newdata,s[shift:]...)
	newdata = append(newdata,first...)
	return newdata
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate_ints(s)
	fmt.Println(s)
}
