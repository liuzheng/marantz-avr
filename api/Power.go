package api

func (avr *AVReceiverStruct) getPowerState() bool {
	return avr.getStateFor("power")
}
func (avr *AVReceiverStruct) SyncPowerState() bool {
	return avr.getPowerState()
}
func (avr *AVReceiverStruct) SetPowerState(state string) {
	// state is "ON" or "OFF"
	avr.sendCommand("PutZone_OnOff/" + state)
}
