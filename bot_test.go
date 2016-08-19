package jusulsa

import (
	"testing"
	"time"
)

func Test_work(t *testing.T) {
	bot := NewBot("015760")
	bot.QueryWorks()
	bot.PrintAll()
}

func Test_worker(t *testing.T) {
	bot := NewBot("015760")
	for i := 0; i < 2; i++ {
		bot.QueryWorks()
		<-time.After(time.Second * 10)
		bot.PrintAt(i)
	}
}
