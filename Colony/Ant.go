package Colony

import (
	"fmt"
	"log"

	"github.com/nanunsin/jusulsa"
)

type Ant struct {
	qa   *QueenAnt
	code string
	opCh chan uint
	bot  *jusulsa.Mark1
}

func NewAnt(code string, qa *QueenAnt) *Ant {

	opCh := make(chan uint, 5)
	bot := jusulsa.NewBot(code)

	return &Ant{qa, code, opCh, bot}
}

func (ant *Ant) Collect() {
	ant.opCh <- COLLECT
}

func (ant *Ant) Report() {
	ant.opCh <- REPORT
}

func (ant *Ant) Done() {
	ant.opCh <- DONE
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
	log.Println("size = 0")
	return nil
}

func (ant *Ant) Run() {
	for {
		select {
		case opcode := <-ant.opCh:
			switch opcode {
			case COLLECT:
				ant.collect()
			case REPORT:
				ant.qa.ReportDataCh <- ant.report()
			case DONE:
				ant.qa.Del(ant)
				return
			}
		}
	}
}
