package main

import (
	"flag"
	"github.com/fresh4less/neopixel-display/neopixeldisplay"
	"time"
)

const minutesLedCount = 59;
const hoursLedCount = 42;

func main() {
	var display neopixeldisplay.PixelDisplay
	noled := flag.Bool("noled", false, "If set, displays the clock as 256-color terminal stdout")
	test := flag.Bool("test", false, "Displays time at a rate determined by test-delay flag")
	testDelay := flag.Float64("test-delay", 1.0/60.0, "Delay in seconds between 'minutes' when running test command")
	flag.Parse()
	if *noled {
		display = neopixeldisplay.NewConsoleColorDisplay(minutesLedCount+hoursLedCount, [][]int{[]int{minutesLedCount,1}, []int{hoursLedCount,1}})
	} else {
		display = neopixeldisplay.NewNeopixelDisplay(18, minutesLedCount+hoursLedCount, 255)
	}
	clockDisplay := NewClockDisplay(display)

	if *test {
		for true {
			for h := 0; h < 24; h++ {
				for m := 0; m < 60; m++ {
					clockDisplay.displayTime(time.Date(0,0,0,h,m,0,0, time.Local))
					time.Sleep(time.Duration(*testDelay*float64(time.Second)))
				}
			}
		}
	}


	clock := NewClockTimer(nil)
	for(true) {
		now := <-clock.C
		clockDisplay.displayTime(now)
	}
}

