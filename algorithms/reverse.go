package algorithms

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var stringy string
	for _, char := range word {
		stringy = string(char) + stringy
	}
	return stringy
}
