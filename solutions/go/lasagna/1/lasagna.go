package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	remainMinutes := OvenTime - actualMinutesInOven 
    return remainMinutes
}

// PreparationTime calculates the time needed to prepare he lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	const Time = 2
    preparationTime := numberOfLayers * 2
    return preparationTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	elaspedTime := numberOfLayers*2 + actualMinutesInOven
    return elaspedTime
}
