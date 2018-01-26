package stk

// Divide 列运算，两列相除
func Divide(x, y []float64) []float64 {
	return forEach(x, y, func(i int) float64 {
		if y[i] == 0.0 {
			return 0.0
		}
		return x[i] / y[i]
	})
}

// Sub 列运算，两列相减
func Sub(x, y []float64) []float64 {
	return forEach(x, y, func(i int) float64 { return x[i] - y[i] })
}

// Multiply 列乘以一个系数
func Multiply(x []float64, a float64) []float64 {
	return forEach(x, nil, func(i int) float64 { return x[i] * a })
}

func forEach(x, y []float64, fn func(int) float64) []float64 {
	result := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		result[i] = fn(i)
	}
	return result

}
