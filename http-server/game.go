package poker

import "time"

// A TexasHoldem holds a PlayerStore and a BlindAlerter that
// schedules alerts to be fired for blind raises at
// pre-specified intervals.
type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

// Start will start running a *Game of the provided numberOfPlayers,
// setting up the needed blind alerts using the BlindAlerter in
// *Game alerter.
func (p *TexasHoldem) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

// Finish records a win into the PlayerStore of the provided
// *Game for the provided winner.
func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}

// NewTexasHoldem constructs a *Game. This is provided as a helpful
// constructor given that we don't want to export a Game struct's
// internal properties, but do wish to be able to set up a Game
// from external packages.
func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}
