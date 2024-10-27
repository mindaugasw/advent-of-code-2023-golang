package internal

// GreatestCommonDivisor formula from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GreatestCommonDivisor[T int | int64](a T, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// LeastCommonMultiple formula from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LeastCommonMultiple[T int | int64](a T, b T, integers ...T) T {
	result := a * b / GreatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = LeastCommonMultiple(result, integers[i])
	}

	return result
}
