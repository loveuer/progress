package main

import (
	"fmt"
	"github.com/lianggaoqiang/progress"
	"time"
)

func main() {
	p := progress.Start()
	b := progress.NewBar()
	p.AddBar(b)

	for i := 0; i <= 100; i++ {
		b.Inc(fmt.Sprintf("(%d MB)", i+1))
		time.Sleep(time.Millisecond * 10)
	}
}
