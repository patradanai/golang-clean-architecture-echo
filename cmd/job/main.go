package main

import "flag"

type Args struct {
	Test string
}

func main() {
	var opts Args

	flag.StringVar(&opts.Test, "flag", "default", "flag description")

	flag.Parse()

	// TODO: implement
}
