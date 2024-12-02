package functions

func ConvertAndCubing(readChannel <-chan uint8, writechannel chan<- float64) {
	for num := range readChannel {
		floatNum := float64(num)
		writechannel <- floatNum * floatNum * floatNum
	}
	close(writechannel)
}
