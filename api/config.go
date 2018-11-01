package api

import (
	"encoding/xml"
)

const (
	STATUS_URL  = "/goform/formMainZone_MainZoneXml.xml"
	POST_URL    = "/index.put.asp"
	COMMAND_URL = "/goform/AppCommand.xml"
	XML_Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

var (
	GetAllZonePowerStatus   = Command{CMD: CMD{Value: "GetAllZonePowerStatus", Id: 1}, Value: 1}
	GetDeletedNetworkSource = Command{CMD: CMD{Value: "GetDeletedNetworkSource", Id: 1}, Value: 1}
	GetDeletedSource        = Command{CMD: CMD{Value: "GetDeletedSource", Id: 1}, Value: 1}
	GetMuteStatus           = Command{CMD: CMD{Value: "GetMuteStatus", Id: 1}}
	GetRenameSource         = Command{CMD: CMD{Value: "GetRenameSource", Id: 1}}
	GetRestorerModeStatus   = Command{CMD: CMD{Value: "GetRestorerModeStatus", Id: 1}}
	GetSourceStatus         = Command{CMD: CMD{Value: "GetSourceStatus", Id: 1}}
	GetSurroundModeStatus   = Command{CMD: CMD{Value: "GetSurroundModeStatus", Id: 1}}
	GetVolumeLevel          = Command{CMD: CMD{Value: "GetVolumeLevel", Id: 1}}
	GetZoneName             = Command{CMD: CMD{Value: "GetZoneName", Id: 1}}

	Source = SourceStruct{
		GAME:           "GAME",
		CBL_SAT:        "SAT/CBL",
		NETWORK:        "NET",
		USB:            "USB/IPOD",
		TUNER:          "TUNER",
		DVD:            "DVD",
		BLUERAY:        "BD",
		HD_RADIO:       "HDRADIO",
		AUX1:           "AUX1",
		AUX2:           "AUX2",
		MEDIA_PLAYER:   "MPLAY",
		TV:             "TV",
		PHONO:          "PHONO",
		INTERNET_RADIO: "IRADIO",
		MXPORT:         "M-XPORT",
		NETHOME:        "NETHOME",
	}
	SurroundModes = SurroundModesStruct{
		MOVIE:       "MOVIE",
		MUSIC:       "MUSIC",
		GAME:        "GAME",
		PURE_DIRECT: "PURE DIRECT",
		DIRECT:      "DIRECT",
		STEREO:      "STEREO",
		STANDARD:    "STANDARD",
		SIMULATION:  "SIMULATION",
		AUTO:        "AUTO",
		LEFT:        "LEFT",
	}
	DefatultSource      = Source.TV
	DefaultSurroundMode = SurroundModes.AUTO
)

type SourceStruct struct {
	GAME           string
	CBL_SAT        string
	NETWORK        string
	USB            string
	TUNER          string
	DVD            string
	BLUERAY        string
	HD_RADIO       string
	AUX1           string
	AUX2           string
	MEDIA_PLAYER   string
	TV             string
	PHONO          string
	INTERNET_RADIO string
	MXPORT         string
	NETHOME        string
}
type SurroundModesStruct struct {
	MOVIE       string
	MUSIC       string
	GAME        string
	PURE_DIRECT string
	DIRECT      string
	STEREO      string
	STANDARD    string
	SIMULATION  string
	AUTO        string
	LEFT        string
}
type AVReceiverStruct struct {
	IpAdderss    string
	Limiter      int
	Source       string
	SurroundMode string
}
type Command struct {
	XMLName xml.Name `xml:"tx"`
	CMD     CMD
	Value   int      `xml:"value,omitempty"`
	Zone    string   `xml:"zone,omitempty"`
}
type CMD struct {
	XMLName xml.Name `xml:"cmd"`
	Id      int      `xml:"id,attr"`
	Value   string   `xml:",chardata"`
}
