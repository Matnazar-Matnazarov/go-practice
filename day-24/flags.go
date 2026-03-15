package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	port    = flag.Int("port", 8080, "HTTP server port")
	verbose = flag.Bool("v", false, "verbose output")
	name    = flag.String("name", "Mehmon", "foydalanuvchi ismi")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Kun 24 CLI demo.\nIshlatish: %s [options] [args...]\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

func flagArgs() []string {
	return flag.Args()
}
