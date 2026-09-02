// This package provides tools for determining whether a number is
// valid according to the Luhn formula.
package luhn

import (
    "strings"
)

func Valid(id string) bool {
	clean_id := strings.ReplaceAll(id, " ", "")  // Removing the spaces from id

    if len(clean_id) <= 1 {						// Strings of length 1 or less are not valid
        return false
    } 

 	var sum int									// Initialized a variable sum for the sum of all the digits
    double := false								// Initialized a variable double for determining whether to double a digit

    for i:=len(clean_id)-1; i>=0; i-- {			// Using a for loop to traverse the digits from right to left
        digit := 0

        switch clean_id[i]{						// Using a switch case to convert each digit into an integer
            case '0':
            	digit = 0
            case '1':
            	digit = 1
            case '2':
            	digit = 2
            case '3':
            	digit = 3
            case '4':
            	digit = 4
            case '5':
            	digit = 5
            case '6':
            	digit = 6
            case '7':
            	digit = 7
            case '8':
            	digit = 8
            case '9':
            	digit = 9
            default:
            	return false
        }

        if double {					 // An if statement to determine whether to double the digit or not (ex. for the first iteration, it does not double the digit)
            digit = digit *2

            if digit > 9 {			// Another if statement to check if the digit is less than or greater than 9. If it is greater then 9, 9 is subtracted from the digit.
                digit = digit - 9
            }
        }
        
        double = !double			// At the end of each iteration, the variable double is flipped to the opposite value.
        sum = sum + digit			// Finally, the digit is added to the sum variable.
    } 
    if sum%10 == 0 {				// Lastly, if the sum of the digits is evenly divisible by 10, then the number is valid.
        return true
    } else {						// Otherwise, it is not. 
        return false
    }
}