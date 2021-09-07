package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// func Reverse(word string) string {
// 	var result string
// 	for _, v := range word {
// 		result = string(v) + result
// 	}
// 	return result
// }
func Reverse(word string) string {
	l := len(word)
	if l == 0 {
		return ""
	}
	return string(word[l-1]) + Reverse(string(word[:l-1]))
}
