package main

import (
	"flag"
	"time"

	"github.com/binkynet/bidib/host"
	"github.com/binkynet/bidib/transport/serial"
	"github.com/rs/zerolog"
)

func main() {
	portName := ""
	flag.StringVar(&portName, "port", "/dev/tty.usbserial-AB0LPGVA", "Name of serial port")
	flag.Parse()

	log := zerolog.New(zerolog.NewConsoleWriter())
	cfg := host.Config{
		Serial: &serial.Config{
			PortName: portName,
		},
	}
	h, err := host.New(cfg, log)
	if err != nil {
		log.Fatal().Err(err).Msg("host.New failed")
	}
	time.Sleep(time.Second * 20)
	h.Close()
}
