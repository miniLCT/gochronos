package defines

type ProcessStatus int

const (
	ProcessStatusContinue ProcessStatus = iota // 同步请求继续执行
	ProcessStatusStop                          // 停止整张图的调度
	ProcessStatusWait                          // 异步请求等待
)
