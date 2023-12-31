package main

import (
	"fmt"
	"time"

	irsdk "github.com/quimcalpe/iracing-sdk"
)

func main() {
	var sdk irsdk.IRSDK
	sdk = irsdk.Init(nil)
	defer sdk.Close()

	for {
		sdk.WaitForData(100 * time.Millisecond)
		gear, _ := sdk.GetVar("CarIdxGear")
		speed, _ := sdk.GetVar("Speed")
		fmt.Printf("Gear: %s Speed: %s\n", gear, speed)
	}
}
