/**

Keyboard-backlight settings for lubuntu on my macbook pro, 13 inch retina early 2015
`keyboard-backlight --up`
`keyboard-backlight --down`
This works on my macbook, it may cause yours to say mean things and slam shut on your fingers.
I also had to
`chown root:root keyboard-backlight`
`chmod u+s keyboard-backlight`
in order for the command to have to rights to change the backlight setting

robert arles
**/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const maxsetting = 200
const minsetting = 0
const defaultStep = 10

var myKeyboardDeviceBrightnessFile = "/sys/class/leds/smc::kbd_backlight/brightness"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {

		fmt.Printf("You must supply 1 parameter\n")
		fmt.Printf("\t--up\t\tto increase the keyboard backlight\n")
		fmt.Printf("\t--down\t\tto decrease the keyboard backlight\n")
		os.Exit(1)
	}
	if args[0] == "--up" {

		var brightnessSetting, err = increaseBrightnessSetting()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Brightness increased to %d\n", brightnessSetting)
		os.Exit(0)
	} else if args[0] == "--down" {

		var brightnessSetting, err = decreaseBrightnessSetting()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Brightness decreased to %d\n", brightnessSetting)
		os.Exit(0)
	} else {
		printHelp()
	}
}

func printHelp() {

	fmt.Printf("You must supply 1 parameter\n")
	fmt.Printf("\t--up\t\tto increase the keyboard backlight\n")
	fmt.Printf("\t--down\t\tto decrease the keyboard backlight\n")
}

func getBrightnessSetting() (response int16, err error) {

	var responseBytes []byte
	responseBytes, err = ioutil.ReadFile(myKeyboardDeviceBrightnessFile)
	if err != nil {
		return response, err
	}

	var response64 int64
	response64, err = strconv.ParseInt(strings.TrimSpace(string(responseBytes)), 10, 16)
	if err != nil {
		return response, err
	}

	response = int16(response64)

	return response, err
}

func setBrightness(desiredLevel int16) (currentSetting int16, err error) {

	currentSetting, err = getBrightnessSetting()
	if err != nil {
		return currentSetting, err
	}

	// if we are trying to set the level *above* max, then just set it to max
	if desiredLevel > maxsetting {
		desiredLevel = maxsetting
	}
	// if we are trying to set the level *below* min, then just set it to min
	if desiredLevel < minsetting {
		desiredLevel = minsetting
	}

	err = ioutil.WriteFile(myKeyboardDeviceBrightnessFile, []byte(strconv.FormatInt(int64(desiredLevel), 10)), 066)
	if err != nil {
		return currentSetting, err
	}

	currentSetting, err = getBrightnessSetting()
	if err != nil {
		return currentSetting, err
	}

	return currentSetting, err
}

func increaseBrightnessSetting() (resultantSetting int16, err error) {

	var currentSetting int16
	currentSetting, err = getBrightnessSetting()
	if err != nil {
		return currentSetting, err
	}

	var desiredSetting = currentSetting + defaultStep

	resultantSetting, err = setBrightness(desiredSetting)

	return resultantSetting, err
}

func decreaseBrightnessSetting() (resultantSetting int16, err error) {

	var currentSetting int16
	currentSetting, err = getBrightnessSetting()
	if err != nil {
		return currentSetting, err
	}

	var desiredSetting = currentSetting - defaultStep // subtracting a value higher than the uint16 causes a loop around to 655xx

	resultantSetting, err = setBrightness(desiredSetting)

	return resultantSetting, err
}
