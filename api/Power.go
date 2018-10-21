package api

func (avr *AVReceiverStruct) getPowerState(cmd string) bool {
	return avr.getStateFor("power")
}
