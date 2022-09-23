package gamegrp

// Status represents the game status.
type Status struct {
	Status  string `json:"status"`
	Address string `json:"address"`
}

// NewBet is what we require from clients when adding a Bet.
type NewBet struct {
	Description      string      `json:"description"`
	Terms            string      `json:"terms"`
	Amount           int         `json:"amount"`
	ModeratorAddress string      `json:"moderatorAddress" validate:"required"`
	Players          []NewPlayer `json:"players" validate:"required"`
	DateExpired      uint        `json:"dateExpired" validate:"required"`
}

// NewPlayer represents the connection between a new Bet and an Account that is in
// a player role.
type NewPlayer struct {
	Address string `json:"address"`
	InFavor bool   `json:"inFavor"`
}

// UpdateBet is what we require from clients when updating a Bet.
type UpdateBet struct {
	Description      *string `json:"description"`
	Terms            *string `json:"terms"`
	Amount           *int    `json:"amount"`
	ModeratorAddress *string `json:"moderatorAddress" validate:"required"`
	DateExpired      *uint   `json:"dateExpired" validate:"required"`
}
