package bubblesort

import "fmt"

func BubbleSortArray(input []int) []int {

	fmt.Println(input)

	swapped := true

	for {

		// array is not swapped in the previous iteration
		// so list is ordered, break.
		if !swapped {
			break
		}

		swapped = false

		for i := 0; i < len(input)-1; i++ {

			//replace current with next item
			if input[i] > input[i+1] {
				temp := input[i+1]
				input[i+1] = input[i]
				input[i] = temp
				swapped = true
			}
		}
	}

	return input
}
