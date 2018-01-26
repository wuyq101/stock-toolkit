package stk

import (
	"fmt"
	"testing"
)

func TestMA(t *testing.T) {
	data := []float64{
		7.1000,
		7.1200,
		7.0400,
		6.9700,
		6.6300,
		6.6800,
		6.6900,
		6.7500,
		6.9100,
		6.8800,
		6.9100,
		7.1600,
		7.1100,
	}
	//	expected := []float64{}
	result := MA(data, 5)
	fmt.Printf("%v\n", result)
}
