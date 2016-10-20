package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/nanunsin/jusulsa"
)

type Ant struct {
	qa       *QueenAnt
	code     string
	doneCh   chan bool
	reportCh chan bool
	bot      *jusulsa.Mark1
}

func NewAnt(code string, qa *QueenAnt) *Ant {

	doneCh := make(chan bool)
	reportCh := make(chan bool)
	bot := jusulsa.NewBot(code)

	return &Ant{qa, code, doneCh, reportCh, bot}
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

func (ant *Ant) SaveBotData() {
	if len(ant.bot.ObjInfo) > 0 {
		fmt.Println("SaveFiles")
		y, m, d := time.Now().Date()
		csvfilename := fmt.Sprintf("%s_%04d%02d%02d.csv", ant.bot.Code, y, m, d)
		jusulsa.WriteCSV(ant.bot, csvfilename)
	}
}

func (ant *Ant) Collect() {
	ant.bot.QueryWorks()
	lastdata := ant.bot.ObjInfo[len(ant.bot.ObjInfo)-1]
	log.Println(fmt.Sprintf("[%s] %d", ant.code, lastdata.Data.Price))
	ant.bot.AnalyzeWorks(false)
}

func (ant *Ant) Listen() {
	for {
		select {
		case <-ant.doneCh:
			ant.qa.Del(ant)
			ant.doneCh <- true
			return
		case <-ant.reportCh:
			ant.Report()
		}
	}
}

func (ant *Ant) Report() *AntReport {
	size := len(ant.bot.ObjInfo)
	if size > 0 {
		lastdata := ant.bot.ObjInfo[len(ant.bot.ObjInfo)-1]
		return &AntReport{ant.code, lastdata.Data.Price, lastdata.Volume, lastdata.Curve}
	}
	return nil
}

func (ant *Ant) Run() {

	go ant.Listen()

	// Collect
	for {
		select {
		case <-ant.doneCh:
			ant.qa.Del(ant)
			ant.doneCh <- true
			return
		default:
			runtime.Gosched()
			if isCont() {
				ant.Collect()
				<-time.After(time.Second * 10)
			} else {
				ant.SaveBotData()
				ant.doneCh <- true
				return
			}
		}
	}
}

/*
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
				ant.SaveBotData()

				}
				return
			}
		}
	}
}
*/
