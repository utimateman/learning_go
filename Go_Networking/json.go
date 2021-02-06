// JSON example
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
{
	"user": "Krai Cham",
	"type": "deposit",
	"amount": 10000.3
}
`

// request is a bank transactions
type Request struct {
	Login string `json:"user"`
	Type string `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	rdr := bytes.NewBufferString(data) // simulate a file/socket

	// decode request
	dec := json.NewDecoder(rdr)

	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	fmt.Printf("got: %+v\n", req)

	// create response
	prevBalance := 8500.0 // loaded from db
	resp := map[string]interface{} {
		"ok": true,
		"balance": prevBalance + req.Amount,
	}

	// Encode response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}
}