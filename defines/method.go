package defines

import "context"

// NodeMethodXXX 个性化的直接在node上绑定方法，通过接口断言实现

// NodeMethodErrHandler node 错误处理实现接口
type NodeMethodErrHandler interface {
	ErrorHandler(ctx context.Context, graphName string, UUID string, err error)
}
