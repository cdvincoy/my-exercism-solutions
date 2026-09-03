// This package provides tools for checking if the 
// provided string is a valid ISBN-10.
package isbnverifier

import (
    "strings"
)

// This function takes a string value and results a boolean value, 
// checking if it is a valid ISBN-10 or not.
func IsValidISBN(isbn string) bool {
	
    // Cleaning the string first, since some have hyphens
    cleanIsbn := strings.ReplaceAll(isbn, "-", "")
    
    // Similar to the previous activity, initializing variables for sum and multiplier
	sum:= 0
    multiplier:= 1

    // String is checked if it has exactly 10 characters
    if len(cleanIsbn) != 10 {
        return false
    }

    // A for loop is used to traverse digits from right to left
	for i:=len(cleanIsbn)-1; i>=0; i-- {
        
        digit := 0

        // An if condition to only allow 'X' as a check character, which is the 10th character or index 9
        if cleanIsbn[i] == 'X' && i != 9 {
            return false
        }

        // A switch case to convert a character into an integer
    	switch cleanIsbn[i] {
      		case 'X':
            	digit = 10
            case '9':
            	digit = 9
            case '8':
            	digit = 8
            case '7':
            	digit = 7
            case '6':
            	digit = 6
            case '5':
            	digit = 5
            case '4':
            	digit = 4
            case '3':
            	digit = 3
            case '2':
            	digit = 2
            case '1':
            	digit = 1
            case '0':
            	digit = 0
            default:
            	return false
    	}

        // The integer is multiplied by its assigned multiplier and added to the previous sum
        sum = (digit*multiplier) + sum
        // Since this is traversing right to left, increasing the multiplier by 1 is sufficient to obtain the assigned multiplier
        multiplier++
        
    }
    // Lastly, an if-else statement to check if it is a valid ISBN-10
    if sum%11 == 0 {
        return true
    } else {
        return false
    }
}
