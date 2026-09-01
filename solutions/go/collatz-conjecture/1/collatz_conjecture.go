// This package provides tools in calculating the number of steps it takes to 
// reach 1 according to the rules of the Collatz Conjecture.
package collatzconjecture

import (
    "errors"
)

// This function returns two values, an integer and a potential error message.
func CollatzConjecture(n int) (int, error) {

    steps := 0					// The steps start at 0.

    if n == 0 || n < 0 {
        return 0, errors.New("value must be greater than 0")
    }
    
    for n > 1 {					// The loop stops if n is equal to 1.
        if n%2 == 0 {			// Divide the number if it is even.
            n = n/2
            steps += 1			// Increase number of steps by 1.
        } else {
        	n = (n*3) + 1		// Muliply by 3 and add 1 if it is odd.
            steps += 1			// Increase number of steps by 1.
        }
	}
	return steps, nil
}
