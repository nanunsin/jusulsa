package jusulsa

import (
	"fmt"
	"time"
)

type Mark1 struct {
	Today   string
	ObjInfo []QryInfo
}

func (bot *Mark1) Print() {
	for i := 0; i < len(bot.ObjInfo); i++ {
		fmt.Printf("%s\t", bot.ObjInfo[i].TimeStr)
		fmt.Printf("%d(%d)\t", bot.ObjInfo[i].Data.Price, bot.ObjInfo[i].Volume)
		fmt.Printf("%d\n", bot.ObjInfo[i].Curve)
	}
}

func Works(code string) {
	fmt.Println(code + " work start")
	var bot Mark1
	for i := 0; i < 5; i++ {
		bot.ObjInfo = append(bot.ObjInfo, *QueryData(code))
		if i == 0 {
			makeInfoStep1(nil, &bot.ObjInfo[0])
		} else {
			makeInfoStep1(&bot.ObjInfo[i-1], &bot.ObjInfo[i])
		}

		<-time.After(time.Second * 5)
	}
	bot.Print()
}
