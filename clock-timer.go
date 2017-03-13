package main;
import (
	"time"
)

//Pushes the wall-clock time to its channel every second, corrects drift

const ClockTimerBuffer = 255;

type ClockTimer struct {
	loc *time.Location;
	C chan time.Time;
	timer *time.Timer;
	Running bool;
}

func NewClockTimer(loc *time.Location) *ClockTimer {
	if(loc == nil) {
		loc = time.Local
	}
	clock := ClockTimer{
		loc,
		make(chan time.Time, ClockTimerBuffer),
		time.NewTimer(DurationUntilNextSecond()),
		true,
	}

	clock.Start()

	return &clock
}

func (ct *ClockTimer) Start() {
	ct.Running = true
	go func() {
		for(true) {
			<-ct.timer.C
			if(!ct.Running) {
				break;
			}

			ct.C <- time.Now().In(ct.loc)
			//restart timer
			ct.timer.Reset(DurationUntilNextSecond())
		}
	}()
}

func (ct *ClockTimer) Stop() {
	ct.Running = false
}

//millisecond precision
func DurationUntilNextSecond() time.Duration {
	return (1000 - time.Duration(time.Now().Nanosecond())/time.Millisecond)*time.Millisecond
}

