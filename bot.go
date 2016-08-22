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

func (bot *Mark1) PrintAll() {
	for i := 0; i < len(bot.ObjInfo); i++ {
		//fmt.Printf("%s\t", bot.ObjInfo[i].TimeStamp.String())
		printTimeStamp(bot.ObjInfo[i].TimeStamp, 10)
		fmt.Printf("\t%d(%d)", bot.ObjInfo[i].Data.Price, bot.ObjInfo[i].Volume)
		fmt.Printf("\t%d", bot.ObjInfo[i].Curve)
		fmt.Println()
	}
}

func (bot *Mark1) PrintAt(index int) {
	if len(bot.ObjInfo) >= index {
		//fmt.Printf("%s\t", bot.ObjInfo[index].TimeStamp.String())
		printTimeStamp(bot.ObjInfo[index].TimeStamp, 10)
		fmt.Printf("\t%d(%d)", bot.ObjInfo[index].Data.Price, bot.ObjInfo[index].Volume)
		fmt.Printf("\t%d", bot.ObjInfo[index].Curve)
		fmt.Println()
	}
}

func (bot *Mark1) QueryWorks() {
	bot.ObjInfo = append(bot.ObjInfo, *QueryData(bot.Code))
	size := len(bot.ObjInfo) - 1

	if size == 0 {
		makeInfoStep1(nil, &bot.ObjInfo[0])
	} else {
		makeInfoStep1(&bot.ObjInfo[size-1], &bot.ObjInfo[size])
	}
}

func AnalyzeWork() {
	return
}
