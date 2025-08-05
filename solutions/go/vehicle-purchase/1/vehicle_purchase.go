package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	finalOptionIs1 := option1 < option2

	personalizedMessage := " is clearly the better choice."

	if finalOptionIs1 {
		return option1 + personalizedMessage
	}

	return option2 + personalizedMessage
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	const (
		MaxCarAgeForTwentyPerCentDiscount float64 = 3.0
		MaxCarAgeForThirtyPerCentDiscount float64 = 10.0
		TwentypPerCentMultiplier          float64 = 0.2
		ThirtypPerCentMultiplier          float64 = 0.3
		FiftyPerCentMultiplier            float64 = 0.5
	)

	if age < MaxCarAgeForTwentyPerCentDiscount {
		return originalPrice - (originalPrice * TwentypPerCentMultiplier)
	}

	if age < MaxCarAgeForThirtyPerCentDiscount {
		return originalPrice - (originalPrice * ThirtypPerCentMultiplier)
	}

	return originalPrice - (originalPrice * FiftyPerCentMultiplier)
}
