package jusulsa

import (
	"fmt"
	"time"
)

type Mark1 struct {
	Date    string
	ObjInfo []QryInfo
	Code    string
}

// NewBot func Create New Bot one
func NewBot(code string) *Mark1 {
	y, m, d := time.Now().Date()
	today := fmt.Sprintf("%04d%02d%02d", y, m, d)
	bot := &Mark1{
		Date: today,
		Code: code,
	}
	return bot
}

func (bot *Mark1) Print() {
	for i := 0; i < len(bot.ObjInfo); i++ {
		fmt.Printf("%s\t", bot.ObjInfo[i].TimeStr)
		fmt.Printf("%d(%d)\t", bot.ObjInfo[i].Data.Price, bot.ObjInfo[i].Volume)
		fmt.Printf("%d\n", bot.ObjInfo[i].Curve)
	}
}

func (bot *Mark1) queryWorks() {
	bot.ObjInfo = append(bot.ObjInfo, *QueryData(bot.Code))
	size := len(bot.ObjInfo) - 1

	if size == 0 {
		makeInfoStep1(nil, &bot.ObjInfo[0])
	} else {
		makeInfoStep1(&bot.ObjInfo[size-1], &bot.ObjInfo[size])
	}
}

func analyzeWork() {
	return
}
