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
	return Divide(Sub(data, ma), ma)
}

// BBI 多空指数（Bull and Bear Index）
// 是一种将不同日数移动平均线加权平均之后的综合指标。属于常用技术指标类因子。
func BBI(close []float64) []float64 {
	em5 := EMA(close, 5)
	em10 := EMA(close, 10)
	em20 := EMA(close, 20)
	sum := Add(Add(em5, em10), em20)
	return Multiply(sum, 1.0/3.0)
}

// UPR 股价压力线，对股价有压制作用
// 一般情况下可以取20天的收盘价进行计算，因此close的长度至少为20
// 传入参数每一列为不同交易日的收盘价，每一行表示不同交易日
// 输出为支撑线列
func UPR(close []float64) []float64 {
	bbi := BBI(close)
	result := make([]float64, 0, len(bbi))
	// 针对0-9位置的数据
	for i, b := range bbi {
		if i < 9 {
			result = append(result, 0)
			continue
		}
		// 压力位
		upr := b - 3*Mean(bbi[i-9:i+1])
		result = append(result, upr)
	}
	return result
}

// DWN 股价支撑线，对股价有支撑作用
// 一般情况下可以取10天的收盘价会进行计算
// 传入参数每一列为不同交易日的收盘价，每一行表示不同交易日
// 输出为支撑线列
func DWN(close []float64, n int) []float64 {
	return nil
}

// HHV 输出data在范围n天内的最大值
func HHV(data []float64, n int) []float64 {
	result := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		max := data[i]
		for k := i - 1; k > i - n && k >= 0; k-- {
			if data[k] > max {
				max = data[k]
			}
		}
		result[i] = max
	}
	return result
}