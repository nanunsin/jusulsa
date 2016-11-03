package Colony

type AntReport struct {
	code   string
	value  int
	volume int
	curve  int
}

const (
	COLLECT = 1
	REPORT  = 2
	DONE    = 3
)
