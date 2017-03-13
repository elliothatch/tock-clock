package main

import (
	"flag"
	"github.com/fresh4less/neopixel-display/neopixeldisplay"
	"time"
)

const minutesLedCount = 60;
const hoursLedCount = 24;

func main() {
	var display neopixeldisplay.PixelDisplay
	noled := flag.Bool("noled", false, "If set, displays the clock as 256-color terminal stdout")
	flag.Parse()
	if *noled {
		display = neopixeldisplay.NewConsoleColorDisplay(minutesLedCount+hoursLedCount, [][]int{[]int{minutesLedCount,1}, []int{hoursLedCount,1}})
	} else {
		display = neopixeldisplay.NewNeopixelDisplay(18, minutesLedCount+hoursLedCount, 255)
	}
	clockDisplay := NewClockDisplay(display)
	clockDisplay.displayTime(time.Now())
}

