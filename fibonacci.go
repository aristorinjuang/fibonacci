package fibonacci

import "math"

var numbers = map[uint]uint{0: 0, 1: 1}

func multiply(f1 [2][2]uint, f2 [2][2]uint) [2][2]uint {
	x := f1[0][0]*f2[0][0] + f1[0][1]*f2[1][0]
	y := f1[0][0]*f2[0][1] + f1[0][1]*f2[1][1]
	z := f1[1][0]*f2[0][0] + f1[1][1]*f2[1][0]
	w := f1[1][0]*f2[0][1] + f1[1][1]*f2[1][1]

	f1[0][0] = x
	f1[0][1] = y
	f1[1][0] = z
	f1[1][1] = w

	return f1
}

func matrixPower(f1 [2][2]uint, index uint) [2][2]uint {
	if index <= 1 {
		return f1
	}

	f1 = matrixPower(f1, index/2)
	f1 = multiply(f1, f1)

	var f2 = [2][2]uint{
		{1, 1},
		{1, 0},
	}
	if index%2 != 0 {
		f1 = multiply(f1, f2)
	}

	return f1
}

// Recursion is a Fibonacci function with recursion method.
func Recursion(index uint) uint {
	if index <= 1 {
		return index
	}

	return Recursion(index-1) + Recursion(index-2)
}

// Memoization is a Fibonacci function with memoization method.
func Memoization(index uint) uint {
	if _, ok := numbers[index]; ok {
		return numbers[index]
	}

	if index <= 2 {
		numbers[index] = 1
	} else {
		numbers[index] = Memoization(index-1) + Memoization(index-2)
	}

	return numbers[index]
}

// Iteration is a Fibonacci function with iteration method.
func Iteration(index uint) uint {
	if index <= 1 {
		return index
	}

	var i, result, prev1, prev2 uint = 2, 1, 0, 0

	for i <= index {
		prev2 = result
		result += prev1
		prev1 = prev2
		i++
	}

	return result
}

// Matrix is a Fibonacci function with matrix method.
func Matrix(index uint) uint {
	if index <= 1 {
		return index
	}

	f := [2][2]uint{
		{1, 1},
		{1, 0},
	}

	f = matrixPower(f, index-1)

	return f[0][0]
}

// Math is a Fibonacci function with math method.
func Math(index uint) uint {
	return uint(math.Round(math.Pow(float64((1+math.Sqrt(5))/2), float64(index)) / math.Sqrt(5)))
}
