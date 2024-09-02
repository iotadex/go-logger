package main

import (
	"log/slog"
	"sync"
	"time"

	"github.com/triplefi/go-logger/wangyi"
)

func main() {

	slog.SetDefault(wangyi.GetDefaultDailyRotatedLogger("logs/out.log", 0, 0))

	slog.Debug("Debug", "bug", 100000)

	var w sync.WaitGroup
	w.Add(10)
	lines := 100000
	for i := 0; i < 10; i++ {
		go func(m int) {
			max := m + lines
			for n := m; n < max; n++ {
				slog.Info("Hello", "value", n)
				time.Sleep(time.Second)
			}
			w.Done()
		}(i * lines)
	}
	w.Wait()
}
