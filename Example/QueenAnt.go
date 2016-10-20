package main

import (
	"fmt"
	"log"
)

// QueenAnt is Ant's Manager
type QueenAnt struct {
	bot      map[string]*Ant
	addCh    chan *Ant // 새로운 Client 발생
	delCh    chan *Ant // Client 이탈
	reportCh chan bool
	doneCh   chan bool
}

func NewQueenAnt() *QueenAnt {
	bot := make(map[string]*Ant)
	addCh := make(chan *Ant)
	delCh := make(chan *Ant)
	reportCh := make(chan bool)
	doneCh := make(chan bool)

	return &QueenAnt{bot, addCh, delCh, reportCh, doneCh}
}

func (qa *QueenAnt) Add(ant *Ant) {
	qa.addCh <- ant
}

func (qa *QueenAnt) Del(ant *Ant) {
	qa.delCh <- ant
}

func (qa *QueenAnt) GetReport() {
	qa.reportCh <- true
}

func (qa *QueenAnt) Done() {
	qa.doneCh <- true
}

func (qa *QueenAnt) Run() {
	fmt.Println("Run")
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
		// job done
		case <-qa.reportCh:
			// report
			log.Println("Reporting...")
			if len(qa.bot) > 0 {
				for _, ant := range qa.bot {
					data := ant.Report()
					if data != nil {
						log.Println(data.code, " : ", data.value)
					}
				}
			} else {
				log.Println("I've no Ant")
			}
		case <-qa.doneCh:
			// all Ants stop and wait
			return
		}
	}
}
