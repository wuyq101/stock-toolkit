package stk

import "math"

// MA 计算MA的工具函数
// 移动平均线，Moving Average，简称MA，MA是用统计分析的方法，将一定时期内的证券价格（指数）加以平均，
// 并把不同时间的平均值连接起来，形成一根MA，用以观察证券价格变动趋势的一种技术指标。
func MA(data []float64, n int) []float64 {
	result := make([]float64, len(data))
	sum := 0.0
	for i := 0; i < len(data); i++ {
		sum += data[i]
		if i >= n {
			sum -= data[i-n]
		}
		if i < n-1 {
			result[i] = 0.0
			continue
		}
		result[i] = sum / float64(n)
	}
	return result
}

// EMA （Exponential Moving Average）是指数平均数指标。
// 也叫EXPMA指标，它也是一种趋向类指标，指数平均数指标是以指数式递减加权的移动平均。
func EMA(data []float64, n int) []float64 {
	result := make([]float64, len(data))
	alpha := 2.0 / float64(n+1)
	result[0] = data[0]
	for i := 1; i < len(data); i++ {
		result[i] = alpha*(data[i]-result[i-1]) + result[i-1]
	}
	return result
}

// FloatEqual 浮点数比较，股票价格等一般精确到小数点之后两位，所以调用的时候，tolerance=0.01
func FloatEqual(a, b float64, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}
