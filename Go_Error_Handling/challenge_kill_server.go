package main

import (
	"fmt"
	"github.com/pkg/errors" //wrap errors
	"io/ioutil" //readfile
	"strconv" //convert file content to integer
	"os"
	"log"
	"strings"
)

func killServer(pidFile string) error {
	data, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "can't open pid file (is server running?)")
	}

	if err := os.Remove(pidFile); err != nil {
		// we can go on if we fail here
		log.Printf("warning: can't remove pid file - %s", err)
	}

	strPID := strings.TrimSpace(string(data))
	pid, err := strconv.Atoi(strPID)
	if err != nil {
		return errors.Wrap(err, "bad process ID")
	}

	// Simulate kill
	fmt.Printf("killing server with pid=%d\n", pid)
	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil { //touch server.pid > hi | 353
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}