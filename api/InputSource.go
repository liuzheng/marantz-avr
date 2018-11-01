package api

func (avr *AVReceiverStruct) getInputSource() bool {
	return avr.getStateFor("input")
}
func (avr *AVReceiverStruct) setInputSource(source string) {
	// source is "ON" or "OFF"
	avr.sendCommand("PutZone_InputFunction/" + source)
}
