package main

import "flag"

var (
	panelCount int
)

func getFlags() {
	flag.IntVar(&panelCount, "panels", 1, "number of panels")

	flag.Parse()
}
