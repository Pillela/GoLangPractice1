package Math_Functions

import (
	"fmt"
	"time"
)
	//using <- within func arguments makes a function which can only send values into the channel.
	func Pinger(c chan<- string){
		for i:=0;i<10 ;i++{
			c<- "ping"
		}
	}

	//using <- after the "chan" keyword, within func arguments makes a function which can only send values into the channel.
	func Ponger(c chan<- string){
		for i:=0;i<10 ;i++{
			c<- "pong"
		}
	}
	// using <- before the chan keyword makes a function which can only receive values from the channel.
	func Printer(c <-chan string){
		for {
			fmt.Println(<-c)
			time.Sleep(time.Second * 1)
		}
	}


