package relu

func Relu(values []int) []int {
	var out []int = make([]int, 0, len(values))
	for _, v := range values {
		if v < 0 {
			v = 0
		}
		out = append(out, v)
	}

	return out
}
