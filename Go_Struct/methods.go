// methods
package main

import (
	"fmt"
)

type Trade struct {
	Symbol string // Stock symbol
	Volume int // Number of shares
	Price  float64 //Trade price
	Buy    bool // true iff buy trade, false is sell trade
}

// Value returns the trade value
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value 
}
func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1.Value())

	
}