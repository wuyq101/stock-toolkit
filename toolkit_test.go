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
	expected := []float64{
		0.0,
		0.0,
		0.0,
		0.0,
		6.97,
		6.89,
		6.80,
		6.74,
		6.73,
		6.78,
		6.83,
		6.92,
		6.99,
	}
	result := MA(data, 5)
	fmt.Printf("%v\n", result)
	for i := 0; i < len(result); i++ {
		if !FloatEqual(result[i], expected[i], 0.01) {
			t.Fatalf("%d expected %f, but was %f", i, expected[i], result[i])
		}
	}
}

func TestEMA(t *testing.T) {
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
	expected := []float64{
		7.10,
		7.11,
		7.08,
		7.05,
		6.91,
		6.83,
		6.78,
		6.77,
		6.82,
		6.84,
		6.86,
		6.96,
		7.01,
	}
	result := EMA(data, 5)
	fmt.Printf("%v\n", result)
	for i := 0; i < len(result); i++ {
		if !FloatEqual(result[i], expected[i], 0.01) {
			t.Fatalf("%d expected %f, but was %f", i, expected[i], result[i])
		}
	}
}

func TestBIAS(t *testing.T) {
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
	expected := []float64{
		0.0,
		0.0,
		0.0,
		0.0,
		-0.049053,
		-0.030197,
		-0.016466,
		0.000890,
		0.026441,
		0.014450,
		0.012009,
		0.034383,
		0.016586,
	}
	result := BIAS(data, 5)
	fmt.Printf("%v\n", result)
	for i := 0; i < len(result); i++ {
		if !FloatEqual(result[i], expected[i], 0.01) {
			t.Fatalf("%d expected %f, but was %f", i, expected[i], result[i])
		}
	}
}
