package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/nanunsin/jusulsa"
)

type Ant struct {
	qa    *QueenAnt
	code  string
	state chan int
	bot   *jusulsa.Mark1
}

func NewAnt(code string, qa *QueenAnt) *Ant {

	state := make(chan int)
	bot := jusulsa.NewBot(code)

	return &Ant{qa, code, state, bot}
}

func isCont() bool {
	now := time.Now()

	target := time.Date(now.Year(), now.Month(), now.Day(), 15, 30, 0, 0, now.Location())
	diff := target.Sub(now)
	if diff.Seconds() > 0 {
		return true
	}
	return false
}

func DoWork(bot *jusulsa.Mark1) {
	bot.QueryWorks()
	// Print
	bot.AnalyzeWorks(false)
}

func (ant *Ant) Run() {
	state := Paused
	for {
		select {
		case state = <-ant.state:
			switch state {
			case Stopped:
				return
			case Running:
				fmt.Println("Run")
			case Paused:
				fmt.Println("Paused")
			}
		default:
			runtime.Gosched()

			if state == Paused {
				break
			}

			// do work
			if isCont() {
				DoWork(ant.bot)
				<-time.After(time.Second * 10)
			} else {

				if len(ant.bot.ObjInfo) > 0 {
					fmt.Println("SaveFiles")
					y, m, d := time.Now().Date()
					csvfilename := fmt.Sprintf("%s_%04d%02d%02d.csv", ant.bot.Code, y, m, d)
					jusulsa.WriteCSV(ant.bot, csvfilename)
				}
				return
			}
		}
	}
}
