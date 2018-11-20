package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/TV4/nids"
)

func main() {
	allowÅÄÖ := flag.Bool("åäö", false, "Allow ÅÄÖ")

	flag.Parse()

	opts := []func(*nids.Nids){}
	if *allowÅÄÖ {
		opts = append(opts, nids.AllowÅÄÖ)
	}

	n := nids.New(opts...)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(n.Case(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
