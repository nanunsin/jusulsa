package main

import (
	"bufio"
	"os"
	"runtime"
	"time"
)

const (
	Stopped = 0
	Paused  = 1
	Running = 2
)

func Query(qa *QueenAnt) {
	for {
		if isCont() {
			runtime.Gosched()
			qa.GetReport()
			<-time.After(time.Second * 10)
		} else {
			return
		}
	}
}

func main() {

	userInput := bufio.NewScanner(os.Stdin)
	qa := NewQueenAnt()
	go qa.Run()
	//	go Query(qa)

	for userInput.Scan() {
		inputData := userInput.Text()
		if len(inputData) == 6 {
			ant := NewAnt(inputData, qa)
			qa.Add(ant)
		} else if "exit" == inputData {
			break
		}
	}
}
