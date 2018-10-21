package api

const (
	STATUS_URL = "/goform/formMainZone_MainZoneXml.xml"
	POST_URL   = "/index.put.asp"
)

var (
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
