// This package provides tools for
// calculating the points scored in a single toss of a Darts game.
package darts

import (
    "math"
)

// This function rewards 4 different amount of points, depending on where the dart lands. 
func Score(x, y float64) int {
    coordinates := (x*x + y*y)		
    radius := math.Pow(coordinates, 0.5) // Calcuate the radius based on the given coordinates.
    score := 0
    
	if radius <= 1 {   				// If the resulting radius ranges from 0 to 1, then that earns 10 points.
    	score = 10    
    } else if radius <= 5 {			// If the resulting radius ranges from 1 to 5, then that earns 5 points.
        score = 5
    } else if radius <= 10 {	   // If the resulting radius ranges from 5 to 10, then that earns 1 point.
        score = 1
    } else {					   // If the resulting radius is greater or lesser, which means it's out of the circle, that earns 0 point.
        score = 0
    } 
    return score
}
