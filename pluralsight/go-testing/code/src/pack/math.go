package pack

func Add(nums ...int) int {
	result := 0

	for _, i := range nums {
		result += i
	}
	return result

}
