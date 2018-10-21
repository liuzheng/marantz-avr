package main

import (
	"github.com/liuzheng/marantz-avr"
)

func main() {
	marantz_avr.New("127.0.0.1", "default", "default")
}
