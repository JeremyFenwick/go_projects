package algorithms

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {
	boolFriend := false
	for _, number := range list {
		if number == num {
			boolFriend = true
		}
	}
	return boolFriend
}
