<<<<<<< HEAD
package jusulsa
=======
package main
>>>>>>> dbcc3dfe2cf8f26f5bf1421ed60bb31e216f33d8

import (
	"fmt"
	"strconv"
	"strings"
)

<<<<<<< HEAD
func removeChar(src, rmchar string) (ret int) {
	src = strings.TrimLeft(src, " ")
	data := strings.Replace(src, ",", "", -1)
	var err error
	if len(data) > 0 {
		ret, err = strconv.Atoi(data)
		if err != nil {
			fmt.Println(fmt.Errorf("[Atoi]%s\n", data))
			panic(err)
		}
=======
func removeChar(src, rmchar string) int {
	src = strings.TrimLeft(src, " ")
	data := strings.Replace(src, ",", "", -1)
	ret, e := strconv.Atoi(data)
	if e != nil {
		fmt.Errorf("[Atoi]%s\n", data)
		panic(e)
>>>>>>> dbcc3dfe2cf8f26f5bf1421ed60bb31e216f33d8
	}
	return ret
}
