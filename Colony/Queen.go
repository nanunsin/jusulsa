package main

import (
	"fmt"
	"log"
)

// QueenAnt is Ant's Manager
type QueenAnt struct {
	bot          map[string]*Ant
	addCh        chan *Ant // 새로운 Client 발생
	delCh        chan *Ant // Client 이탈
	collCh       chan bool
	reportCh     chan bool
	doneCh       chan bool
	ReportDataCh chan *AntReport
}

// NewQueenAnt Create New Queen
func NewQueenAnt() *QueenAnt {
	bot := make(map[string]*Ant)
	addCh := make(chan *Ant)
	delCh := make(chan *Ant)
	collCh := make(chan bool)
	reportCh := make(chan bool)
	doneCh := make(chan bool)
	ReportDataCh := make(chan *AntReport, 10)

	return &QueenAnt{bot, addCh, delCh, collCh, reportCh, doneCh, ReportDataCh}
}

func (qa *QueenAnt) Add(ant *Ant) {
	qa.addCh <- ant
}

func (qa *QueenAnt) Del(ant *Ant) {
	qa.delCh <- ant
}

func (qa *QueenAnt) Collect() {
	qa.collCh <- true
}

func (qa *QueenAnt) GetReport() {
	qa.reportCh <- true
}

func (qa *QueenAnt) Done() {
	qa.doneCh <- true
}

func (qa *QueenAnt) Run() {
	fmt.Println("Queen Run")
	for {
		select {
		// Add Ant
		case ant := <-qa.addCh:
			log.Println("Added new ant")
			qa.bot[ant.code] = ant
			log.Println("Now", len(qa.bot), "ants spawned.")
			go ant.Run()
		// Delete Ant
		case ant := <-qa.delCh:
			log.Println("Delete ant")
			delete(qa.bot, ant.code)

		case <-qa.collCh:
			// report
			log.Println("Collect...")
			if len(qa.bot) > 0 {
				for _, ant := range qa.bot {
					ant.Collect()
				}
			} else {
				log.Println("I've no Ant")
			}

		case <-qa.reportCh:
			// report
			log.Println("Reporting...")
			if len(qa.bot) > 0 {
				for _, ant := range qa.bot {
					ant.Report()
				}
			} else {
				log.Println("I've no Ant")
			}
		// job done
		case <-qa.doneCh:
			// all Ants stop and wait
			return
		}
	}
}
