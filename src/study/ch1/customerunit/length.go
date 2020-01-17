package customerunit

// 声明千米类型
type KM float64

// 声明米类型
type M float64

// 声明厘米类型
type CM float64

// 进制 小写字母开头 对外不可见
var base float64 = 1000

// km 转换 m
func KM2M(k KM) M {
	return M(k * KM(base))
}

// km 转 cm 小写字母开头 对外不可见
func kM2CM(k KM) CM {
	return CM(k * KM(base*base))
}
