package main

import (
	"fmt"
	"github.com/liuzheng/marantz-avr"
)

func main() {
	var avr = marantz_avr.New("127.0.0.1", "default", "default")
	fmt.Println(avr.SurroundMode)
}
