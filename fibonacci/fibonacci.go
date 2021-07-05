package fibonacci

func Fibonacci(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 || n == 1 {
		return n
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
