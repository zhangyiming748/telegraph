package constant

import "runtime"

var (
	Root    = "/data" // 工作目录 如果为空  默认/data
	Proxy   = "http://192.168.1.5:8889"
	CpuNums = runtime.NumCPU() // 核心数
)

const (
	MaxCPU = 12
)

func GetCpuNums() int {
	return CpuNums
}

func GetRoot() string {
	return Root
}
func SetRoot(s string) {
	Root = s
}
func GetProxy() string {
	return Proxy
}
func SetProxy(s string) {
	Proxy = s
}
