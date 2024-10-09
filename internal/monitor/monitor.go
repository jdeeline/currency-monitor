package monitor

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jdeeline/currency-monitor/internal/grabber"
	"github.com/jdeeline/currency-monitor/internal/render"
)

// Run a currency monitor.
func Run(interval int) {
	update()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)

	for {
		select {
		case <-quit:
			return
		case <-time.After(time.Duration(interval) * time.Minute):
			update()
		}
	}
}

func update() {
	currencies, err := grabber.Grab()
	if err != nil {
		fmt.Println(err)
		return
	}

	render.Table(currencies)
}
