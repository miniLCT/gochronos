package defines

import "context"

type contextKey int

const (
	NoUse           contextKey = iota // 未使用
	DpipeContextKey                   // dpipe context key
	GraphStateKey                     // 存储graph state
	NodeStateKey                      // 存储node state
	NodeDataKey                       // 存储节点数据
)

type KVStore interface {
	Load(key any) (value any, ok bool)
	Store(key, value any)
	Delete(key any)
	LoadOrStore(key, value any) (actual any, loaded bool)
}

func MustKVStore(ctx context.Context) KVStore {
	if ctx == nil {
		panic("context is nil, can't get kvstore")
	}
	if v := ctx.Value(DpipeContextKey); v != nil {
		return v.(KVStore)
	}
	panic("context has no kvstore interface")
}
