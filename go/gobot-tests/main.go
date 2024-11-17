package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/v2"
	"gobot.io/x/gobot/v2/drivers/gpio"
	"gobot.io/x/gobot/v2/platforms/raspi"
)

func main() {
	servo()
}

func newLedWork(ld *gpio.LedDriver) func() {
	return func() {
		gobot.Every(1*time.Second, func() {
			fmt.Println("toggling led...")
			if err := ld.Toggle(); err != nil {
				panic(err)
			}
			fmt.Println("after toggling light")
		})
	}

}

func servo() {
	r := raspi.NewAdaptor()
	sd := gpio.NewServoDriver(r, "12")

	work := func() {
		for {
			if err := sd.ToMin(); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
			if err := sd.ToCenter(); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
			if err := sd.ToMax(); err != nil {
				panic(err)
			}
			time.Sleep(1 * time.Second)
		}
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{r},
		[]gobot.Device{sd},
		work,
	)

	if err := robot.Start(); err != nil {
		panic(err)
	}

}
