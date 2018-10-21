package marantz_avr

import (
	"github.com/liuzheng/marantz-avr/api"
)

var Name = "marantz_avr"
var Version = "0.1"

func New(ip, source, surroundMode string) api.AVReceiverStruct {
	var AVReceiver = api.AVReceiverStruct{
		IpAdderss: ip,
	}

	if source == "default" {
		AVReceiver.Source = api.DefatultSource
	} else if source == "current" {
		AVReceiver.SyncCurrentSource()
	}
	if surroundMode == "default" {
		AVReceiver.SurroundMode = api.DefaultSurroundMode
	} else if surroundMode == "current" {
		AVReceiver.SyncCurrentSurroundMode()
	}

	return AVReceiver
}
