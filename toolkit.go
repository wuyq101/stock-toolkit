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

// BIAS 乖离率，是用百分比来表示价格与MA间的偏离程度(差距率)。
// BIAS=(收盘价-收盘价的N日简单平均)/收盘价的N日简单平均*100
func BIAS(data []float64, n int) []float64 {
	ma := MA(data, n)
	result := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		if i < n-1 {
			result[i] = 0.0
			continue
		}
		result[i] = (data[i] - ma[i]) / ma[i]
	}
	return result
}

// BIASHIGH 取最高价
func BIASHIGH(high, close []float64, n int) []float64 {
	//TODO
	return nil
}

// BBI 多空指数（Bull and Bear Index）
// 是一种将不同日数移动平均线加权平均之后的综合指标。属于常用技术指标类因子。
func BBI(close []float64) []float64 {
	em5 := EMA(close, 5)
	em10 := EMA(close, 10)
	em20 := EMA(close, 20)
	return Add(Add(em5, em10), em20)
}

// UPR 股价压力线，对股价有压制作用
// 一般情况下可以取10天的收盘价进行计算
// 传入参数每一列为不同交易日的收盘价，每一行表示不同交易日
// 输出为支撑线列
func UPR(close []float64, n int) []float64 {

}

// DWN 股价支撑线，对股价有支撑作用
// 一般情况下可以取10天的收盘价会进行计算
// 传入参数每一列为不同交易日的收盘价，每一行表示不同交易日
// 输出为支撑线列
func DWN(close []float64, n int) []float64 {

}
