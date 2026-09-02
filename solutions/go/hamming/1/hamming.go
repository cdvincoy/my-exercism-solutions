// This program attepts to calculate the Hamming distance betwwen two DNA strands.
package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
	// First, compare the length of the two strings.
    // If they are not of equal length, return an error message.
    // If yes, proceed to the rest of the program.
    if len(a) != len(b){
        return 0, errors.New("strings must have the same length")
    }

    // Count the number of differences on the two strings.
    count := 0

    // Compare each character on the two strings.
    // If they are different, increase count by 1.
	for i:=0; i <len(a); i++ {
        if a[i] != b[i] {
            count += 1
        }
    }
	// Since it expects to return two values, nil is used to indicate that there are no errors.
    return count, nil
}
