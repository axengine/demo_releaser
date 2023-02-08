// main.go
package main

import (
	"github.com/axengine/utils/log"
	"go.uber.org/zap"
	"time"
)

func main() {
	tk := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-tk.C:
			log.Logger.Info("ticker", zap.Time("now", time.Now()))
		}
	}
}
