package api

func (avr *AVReceiverStruct) getPowerState() bool {
	return avr.getStateFor("power")
}
func (avr *AVReceiverStruct) SyncPowerState() bool {
	return avr.getPowerState()
}
