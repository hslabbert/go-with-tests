package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// A Sleeper implements a Sleep() method
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper implements an arbitrary Sleep duration.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep will sleep for c.duration using the c.sleep function.
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown takes an io.Writer write a countdown from 3, 2, 1, Go!
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
