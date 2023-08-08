package validatesequence

/*
TEST INPUT

array := []int{5, 1, 22, 25, 6, -1, 10, 8}
subArray := []int{1, 6, -1, 10}

*/

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.

	//markerArray := make(map[int]bool)

	tracker := 0

	for i := 0; i < len(array); i++ {

		if tracker == len(sequence) {
			return true
		}

		if array[i] == sequence[tracker] {
			tracker++
		}

		// for j := tracker; j < len(sequence); j++ {
		// 	if array[i] == sequence[j] {
		// 		tracker++
		// 		break
		// 	}
		// }
	}

	return tracker == len(sequence)

}
