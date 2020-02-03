package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var args Arguments
	flag.StringVar(&args.Network, "network", "tcp", "network to dial, default is tcp.")
	flag.StringVar(&args.Host, "host", "", "host to dial, e.g. 127.0.0.1")
	flag.IntVar(&args.Port, "port", 0, "port to dial, e.g. 80")
	flag.StringVar(&args.Timeout, "timeout", "10s", "dial timeout, default is 10s")
	flag.IntVar(&args.Retry, "retry", 0, "times to retry, default is 0 - no retry")
	flag.StringVar(&args.Delay, "delay", "1s", "duration to delay retry, default is 1s")
	flag.BoolVar(&args.Verbose, "verbose", false, "default is off")
	flag.Parse()

	if err := args.Parse(); err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}
	address := args.GetAddress()

	o := NewOutputer(args.Verbose)

retry:
	o.Printf("dialing to '%s'\n", address)
	conn, err := net.DialTimeout(args.Network, address, args.GetTimeout())
	if err != nil {
		o.Printf("failed to dial to '%s': %v\n", args.GetAddress(), err)
		if args.Retry > 0 {
			args.Retry--
			time.Sleep(args.GetDelay())
			goto retry
		}
		os.Exit(1)
	}

	o.Printf("dial to '%s' successfully\n", address)
	conn.Close()
}
