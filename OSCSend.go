package main

import (
	"fmt"
	"time"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
	client := osc.NewClient("127.0.0.1", 9000)
	msgMin := osc.NewMessage("/avatar/parameters/LeftMin")
	msgHour := osc.NewMessage("/avatar/parameters/RightHour")
	for {
		var t = time.Now()
		hour := t.Hour()
		minute := t.Minute()
		fmt.Print(t.Local().Hour(), t.Local().Minute(), "\n")
		fminute := float32(minute) / float32(60)
		fhour := float32(int(hour%12))/float32(12) + fminute/float32(12)
		fmt.Print(fhour, fminute, "\n")
		msgMin.Append(float32(fminute))
		msgHour.Append(float32(fhour))
		client.Send(msgHour)
		client.Send(msgMin)
		time.Sleep(time.Second)
	}
}
