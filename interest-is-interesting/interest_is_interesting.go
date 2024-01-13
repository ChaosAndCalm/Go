package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {

	var interstrate float32

	if balance < 0 {
		interstrate = 3.213
	} else if balance >= 0 && balance < 1000 {
		interstrate = 0.5
	} else if balance >= 1000 && balance < 5000 {
		interstrate = 1.621
	} else {
		interstrate = 2.475
	}

	return interstrate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interstrate := InterestRate(balance)

	return float64(interstrate/100) * balance

}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {

	var noOfYears int

	for balance < targetBalance {
		balance = balance + Interest(balance)
		noOfYears++
	}

	return noOfYears
}
