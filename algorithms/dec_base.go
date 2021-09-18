package algorithms

import (
	"fmt"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	var answer string
	const charset = "0123456789ABCDEF"
	for dec > 0 {
		x := dec % base
		answer = fmt.Sprintf("%v", string(charset[x])) + answer
		dec = dec / base
	}
	return answer
}
