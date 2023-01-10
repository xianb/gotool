package utils

//FloatProtocal 强化float的扩展方法
type FloatProtocal interface {
	Float32(def ...float32) float32
	Float64(def ...float64) float64
}
