package main;
import (
	"time"
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
		neopixeldisplay.NewScreenView(pixelDisplay, 0, hoursLedCount, 1, 1.0, neopixeldisplay.Error),
	}
}

func (cd *ClockDisplay) displayTime(time time.Time) {
	minutesFrame := cd.minutesScreen.GetFrame()
	hoursFrame := cd.hoursScreen.GetFrame()

	minutesFrame.SetAll(neopixeldisplay.MakeColor(0,0,0))
	hoursFrame.SetAll(neopixeldisplay.MakeColor(0,0,0))

	minutesFrame.Set(time.Minute(), 0, neopixeldisplay.MakeColor(255,255,255), neopixeldisplay.Error)
	hoursFrame.Set(time.Hour()%12, 0, neopixeldisplay.MakeColor(255,255,255), neopixeldisplay.Error)
	hoursFrame.Set((time.Hour()%12)+1, 0, neopixeldisplay.MakeColor(255,255,255), neopixeldisplay.Error)

	cd.minutesScreen.Draw()
	cd.hoursScreen.Draw()
}
