package go_d_text

import (
"time"
)

func Start(element string, ch chan string, p int) chan bool{
	ctrlCh := make(chan bool)
	switch element {
	case "spinner":
			go Spinner(ch, p, ctrlCh)
	default:
	}
	return ctrlCh
}

func Spinner(outCh chan string, p int, ctrlCh chan bool){
	x := "-"
	for {
		moment := time.After(time.Duration(p) * time.Millisecond);

		select{
		case <- ctrlCh:
			outCh <- ""
			return;
		case <- moment:
			outCh <- x
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
		}
	}
}
