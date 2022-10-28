package buzzer

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

var (
	high uint = 1
	low  uint = 0
)

func ConfigureBuzzer(gpioNumber uint) {

	cmd := exec.Command("raspi-gpio", []string{
		"set",
		fmt.Sprintf("%d", gpioNumber),
		"op",
		"pn",
		"dh",
	}...)

	err := cmd.Run()
	if err != nil {
		log.Printf("[ERROR] Failed to configure pin %v", gpioNumber)
	}
}

func digitalWrite(gpioNumber, level uint) {

	var state string

	if level == low {
		state = "dh"
	} else {
		state = "dl"
	}

	cmd := exec.Command("raspi-gpio", []string{
		"set",
		fmt.Sprintf("%d", gpioNumber),
		state,
	}...)

	err := cmd.Run()
	if err != nil {
		log.Printf("[ERROR] Failed to write value to pin: %v", err)
	}
}

func BeepBuzzer(gpioNumber, timeout, times uint) {
	var v uint

	log.Printf("[INFO] Beeping buzzer %v times ...", times)

	for v = 0; v < times; v++ {
		digitalWrite(gpioNumber, high)
		time.Sleep(time.Millisecond * time.Duration(timeout))
		digitalWrite(gpioNumber, low)
		time.Sleep(time.Millisecond * time.Duration(timeout))
	}
}

func SingleBeep(gpioNumber, timeout uint) {
	BeepBuzzer(gpioNumber, timeout, 1)
}
