package Colony

import (
	"fmt"
	"log"

	"github.com/nanunsin/jusulsa"
)

type Ant struct {
	qa       *QueenAnt
	code     string
	doneCh   chan bool
	collCh   chan bool
	reportCh chan bool
	bot      *jusulsa.Mark1
}

func NewAnt(code string, qa *QueenAnt) *Ant {

	doneCh := make(chan bool)
	collCh := make(chan bool)
	reportCh := make(chan bool)
	bot := jusulsa.NewBot(code)

	return &Ant{qa, code, doneCh, collCh, reportCh, bot}
}

func (ant *Ant) Collect() {
	ant.collCh <- true
}

func (ant *Ant) collect() {
	ant.bot.QueryWorks()
	lastdata := ant.bot.ObjInfo[len(ant.bot.ObjInfo)-1]
	log.Println(fmt.Sprintf("[%s] %d", ant.code, lastdata.Data.Price))
	ant.bot.AnalyzeWorks(false)
}

func (ant *Ant) report() *AntReport {
	size := len(ant.bot.ObjInfo)
	if size > 0 {
		lastdata := ant.bot.ObjInfo[len(ant.bot.ObjInfo)-1]
		return &AntReport{ant.code, lastdata.Data.Price, lastdata.Volume, lastdata.Curve}
	}
	return nil
}

func (ant *Ant) Report() {
	ant.reportCh <- true
}

func (ant *Ant) Run() {
	for {
		select {
		case <-ant.doneCh:
			ant.qa.Del(ant)
			ant.doneCh <- true
			return
		case <-ant.collCh:
			ant.collect()
		case <-ant.reportCh:
			ant.qa.ReportDataCh <- ant.report()
		}
	}
}
