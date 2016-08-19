package jusulsa

import (
	"fmt"
	"time"
)

func getCurTimeString() (ret string) {
	h, m, s := time.Now().Round(time.Second * 10).Clock()
	//time.Now().Round(time.Second*10)
	ret = fmt.Sprintf("%02d%02d%02d", h, m, s)
	return ret
}
