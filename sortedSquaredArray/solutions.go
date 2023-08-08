package sortedsquaredarray

/*
TEST INPUT

//inputArray := []int{-7, -3, 1, 9, 22, 30}

*/
func SortedSquaredArray(array []int) []int {

	// Write your code here.
	returnArray := make([]int, len(array))

	for i := 0; i < len(array); i++ {
		absVal := array[i] * array[i]

		//inserting to output array in ascending order
		for j := 0; j < len(returnArray); j++ {

			if absVal > returnArray[j] && j < len(returnArray) {

				temp := returnArray[j]
				returnArray[j] = absVal

				if j == 0 {
					returnArray[0] = temp
					returnArray[j] = absVal
				} else {
					returnArray[j-1] = temp
				}
			}
		}
	}

	return returnArray
}
