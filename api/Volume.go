package api

func (avr *AVReceiverStruct) volumeUp() {
	avr.sendCommand("PutMasterVolumeBtn/>")
}
func (avr *AVReceiverStruct) volumeDown() {
	avr.sendCommand("PutMasterVolumeBtn/<")
}
func (avr *AVReceiverStruct) getVolumeLevel() bool {
	return avr.getStateFor("volumeLevel")
}

func (avr *AVReceiverStruct) setVolumeLevel(level string) {
	avr.sendCommand("'PutMasterVolumeSet/" + level)
}
func (avr *AVReceiverStruct) SetMuteState(muted string) {
	// muted is "ON" or "OFF"
	avr.sendCommand("PutVolumeMute/" + muted)
}
func (avr *AVReceiverStruct) getMuteState() bool {
	return avr.getStateFor("mute")
}
