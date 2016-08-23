package jusulsa

import (
	"fmt"
	"time"
)

// Mark1 is mybot
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
}

// AnalyzeWorks analyze QueryData
func (bot *Mark1) AnalyzeWorks() {
	index := len(bot.ObjInfo) - 1
	CurInfo := bot.ObjInfo[index]
	// SB
	if CurInfo.Data.Sell != 0 {
		CurInfo.SBRatio = float32(CurInfo.Data.Buy) / float32(CurInfo.Data.Sell)
	}

	if (CurInfo.Curve > 2) && (CurInfo.SBRatio) > 1.0 {
		SendToU(bot.Code, "BBB", true)
	} else if (CurInfo.Curve < -2) && (CurInfo.SBRatio) < 1.0 {
		SendToU(bot.Code, "SSS", true)
	}
	return
}
