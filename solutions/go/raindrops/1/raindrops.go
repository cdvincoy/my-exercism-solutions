package raindrops

import (
    "strconv" 	// package needed to convert number as  a string
)

func Convert(number int) string {
	str := ""  // empty string to return result later
    
	if number % 3 == 0 {
        str = str + "Pling"
    }
    
    if number % 5 == 0 {
        str = str + "Plang"
    }

    if number % 7 == 0 {
        str = str + "Plong"
    }

	if str == "" {
        str = strconv.Itoa(number)
    }

    return str
}
