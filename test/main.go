package main

import (
	"github.com/liuzheng/marantz-avr"
	"fmt"
)

func main() {
	var avr = marantz_avr.New("127.0.0.1", "default", "default")
	fmt.Println(avr.SurroundMode)
}
