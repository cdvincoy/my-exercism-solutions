// This program determines whether a year is a leap year or not.
package leap

// This function checks different conditions to determine if it's a leap year or not.
func IsLeapYear(year int) bool {
	//str := ""
    status := false

    if year%100 == 0 {										// A year can be a leap year if it is divisible by 100,
        if year%400 == 0 {									// and if it is divisible by 400.
            //str = str + "This is a leap year!"
            status = true
        } else {
            //str = str + "This is not a leap year!"			// Or else, it is not a leap year.
            status = false
        }
    } else if year%4 == 0 {									// A year can also be a leap year if it is divisible by 4.
         //str = str + "This is a leap year!"
         status = true
    } else {
		//str = str + "This is not a leap year!"        		// Or else, it is not a leap year.
        status = false
    }
    return status
}
