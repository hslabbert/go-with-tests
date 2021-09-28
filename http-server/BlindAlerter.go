package poker

import (
	"fmt"
	"os"
	"time"
)

// A BlindAlerter implements a ScheduleAlertAt() method.
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// A BlindAlerterFunc function takes a time & amount for
// blinds to provide an implementation-independent interface.
type BlindAlerterFunc func(duration time.Duration, amount int)

func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int) {
	a(duration, amount)
}

// StdOutAlerter takes a time and amount, spinning that into a
// goroutine via time.AfterFunc() to print updated Blind values
// to os.Stdout
func StdOutAlerter(duration time.Duration, amount int) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(os.Stdout, "Blind is now %d\n", amount)
	})
}
