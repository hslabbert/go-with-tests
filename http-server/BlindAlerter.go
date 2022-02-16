package poker

import (
	"fmt"
	"io"
	"time"
)

// A BlindAlerter implements a ScheduleAlertAt() method.
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int, to io.Writer)
}

// A BlindAlerterFunc function takes a time & amount for
// blinds to provide an implementation-independent interface.
// This allows you to implement BlindAlerter with a function.
type BlindAlerterFunc func(duration time.Duration, amount int, to io.Writer)

// ScheduleAlertAt is BlindAlerterFunc implementation of BlindAlerter
// It hands a duration and time to the provided BlindAlerterFunc.
func (a BlindAlerterFunc) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	a(duration, amount, to)
}

// Alerter is a BlindAlerterFunc. It takes a time and amount,
// spinning that into a goroutine via time.AfterFunc() to print
// updated Blind values to os.Stdout.
func Alerter(duration time.Duration, amount int, to io.Writer) {
	time.AfterFunc(duration, func() {
		fmt.Fprintf(to, "Blind is now %d\n", amount)
	})
}
