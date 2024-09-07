//go:build linux
// +build linux

package rotate

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Example of how to rotate in response to SIGHUP.
func ExampleLoggerRotate() {
	l := &Writer{}
	log.SetOutput(l)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for {
			<-c
			l.Rotate()
		}
	}()
}
