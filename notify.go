package jusulsa

import (
	"fmt"

	"github.com/jsegura/goxcar"
)

func SendToU(title, msg string, send bool) {
	fmt.Println("send\n" + msg)

	if send {
		goxcar.Notify("GXjKV6MrA2oQHYgknYd", title, msg, "default")
	}
}
