package factorial

// Factorial is called in main.go
func Factorial(i int) int {
	if i == 1 || i == 0 {
		return 1
	} else {
		return i * Factorial(i-1)
	}
}
