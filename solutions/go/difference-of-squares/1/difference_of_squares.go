// This package provides tools for finding the difference
// between the square of the sum and the sum of the squares of the first N natural numbers.
package differenceofsquares

// This function calculates the square of the sum of numbers from 1 to n.
func SquareOfSum(n int) int {
    sum := 0
	for i:= 1; i <= n; i++ {
        sum = sum + i
    }
    squareOfSum := sum*sum
    return squareOfSum
}

// This function calculates the sum of the squares of numbers from 1 to n.
func SumOfSquares(n int) int {
	squared := 0
    for i:= 1; i <= n; i++ {
        squared = squared + i*i
    }
    return squared
}

// The difference between square of the sum and sum of the squares is calculated by calling the two function earlier and subtracting them.
func Difference(n int) int {
	difference := SquareOfSum(n) - SumOfSquares(n)
    return difference
}
