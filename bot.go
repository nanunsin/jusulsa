package jusulsa

import (
	"fmt"
	"time"
)

type Average struct {
	Price  int
	Volume int
}

// Mark1 is mybot
type Mark1 struct {
	Date    string
	ObjInfo []QryInfo
	Code    string
	Avg     Average
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

// PrintAll print all objInfo data
func (bot *Mark1) PrintAll() {
	for i := 0; i < len(bot.ObjInfo); i++ {
		printTimeStamp(bot.ObjInfo[i].TimeStamp, 10)
		fmt.Printf("\t%d(%d)", bot.ObjInfo[i].Data.Price, bot.ObjInfo[i].Volume)
		fmt.Printf("\t%d", bot.ObjInfo[i].Curve)
		fmt.Println()
	}
}

// PrintAt print objInfo[index] data
func (bot *Mark1) PrintAt(index int) {
	if len(bot.ObjInfo) >= index {
		printTimeStamp(bot.ObjInfo[index].TimeStamp, 10)
		fmt.Printf("\t%d(%d)", bot.ObjInfo[index].Data.Price, bot.ObjInfo[index].Volume)
		fmt.Printf("\t%d", bot.ObjInfo[index].Curve)
		fmt.Printf("\t[%d|%d(%f)]", bot.ObjInfo[index].Data.Sell, bot.ObjInfo[index].Data.Buy, bot.ObjInfo[index].SBRatio)
		fmt.Printf("\t%d", bot.Avg.Volume)
		fmt.Println()
	}
}

// QueryWorks collect data and add list
// some works
func (bot *Mark1) QueryWorks() {
	bot.ObjInfo = append(bot.ObjInfo, *QueryData(bot.Code))
	size := len(bot.ObjInfo) - 1

	if size == 0 {
		makeInfoStep1(nil, &bot.ObjInfo[0])
	} else {
		makeInfoStep1(&bot.ObjInfo[size-1], &bot.ObjInfo[size])
	}

	bot.ObjInfo[size].Index = getTimeStampIndex()
	if bot.ObjInfo[size].Index == 0 {
		bot.Avg.Volume = bot.ObjInfo[size].Volume
	} else {
		bot.Avg.Volume = (bot.ObjInfo[size].Data.TotalVolume / bot.ObjInfo[size].Index)
	}
}

// AnalyzeWorks analyze QueryData
func (bot *Mark1) AnalyzeWorks(alarm bool) {
	index := len(bot.ObjInfo) - 1
	CurInfo := bot.ObjInfo[index]

	if (CurInfo.Curve > 2) && (CurInfo.SBRatio) > 1.0 {
		title := fmt.Sprintf("%s-BBB", bot.Code)
		msg := fmt.Sprintf("%d  %f", CurInfo.Data.Price, CurInfo.SBRatio)
		SendToU(title, msg, alarm)
	} else if (CurInfo.Curve < -2) && (CurInfo.SBRatio) < 1.0 {
		title := fmt.Sprintf("%s-SSS", bot.Code)
		msg := fmt.Sprintf("%d  %f", CurInfo.Data.Price, CurInfo.SBRatio)
		SendToU(title, msg, alarm)
	}
	return
}
