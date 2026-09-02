// This package provides tools for determining
// the date and time one gigasecond after a certain date.
package gigasecond

// Import the time package from the standard library
import "time"

// This function computes the date and time one gigasecond after a certain date.
func AddGigasecond(t time.Time) time.Time {

    // A gigasecond is one thousand million seconds. That is one with nine zeros after it.
	gigasecond := time.Second * 1000000000

    // Adding the gigasecond to the input time to produce a new date and time
    newTime := t.Add(gigasecond)
    
	return newTime
}
