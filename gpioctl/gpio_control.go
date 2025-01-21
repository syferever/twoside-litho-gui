package twosidelithogui

import (
	"log"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

type Pin struct {
	Conf  gpio.PinIO
	Mode  string
	Level gpio.Level
}

func NewPin(name, mode string) *Pin {
	new_pin := gpioreg.ByName(name)
	if new_pin == nil {
		log.Println("Failed to find pin ", name)
	}
	return &Pin{Conf: new_pin, Mode: "NS", Level: gpio.Low}
}

func HostInit() {
	_, err := host.Init()
	if err != nil {
		log.Println("Failed to initialize periph: ", err)
	}
}

func (pin *Pin) SetPinMode(mode string) {
	switch mode {
	case "OUT", "IN":
		pin.Mode = mode
	default:
		log.Println("Mode is incorrect. Waited OUT or IN, recieved:", mode)
	}
}

func (pin *Pin) SetPinLevel(lvl string) {
	switch lvl {
	case "LOW":
		pin.Level = gpio.Low
	case "HIGH":
		pin.Level = gpio.High
	default:
		log.Println("Pin level is incorrect. Waited LOW or HIGH, recieved: ", lvl)
	}
	switch pin.Mode {
	case "OUT":
		err := pin.Conf.Out(pin.Level)
		if err != nil {
			log.Println("Failed to set pin as output:", err)
		}
	case "IN":
		log.Println("Pin mode is set to IN, couldn't run it as OUT.")
	case "NS":
		log.Println("Pin mode is not set.")
	}
}
