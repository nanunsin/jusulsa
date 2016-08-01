package jusulsa

import (
	"fmt"
	"strconv"
	"strings"
)

func removeChar(src, rmchar string) (ret int) {
	src = strings.TrimLeft(src, " ")
	data := strings.Replace(src, ",", "", -1)
	
	ret = 0
	var err error
	if len(data) > 0 {
		ret, err = strconv.Atoi(data)
		if err != nil {
			fmt.Println(fmt.Errorf("[Atoi]%s\n", data))
			panic(err)
		}
	}
	return ret
}
