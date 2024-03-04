package main

import (
	"flag"

	"github.com/kubeclipper/kubeclipper/tools/kcctldocs-gen/generators"
)

func main() {
	/*
		1.generate markdown files.
		2.use docker image brianpursley/brodocs to generate html files.
	*/
	flag.Parse()
	generators.GenerateFiles()
}
