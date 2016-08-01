package main

import (
	"fmt"
	"strconv"
	"strings"
)

func removeChar(src, rmchar string) int {
	src = strings.TrimLeft(src, " ")
	data := strings.Replace(src, ",", "", -1)
	ret, e := strconv.Atoi(data)
	if e != nil {
		fmt.Errorf("[Atoi]%s\n", data)
		panic(e)
	}
	return ret
}
