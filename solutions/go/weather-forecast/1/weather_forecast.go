// Package weather provides tools to forecast the 
// current weather condition of various cities in Goblinocus.
package weather

var (
    // CurrentCondition represents the current weather condition.
	CurrentCondition string
    // CurrentLocation represents a city paired with the current weather condition.
	CurrentLocation  string		
)

// Forecast function returns a string value.
// The string is about the current city with its current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
