package main

import "fmt"

func main() {
	// Arrays
	// Store items of one type and are fixed length
	// Array of ints of length 5
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a, "length:", len(a))

	// Can also let the compiler infer the length of the array
	b := [...]string{"a", "b"}
	fmt.Println(b, "length:", len(b))

	// Slices
	// Slices are similar to array but they are not fixed in size
	c := []int{1, 2, 3, 4, 5}
	fmt.Println(c) // [1,2,3,4,5]

	// we can extend c
	c = append(c, 6)
	fmt.Println(c) // [1,2,3,4,5,6]

	// append can also take a slice as a second arg as long as it is unfurled
	d := append(c, []int{7, 8, 9}...)
	fmt.Println(d) // [1,2,3,4,5,6,7,8,9]

	// slicing works simiarly to Python with the colon operator
	// e.g. get a subslice from index 2 (inclusive) to end
	c = c[2:]
	fmt.Println(c) // [3,4,5,6]

	/*
		A slice always has an underlying array where the value are stored
		and to which it points. A slice has two properties:
		- length: the number of values stored in the underlying array
		- capacity: the size of the underlying array
		When we append elements to a slice two things can happen:
		1 - if the length of the new modified slice is below or equal to its
			capacity, the underlying array stays the same
		2 - if the lenght of the modified slice is above its capacity, the
			underlying array is substituted with one with a size that can accomodate
			the increase number of elements. By the looks of it, the array is doubled
			in size every time the capacity limit is breached (assume is to avoid
			discarding and recreating the underlying array every time a single element is
			appended so it is not too costly performance wise to extend the slice)

			Looking at the source code it looks like the doubling rule applies to "small"
			slices and the growth rate is reduced for bigger ones.
			(see: https://cs.opensource.google/go/go/+/refs/tags/go1.23.1:src/runtime/slice.go;l=288)

			Also if the new length is more than double the capacity, different rules apply
	*/
	test := []int{1, 2, 3}
	fmt.Println(len(test)) // 3
	fmt.Println(cap(test)) // 3
	test = append(test, 1, 2)
	fmt.Println(len(test)) // 5
	fmt.Println(cap(test)) // 6
	test = append(test, 1, 2)
	fmt.Println(len(test)) // 7
	fmt.Println(cap(test)) // 12
	test = append(test, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(test)) // 14
	fmt.Println(cap(test)) // 24
	test = append(test, make([]int, 35)...)
	fmt.Println(len(test)) // 49
	fmt.Println(cap(test)) // 52 (???, why not 49?)

	// When we assign a slice to anothe variable, they share the same underlyinga array
	one := []int{1, 2, 3}
	two := one

	two[2] = 4

	fmt.Println(one) // [1,2,4]
	fmt.Println(two) // [1,2,4]

	// to avoid modifying the underlying array we need to create a copy instead
	three := []string{"a", "b", "c"}
	four := make([]string, 3)
	copy(four, three)

	four[2] = "d"
	fmt.Println(three) // ["a","b","c"]
	fmt.Println(four)  // ["a","b","d"]

	/*
		https://www.adityathebe.com/anatomy-of-go-slices/
		Notes on capacity

		While above I stated "- capacity: the size of the underlying array", that is not correct

		The actual definition is:
		"The Capacity of a slice is equal to the length of the underlying array,
		minus the index in the array of the first element of the slice."


		See sample code below:
	*/
	parentSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Printf("Parent Slice's Address: %p\n\n", parentSlice)

	sliceA := parentSlice[:5]
	fmt.Printf("sliceA's Address: %p\n", sliceA)        // same as parentSlice
	fmt.Printf("sliceA's Capacity %d\n\n", cap(sliceA)) // sliceA's Capacity 20

	// Has same underlying array (parentSlice) as sliceA but its first element
	// points to a different index on it
	sliceB := parentSlice[12:15]
	fmt.Printf("sliceB's Address: %p\n", sliceB)      // will be the address of parentSlice's 12th element
	fmt.Printf("sliceB's Capacity %d\n", cap(sliceB)) // sliceB's Capacity 8

}
