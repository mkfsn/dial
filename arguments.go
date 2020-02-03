package main

import (
	"fmt"
	"time"
)

type Arguments struct {
	Network string // default tcp
	Host    string
	Port    int
	Timeout string // default 10s
	Retry   int    // default 0
	Delay   string // default 1s
	Verbose bool   // default is off

	parsedTimeout time.Duration
	parsedDelay   time.Duration
}

func (args *Arguments) GetAddress() string {
	return fmt.Sprintf("%s:%d", args.Host, args.Port)
}

func (args *Arguments) GetTimeout() time.Duration {
	return args.parsedTimeout
}

func (args *Arguments) GetDelay() time.Duration {
	return args.parsedDelay
}

func (args *Arguments) Parse() error {
	if args.Host == "" {
		return fmt.Errorf("empty host")
	}

	if args.Port == 0 {
		return fmt.Errorf("empty port")
	}

	timeout, err := time.ParseDuration(args.Timeout)
	if err != nil {
		return fmt.Errorf("failed to parse timeout '%s': %w\n", args.Timeout, err)
	}
	args.parsedTimeout = timeout

	delay, err := time.ParseDuration(args.Delay)
	if err != nil {
		return fmt.Errorf("failed to parse delay '%s': %w\n", args.Delay, err)
	}
	args.parsedDelay = delay

	return nil
}
