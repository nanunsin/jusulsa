package main

import (
	"sync"
	"time"
)

const (
	Stopped = 0
	Paused  = 1
	Running = 2
)

func main() {

	var wg sync.WaitGroup

	code := "006060"
	bots := make(map[string]*Ant)
	wg.Add(1)
	bots[code] = NewAnt(code)
	go bots[code].Run()
	bots[code].state <- Running

	<-time.After(time.Second * 20)
	wg.Done()

	wg.Wait()

}
