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

func TestBBI(t *testing.T) {
	t.Skip()
	data := []float64{
		45.06,
		49.57,
		54.53,
		59.98,
		60.7,
		61.41,
		64.97,
		60.76,
		55.7,
		56.78,
		57.55,
		56.66,
		57.89,
		59.3,
		58.26,
		60.5,
		59.87,
		56.84,
		52.83,
		53.86,
		54.25,
		53.83,
		53.68,
		52.54,
		53.41,
		53.88,
		52.96,
	}
	expected := []float64{
		45.06,
		45.977599999999995,
		47.674033333333334,
		50.06163333333334,
		52.01793333333333,
		53.66593333333333,
		55.67246666666667,
		56.371633333333335,
		55.92576666666667,
		55.86583333333334,
		57.54999999999999,
		57.368900000000004,
		57.48356666666666,
		59.29999999999999,
		58.26,
		57.5177,
		57.86606666666666,
		57.53133333333333,
		56.48493333333332,
		55.924333333333344,
		55.58566666666667,
		55.243533333333325,
		54.951766666666664,
		54.49453333333333,
		54.32153333333333,
		54.27766666666667,
		54.04836666666667,
	}

	ema := EMA(data, 12)
	for i, e := range ema {
		fmt.Printf("i: %d, e: %.4f\n", i, e)
	}
	// fmt.Printf("%v\n", EMA(data, 10))
	result := BBI(data)
	// fmt.Printf("bbi:%v\n", result)
	for i := 0; i < len(result); i++ {
		if !FloatEqual(result[i], expected[i], 0.01) {
			t.Fatalf("%d expected %f, but was %f", i, expected[i], result[i])
		}
	}
}
