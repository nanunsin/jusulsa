package jusulsa

import (
	"fmt"
	"time"
)

func getCurTimeString(round int) (ret string) {
	h, m, s := time.Now().Round(time.Second * time.Duration(round)).Clock()
	//time.Now().Round(time.Second*10)
	ret = fmt.Sprintf("%02d%02d%02d", h, m, s)
	return ret
}

func printTimeStamp(timestamp time.Time, round int) {
	h, m, s := timestamp.Round(time.Second * time.Duration(round)).Clock()
	fmt.Printf("%02d%02d%02d", h, m, s)
}

func getTimeStamp(timestamp time.Time, round int) string {
	h, m, s := timestamp.Round(time.Second * time.Duration(round)).Clock()
	return fmt.Sprintf("%02d%02d%02d", h, m, s)
}

func getTimeStampIndex() int {
	now := time.Now()
	base := time.Date(now.Year(), now.Month(), now.Day(), 9, 00, 0, 0, now.Location())
	data := time.Now().Sub(base).Seconds()
	return (int)(data / 10)
}
