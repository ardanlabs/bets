package bet

// Bet represents an individual bet.
type Bet struct {
	ID int `json:"id"` // Unique identifier.
}

// NewBet is what we require from clients when adding a Bet.
type NewBet struct {
	BetID int `json:"betId" validate:"required"`
}
