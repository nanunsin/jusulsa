package jusulsa

import (
	"testing"
	"time"
)

func Test_data_1(t *testing.T) {
	bot := NewBot("015760")
	for i := 0; i < 2; i++ {
		bot.QueryWorks()
		bot.PrintAt(i)
		<-time.After(time.Second * 10)
	}

	WriteCSV(bot, "015760.csv")
}
