package main

import (
	"fmt"
	"github.com/loveuer/progress"
	"time"
)

func main() {
	p := progress.Start()
	b := progress.NewBar()
	p.AddBar(b)

	for i := 0; ; i++ {
		b.Inc(fmt.Sprintf("(%d MB)", i+1))
		time.Sleep(time.Millisecond * 10)
		if i > 200 {
			break
		}
		if i > 70 {
			p.Stop()
		}
	}

}
