package go_d_text

import (
"time"
)

func Start(element string, ch chan string, p int) chan bool{
	controlChannel := make(chan bool)
	switch element {
	case "spinner":
			go Spinner(ch, p, controlChannel)
	default:
	}
	return controlChannel
}

func Spinner(ch chan string, p int, outCh chan bool){
	x := "-"
	for {

		select{
		case <- outCh:
			ch <- ""
			return;
		default:
		}

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
