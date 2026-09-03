package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	floatPR := float64(productionRate)
	return floatPR * (successRate/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerMinute := CalculateWorkingCarsPerHour(productionRate, successRate) / 60
    intCarsPM := int(carsPerMinute)
    return intCarsPM
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens := carsCount / 10
    remaining := carsCount%10

    cost := (tens*95000) + (remaining*10000)
    finCost := uint(cost)

    return finCost
}
