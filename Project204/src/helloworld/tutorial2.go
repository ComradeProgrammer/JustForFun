package helloworld

/*练习一下go test*/
func IsEven(x int) bool {
	if x == 1 {
		return false
	} else if x%2 == 0 {
		return true
	} else if x%2 != 0 {
		return false
	}
	return false
}
