package main

import (
        "time"
	"fmt"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/raspi"
)

func main() {
        r := raspi.NewAdaptor()
        //led := gpio.NewLedDriver(r, "38") //Physical pin 38, GPIO 20.
        pump1 := gpio.NewLedDriver(r, "40") //Physical pin 40, GPIO 21.
        pump2 := gpio.NewLedDriver(r, "16") //Physical pin 16, GPIO 23.

        work := func() {
                gobot.Every(4*time.Second, func() {

                    fmt.Println("ON!");

                    pump1.On();
                    pump2.On();

                    time.Sleep(2 * time.Second)

                    fmt.Println("OFF!");
                    pump1.Off();
                    pump2.Off();

                    time.Sleep(2 * time.Second)

                })
        }

        robot := gobot.NewRobot("blinkBot",
                []gobot.Connection{r},
                //[]gobot.Device{led},
                []gobot.Device{pump1,pump2},
                work,
        )

        robot.Start()
}
