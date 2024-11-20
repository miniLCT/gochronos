package defines

import "context"

// Operator 算子需要实现的接口
type Operator interface {
	SetConfig(map[string]string) error              // 设置算子配置，服务启动时设置一次
	Reset()                                         // 重置算子状态，每次请求都会执行一次
	Process(context.Context) (ProcessStatus, error) // 执行算子具体逻辑
}

type Caller interface {
	CallbackHandler(ctx context.Context) error
	SetBackParam(ctx context.Context, data []byte, contentType string) error
}
