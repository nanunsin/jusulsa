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
	for i := 0; i < 5; i++ {
		bot.QueryWorks()
		bot.PrintAt(i)
		<-time.After(time.Second * 3)
	}
}
