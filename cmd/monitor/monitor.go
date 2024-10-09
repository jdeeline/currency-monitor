package main

import (
	"flag"

	"github.com/jdeeline/currency-monitor/internal/monitor"
)

func main() {
	var interval int
	flag.IntVar(&interval, "i", 60, "an update interval in minutes")
	flag.Parse()

	monitor.Run(interval)
}
