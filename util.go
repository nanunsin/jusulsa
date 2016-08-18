package jusulsa

import (
	"fmt"
	"time"
)

func getCurTimeString() (ret string) {
	h, m, s := time.Now().Clock()
	ret = fmt.Sprintf("%02d%02d%02d", h, m, s)
	return ret
}
