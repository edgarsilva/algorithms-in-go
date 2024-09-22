package main

import "fmt"

// Clarifies once and for all the difference of creating new variables from arrays
// or/and slices, composite types vs discrete types.

//  1. Arrays behavior - discrete type
//     creating a new array var from another array allocates new memory
//     and copies over the whole chunk of memory from the original array
//     to the new location.
//
// This might be confusing with the behavior of slices (composite type),
// which share the same backing array (pointer) and memory allocation for it,
// when you create a new slice variable from another one.
//
//  2. Slice behavior - composite type
//     if you create a new slice var from an existing slice, the same mem
//     location for the backing array gets carried over, which makes sense
//     since we are only copying the pointer address of the backing array
//     to the new variable, this changes once we start appending to the new slice,
//     now if there's not enouch capacity a new backing array is created and a
//     corresponding new block of memory with enough size is allocated for it,
//     this slice gets a new backing array pointer, and now the two slices no
//     longer share the same backing array memory.
func main() {
	arrayBehavior()
	fmt.Println("------------------------------------")
	sliceBehavior()
}

func arrayBehavior() {
	arr1 := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Arr1 was defined as [1,2,3,4,5]:")
	fmt.Printf("%#v\n", arr1)
	fmt.Printf("arr1 mem address %p\n\n", &arr1)

	fmt.Println("When creating a new array variable(arr2) from an existing array(arr1)")
	fmt.Println("the whole array (the whole chunk of mem) is copied into a new memory location,")
	fmt.Printf("so modifing the second array does not modify the array from wich it was originated.\n\n")

	arr2 := arr1
	fmt.Printf("Arr2 was defined as arr2 := arr1(%#v)", arr1)
	fmt.Printf("%#v\n", arr2)
	fmt.Printf("arr2 mem address %p\n\n", &arr2)

	fmt.Println("So modifing index 0 and 1 to the values 6,7 won't affect the original arr1")
	arr2[0] = 6
	arr2[1] = 7
	fmt.Printf("Arr2 %#v\n", arr2)
	fmt.Printf("arr2 mem address %p\n\n", &arr2)
}

func sliceBehavior() {
	sli1 := []int{1, 2, 3, 4, 5}

	fmt.Println("sli1 was defined as [1,2,3,4,5]:")
	fmt.Printf("%#v\n", sli1)
	fmt.Printf("sli1 mem address %p\n\n", &sli1)

	fmt.Println("When creating a new slice var(sli2) from an existing slice(sli1)")
	fmt.Println("the two share the same backing array (pointer), same memory location.")
	fmt.Printf("Modifing either of the two updates the same backing array in memory.\n\n")

	sli2 := sli1
	fmt.Printf("sli2 was defined as sli2 := sli1(%#v)\n", sli1)
	fmt.Printf("%#v\n", sli2)
	fmt.Printf("sli2 mem address %p\n\n", &sli2)

	fmt.Println("Modifing position 0 and 1 in sli2 with values 6,7 also updates sli1")
	fmt.Println("since they are using the same backing array in memory")
	sli2[0] = 6
	sli2[1] = 7
	fmt.Printf("sli1 %#v\n", sli1)
	fmt.Printf("sli2 %#v\n\n", sli2)

	fmt.Println("Once the cap of the slice is not enough to append more elements")
	fmt.Println("a new backing array with enough capacity needs to be allocated,")
	fmt.Println("and the backing array pointer in the slice updated, now sli2 points")
	fmt.Printf("to a diff backing array in memory, and modifing either won't affect the other.\n")

	sli2 = append(sli2, 8, 9, 10)
	fmt.Printf("sli2.append(8, 9, 10)\n")
	sli2[0] = 1
	sli2[1] = 2

	fmt.Printf("sli1 %#v\n", sli1)
	fmt.Printf("sli2 %#v\n\n", sli2)

	fmt.Println("Since the original slice is still pointing to the original location in")
	fmt.Println("memory, for its backing array now updating either won't affect the other.")
}
