package utils

func Sum(numbers []int) int {
	ttsum := 0
	//for i := 0; i < len(numbers); i++ {
	//	ttsum += numbers[i]
	//}

	for _, num := range numbers {
		ttsum += num
	}
	return ttsum
}
