package main;
import (
	"time"
	//"math"
	//"fmt"
	"github.com/fresh4less/neopixel-display/neopixeldisplay"
)

type ClockDisplay struct {
	pixelDisplay neopixeldisplay.PixelDisplay;
	minutesScreen *neopixeldisplay.ScreenView;
	hoursScreen *neopixeldisplay.ScreenView;
}

func NewClockDisplay(pixelDisplay neopixeldisplay.PixelDisplay) *ClockDisplay {
	return &ClockDisplay{
		pixelDisplay,
		neopixeldisplay.NewScreenView(pixelDisplay, 0, minutesLedCount, 1, 1.0, neopixeldisplay.Error),
		neopixeldisplay.NewScreenView(pixelDisplay, minutesLedCount, hoursLedCount, 1, 1.0, neopixeldisplay.Error),
	}
}

func (cd *ClockDisplay) displayTime(time time.Time) {
	minutesFrame := cd.minutesScreen.GetFrame()
	hoursFrame := cd.hoursScreen.GetFrame()

	minutesFrame.SetAll(neopixeldisplay.MakeColor(0,0,0))
	hoursFrame.SetAll(neopixeldisplay.MakeColor(0,0,0))

	for i := 0; i < time.Minute(); i++ {
		minutesFrame.Set(58-i, 0, neopixeldisplay.MakeColor(255, 0, 100), neopixeldisplay.Error)
	}
	hoursColor := neopixeldisplay.MakeColor(118, 0, 255)
	//if time.Hour() != 0 {
	hourIndex := (hoursLedCount) - int(
		(hoursLedCount*((float64(time.Hour()%12)/12.0) +
						(float64(time.Minute()%60)/60.0)/12.0)))
	for i := 0; i < hourIndex; i++ {
		hoursFrame.Set(i, 0, hoursColor, neopixeldisplay.Error)
	}
	//hoursFrame.Set(hourIndex, 0, hoursColor, neopixeldisplay.Error)
	//hoursFrame.Set((hourIndex-1)%hoursLedCount, 0, hoursColor, neopixeldisplay.Error)
	//hoursFrame.Set((hourIndex-2)%hoursLedCount, 0, hoursColor, neopixeldisplay.Error)
	//}

	cd.minutesScreen.Draw()
	cd.hoursScreen.Draw()
}
