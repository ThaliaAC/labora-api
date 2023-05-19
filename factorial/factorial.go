package factorial

func Factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
