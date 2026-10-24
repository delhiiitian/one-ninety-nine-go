// Chapter 4.3: Slicing, copying and the slices package
// Run: go run ./ch04-data-structures/03-slice-operations
package main

import (
	"fmt"
	"slices"
)

func main() {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	// s[low:high] includes low, excludes high
	q1 := months[0:3]
	q2 := months[3:6]
	secondHalf := months[6:]
	fmt.Println("Q1:", q1, "Q2:", q2)
	fmt.Println("second half:", secondHalf)

	// Careful: a sub-slice shares memory with the original!
	q1[0] = "JANUARY"
	fmt.Println("months[0] is now:", months[0])
	months[0] = "Jan"

	// copy makes an independent duplicate.
	backup := make([]string, len(q2))
	copy(backup, q2)
	backup[0] = "APRIL"
	fmt.Println("q2 unchanged:", q2, "backup:", backup)

	// Remove the item at index 2 from a list
	list := []string{"milk", "eggs", "bread", "jam"}
	list = slices.Delete(list, 2, 3)
	fmt.Println("after delete:", list)

	// Insert at index 1
	list = slices.Insert(list, 1, "butter")
	fmt.Println("after insert:", list)

	// Search
	fmt.Println("has eggs?", slices.Contains(list, "eggs"))
	fmt.Println("index of jam:", slices.Index(list, "jam"))

	nums := []int{5, 2, 9, 1}
	fmt.Println("max:", slices.Max(nums), "min:", slices.Min(nums))
	fmt.Println("clone equal?", slices.Equal(nums, slices.Clone(nums)))
}
