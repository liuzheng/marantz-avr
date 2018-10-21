package api

func (*AVReceiverStruct) getCurrentSource() string {
	return ""
}
func (avr *AVReceiverStruct) SyncCurrentSource() {
	avr.Source = avr.getCurrentSource()
}
