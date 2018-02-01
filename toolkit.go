package stk

import (
	"math"
)

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

// MACD 称为指数平滑移动平均线，是从双指数移动平均线发展而来的，由快的指数移动平均线（EMA12）减去慢的指数移动平均线（EMA26）得到快线DIF，
// 再用2×（快线DIF-DIF的9日加权移动均线DEA）得到MACD柱。MACD的意义和双移动平均线基本相同，即由快、慢均线的离散、聚合表征当前的多空状态和股价可能的发展变化趋势，
// 但阅读起来更方便。当MACD从负数转向正数，是买的信号。当MACD从正数转向负数，是卖的信号。
// 当MACD以大角度变化，表示快的移动平均线和慢的移动平均线的差距非常迅速的拉开，代表了一个市场大趋势的转变。
func MACD(data []float64, ema1, ema2, dif int) []float64 {
	e1 := EMA(data, ema1)
	e2 := EMA(data, ema2)
	d := Sub(e1, e2)
	dea := EMA(d, dif)
	return Multiply(Sub(d, dea), 2.0)
}

// FloatEqual 浮点数比较，股票价格等一般精确到小数点之后两位，所以调用的时候，tolerance=0.01
func FloatEqual(a, b float64, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

// BIAS 乖离率，是用百分比来表示价格与MA间的偏离程度(差距率)。
// BIAS=(收盘价-收盘价的N日简单平均)/收盘价的N日简单平均*100
func BIAS(data []float64, n int) []float64 {
	ma := EMA(data, n)
	return Divide(Sub(data, ma), ma)
}

// BBI 多空指数（Bull and Bear Index）
// 是一种将不同日数移动平均线加权平均之后的综合指标。属于常用技术指标类因子。
// 默认传入为5日、10日和20日
func BBI(close []float64, n ...int) []float64 {
	result := make([]float64, len(close))
	for _, v := range n {
		result = Add(result, EMA(close, v))
	}
	return Multiply(result, 1.0/float64(len(n)))
}

// UPR 股价压力线，对股价有压制作用
// 一般情况下可以取20天的收盘价进行计算，因此close的长度至少为20
// 传入参数每一列为不同交易日的收盘价，每一行表示不同交易日
// 输出为支撑线列
// n 表示向前查找的bbi天数
// m 表示彼标准差常量参数
// 推荐 n=10，m=3.0
func UPR(close []float64, n int, m float64) []float64 {
	bbi := BBI(close, 5, 10, 20)
	return Add(bbi, Multiply(STD(bbi, n), m))
}

// DWN 股价支撑线，对股价有支撑作用
// 一般情况下可以取10天的收盘价会进行计算
// 传入参数每一列为不同交易日的收盘价，每一行表示不同交易日
// 输出为支撑线列
// n 表示向前查找的bbi天数
// m 表示彼标准差常量参数
// 推荐 n=10，m=-3.0
func DWN(close []float64, n int, m float64) []float64 {
	bbi := BBI(close)
	return Add(bbi, Multiply(STD(bbi, n), m))
}

// STD 求标准差
func STD(x []float64, n int) []float64 {
	return forEach(n, func(i int) float64 {
		if i < n-1 {
			return 0.0
		}
		return Std(x[i-n+1 : i+1])
	})
}

// HHV 输出data在范围n天内的最大值
func HHV(data []float64, n int) []float64 {
	result := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		max := data[i]
		for k := i - 1; k > i-n && k >= 0; k-- {
			if data[k] > max {
				max = data[k]
			}
		}
		result[i] = max
	}
	return result
}
