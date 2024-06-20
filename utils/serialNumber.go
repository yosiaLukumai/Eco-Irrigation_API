package utils

import (
	"fmt"
	"strconv"
	"time"
)

func GenerateSerialNumber(companName string, branchNameShortForm string, currentC uint64) string {
	return fmt.Sprintf("%v-%v-%v", companName[0:2], branchNameShortForm, NumberGenerator(currentC))
}

// getting the time
func TimeFormats() string {
	return time.Now().Format("20060521")
}

func NumberGenerator(currentMeter uint64) string {
	if currentMeter < 10 {
		currentMeter++
		return fmt.Sprintf("00%v", currentMeter)
	} else if currentMeter < 99 {
		currentMeter++
		return fmt.Sprintf("0%v", currentMeter)
	} else {
		currentMeter++
		return strconv.FormatUint(currentMeter, 10)
	}
}
