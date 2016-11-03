package Colony

import (
	"testing"
	"time"
)

func Test_Queen(t *testing.T) {
	queen := NewQueenAnt()
	go queen.Run()

	time.Sleep(time.Second * 1)

	ant1 := NewAnt("006060", queen)
	queen.Add(ant1)

	time.Sleep(time.Second * 1)

	ant2 := NewAnt("007070", queen)
	queen.Add(ant2)

	queen.Collect()
	time.Sleep(time.Second * 1)
	queen.GetReport()

	t.Log(<-queen.ReportDataCh)
	t.Log(<-queen.ReportDataCh)

	time.Sleep(time.Second * 1)
	queen.Done()
}
