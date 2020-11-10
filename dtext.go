package go_d_text

import (
"time"
)

func spiner(ch chan string, p int) {
	x := "-"
	for {
		ch <- x
		switch x {
		case "-":
			x = "\\"
		case "\\":
			x = "|"
		case "|":
			x = "/"
		case "/":
			x = "-"
		}
		time.Sleep(time.Duration(p) * time.Millisecond)
	}
}
