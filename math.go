package stk

import "math"

// Divide 列运算，两列相除
func Divide(x, y []float64) []float64 {
	return forEach(len(x), func(i int) float64 {
		if y[i] == 0.0 {
			return 0.0
		}
		return x[i] / y[i]
	})
}

// Sub 列运算，两列相减
func Sub(x, y []float64) []float64 {
	return forEach(len(x), func(i int) float64 { return x[i] - y[i] })
}

// Add 列运算，两列相加
func Add(x, y []float64) []float64 {
	return forEach(len(x), func(i int) float64 { return x[i] + y[i] })
}

// Multiply 列乘以一个系数
func Multiply(x []float64, a float64) []float64 {
	return forEach(len(x), func(i int) float64 { return x[i] * a })
}

func forEach(length int, fn func(int) float64) []float64 {
	result := make([]float64, length)
	for i := 0; i < length; i++ {
		result[i] = fn(i)
	}
	return result
}

// ColumnSum 列求和
func ColumnSum(x []float64) float64 {
	r := 0.0
	for _, num := range x {
		r += num
	}
	return r
}

// Mean 求平均值
func Mean(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	sum := ColumnSum(x)
	return sum / float64(len(x))
}

// Std 求标准差
func Std(x []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	mean := Mean(x)
	size := len(x)
	var sd float64
	for _, num := range x {
		sd += math.Pow(num-mean, 2)
	}
	return math.Sqrt(sd / float64(size))
}
