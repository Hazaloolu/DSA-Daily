package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	WorkingCarsPerHour := float64(productionRate) * successRate/100
    return WorkingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsProducedPerHour := float64(productionRate) * successRate/100
    carsProducedPerMinute := carsProducedPerHour/60
    return int(carsProducedPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const costPerCar = 10000
    const costPerTenCars = 95000

    carUnits := carsCount/10
    individualCar := carsCount%10

    totalCost := carUnits*costPerTenCars + individualCar*costPerCar
    return uint(totalCost)

    
    
}
