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

// DefaultSleeper implements a time.Second-driven Sleep method.
type DefaultSleeper struct{}

// Sleep implements a time.Second-driven method on DefaultSleeper.
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
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
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
