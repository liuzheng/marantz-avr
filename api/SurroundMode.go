package api

func (avr *AVReceiverStruct) getCurrentSurroundMode() string {
	return ""
}
func (avr *AVReceiverStruct) SyncCurrentSurroundMode() {
	avr.SurroundMode = avr.getCurrentSurroundMode()
}
