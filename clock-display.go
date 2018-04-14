package main;
import (
	"time"
	"math"
	//"fmt"
	//"image/color"
	"github.com/fresh4less/neopixel-display/neopixeldisplay"
	"github.com/lucasb-eyer/go-colorful"
)

var dailyGradient GradientTable;
var hourlyHueStep = -360.0/24.0

func init() {
	for i := 0; i < 24; i++ {
		hue := math.Mod(float64(i)*hourlyHueStep + 225.0, 360.0)
		if hue < 0 {
			hue += 360.0
		}
		c := colorful.Hsv(hue, 1.0, 1.0)
		dailyGradient = append(dailyGradient, struct{Col colorful.Color; Pos float64}{c , float64(i)/24.0})
	}
}

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

	//minutesColor := neopixeldisplay.MakeColor(255, 0, 100)
	//hoursColor := neopixeldisplay.MakeColor(118, 0, 255)

	// tis the season
	//minutesColor := neopixeldisplay.MakeColor(10, 255, 10)
	//hoursColor := neopixeldisplay.MakeColor(10, 10, 255)

	// happy new year
	//minutesColor := neopixeldisplay.MakeColor(0, 0, 255)
	//hoursColor := neopixeldisplay.MakeColor(0, 125, 255)

	for i := 0; i < time.Minute(); i++ {
		//r, g, b, _ := colorful.Hsv(0.0, 0.8, 1.0).RGBA()
		//r, g, b, _ := colorful.Hsv(0.0, 1.0, 1.0).BlendHcl(
			//colorful.Hsv(30.0, 1.0, 1.0), float64(i)/minutesLedCount).Clamped().RGBA()
		t := (float64(i)/minutesLedCount)/24.0 + float64(time.Hour())/24.0
		r, g, b, _ := dailyGradient.GetInterpolatedColorFor(t).Clamped().RGBA()
		ledColor := neopixeldisplay.MakeColor(b>>8, g>>8, r>>8)
		minutesFrame.Set(58-i, 0, ledColor, neopixeldisplay.Error)
	}
	//if time.Hour() != 0 {
	hourIndex := (hoursLedCount) - int(
		(hoursLedCount*((float64(time.Hour()%12)/12.0) +
						(float64(time.Minute()%60)/60.0)/12.0)))
	for i := 0; i < hourIndex; i++ {
		//r, g, b, _ := colorful.Hsv(0.0, 1.0, 1.0).BlendHcl(
			//colorful.Hsv(30.0, 1.0, 1.0), float64(i)/hoursLedCount).Clamped().RGBA()
		//t := (float64(hourIndex-i)/hoursLedCount)*0.5 + float64(time.Hour()/12)*0.5
		t := (float64(time.Minute())/60.0)/24.0 + float64(time.Hour())/24.0 +
			(float64(hourIndex-i)/hoursLedCount)/2.0
		r, g, b, _ := dailyGradient.GetInterpolatedColorFor(t).Clamped().RGBA()
		ledColor := neopixeldisplay.MakeColor(b>>8, g>>8, r>>8)
		hoursFrame.Set(i, 0, ledColor, neopixeldisplay.Error)
	}
	//hoursFrame.Set(hourIndex, 0, hoursColor, neopixeldisplay.Error)
	//hoursFrame.Set((hourIndex-1)%hoursLedCount, 0, hoursColor, neopixeldisplay.Error)
	//hoursFrame.Set((hourIndex-2)%hoursLedCount, 0, hoursColor, neopixeldisplay.Error)
	//}

	cd.minutesScreen.Draw()
	cd.hoursScreen.Draw()
}
