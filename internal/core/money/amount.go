package money

// Amount is an abstraction representing an amount of money.
type Amount int64

// Add adds one amount to another in place.
func (a *Amount) Add(other Amount) {
	*a += other
}

// Subtract subtracts one amount from another in place.
func (a *Amount) Subtract(other Amount) {
	*a -= other
}
