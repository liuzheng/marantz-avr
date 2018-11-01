package api

import (
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
)

func (avr *AVReceiverStruct) getState() bool {
	resp, err := http.Get("http://" + avr.IpAdderss + STATUS_URL)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(body)
	return false
}
func (avr *AVReceiverStruct) getStateFor(name string) bool {
	return avr.getState()
}

func (avr *AVReceiverStruct) sendCommand(cmd string) (string, error) {

	resp, err := http.Post(
		"http://"+avr.IpAdderss+"/MainZone"+POST_URL,
		"text/html",
		strings.NewReader(cmd))
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body), err
}
