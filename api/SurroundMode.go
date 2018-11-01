package api

func (avr *AVReceiverStruct) getSurroundMode() string {
	 return avr.getStateFor("surroundMode")
}
func (avr *AVReceiverStruct) SyncCurrentSurroundMode() {
	avr.SurroundMode = avr.getSurroundMode()
}
