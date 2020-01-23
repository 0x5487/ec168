package sport

import "time"

type MarketStatus int32

const (
	// INACTIVE indicates the market should NOT be displayed and bets on it should NOT be accepted
	INACTIVE MarketStatus = 1
	// ACTIVE indicates that the market should be displayed and bets on it should be accepted
	ACTIVE    MarketStatus = 2
	SUSPENDED MarketStatus = 3
	SETTLED   MarketStatus = 4
)

type MatchStatus int32

type Match struct {
	ID        string
	GameID    int64
	Status    MatchStatus
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Market struct {
	ID              string
	GameID          int64
	RoundID         string
	Name            string
	MarketStatus    MarketStatus
	StartedAt       time.Time `json:"started_at" db:"started_at"`
	EndedAt         time.Time `json:"ended_at" db:"ended_at"`
	SeattledAt      time.Time `json:"seattle_at" db:"seattle_at"`
	SeattledVersion string
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type Outcome struct {
	ID        string
	MarketID  string
	Name      string
	Odds      int32
	IsWin     bool
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type FindOutcomeOptions struct {
}

type MarketServicer interface {
	Markets() ([]Market, error)
	GetOutcomeSettlements(marketID int64) ([]Outcome, error)
	Outcomes() ([]Outcome, error)
}
