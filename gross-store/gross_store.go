package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {

	grossMap := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return grossMap

}

// NewBill creates a new bill.
func NewBill() map[string]int {
	p := map[string]int{}
	return p
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]

	if !ok {
		return false
	}

	bill[item] += units[unit]
	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	val, ok := bill[item]

	if !ok {
		return false
	}

	quantity, unitCheck := units[unit]
	if !unitCheck {
		return false
	}

	if val-quantity < 0 {
		return false
	} else if val-quantity == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = val - quantity
	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, ok := bill[item]

	if !ok {
		return 0, false
	}

	return val, true

}
