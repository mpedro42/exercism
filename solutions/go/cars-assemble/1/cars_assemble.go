package cars

const (
	SingleCarCostProduction int     = 10000
	GroupCarsCounter        int     = 10
	GroupCarProduceDiscount float64 = 0.05
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	const PercentageDivisor float64 = 100.0
	successRate = successRate / PercentageDivisor

	return successRate * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	const HourInMinutes int = 60

	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / HourInMinutes
}

func CarGroupHandler(carsCount int) (int, int) {
	amountCarGroups := carsCount / GroupCarsCounter
	remainingCars := carsCount % GroupCarsCounter

	return amountCarGroups, remainingCars
}

func CarGroupProdutionCostsHandler(amountCarGroups int) (int, int) {
	carGroupCostProduction := amountCarGroups * SingleCarCostProduction * GroupCarsCounter
	carGroupCostProductionWithDiscount := carGroupCostProduction - int(float64(carGroupCostProduction)*GroupCarProduceDiscount)

	return carGroupCostProduction, carGroupCostProductionWithDiscount
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	amountCarGroups, remainingCars := CarGroupHandler(carsCount)
	_, carGroupCostProductionWithDiscount := CarGroupProdutionCostsHandler(amountCarGroups)

	return uint(carGroupCostProductionWithDiscount + (remainingCars * SingleCarCostProduction))

}
