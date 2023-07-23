package for_test

func ForLoop() int {
	var sum int = 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	return sum
}