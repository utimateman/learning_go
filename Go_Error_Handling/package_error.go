// classic error handling / package error handling / adding log file
package main

import (
	"fmt"
	"os"
	"log"
	"github.com/pkg/errors"
)

type Config struct {
	// configuration fields go here (redacted)
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		// [ classic ]
		//return nil, err
		// [ package ] 
		return nil, errors.Wrap(err, "can't open configuration file")
	}

	defer file.Close()

	cfg := &Config{}
	// Parse file here (redacted)
	return cfg, nil

}

// [ add logging ]
func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	log.SetOutput(out)
}
func main() {
	setupLogging()
	cfg, err := readConfig("/path/to/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	} else {
		fmt.Println(cfg)
	}
}