package backend

import (
	"context"
	"errors"

	"github.com/miniLCT/gochronos/defines"
)

var (
	ErrDataNotFound  = errors.New("backend data not found")
	ErrGraphNotFound = errors.New("backend graph not found")
)

type Backend interface {
	Type() string

	// 设置&获取节点产出数据
	LoadData(ctx context.Context, UUID, key string) (data any, err error)

	StoreData(ctx context.Context, graphName, UUID, key string, data any) error
	StoreDatas(ctx context.Context, graphName, UUID string, datas []defines.Pair[string, any]) error
	LoadKeys(ctx context.Context, UUID string) ([]string, error)

	// 设置节点状态
	SetNodeStatePending(ctx context.Context, UUID, graphName, nodeName string) error
	SetNodeStateSkipped(ctx context.Context, UUID, graphName, nodeName string) error
	SetNodeStateReceive(ctx context.Context, UUID, graphName, nodeName string) error
	SetNodeStateStart(ctx context.Context, UUID, graphName, nodeName string) error
	SetNodeStateWait(ctx context.Context, UUID, graphName, nodeName string) error
	SetNodeStateRetry(ctx context.Context, UUID, graphName, nodeName string) error
	SetNodeStateSuccess(ctx context.Context, UUID, graphName, nodeName string) error
	SetNodeStateFailed(ctx context.Context, UUID, graphName, nodeName, errmsg string) error

	// 获取节点状态
	GetNodeState(ctx context.Context, UUID, nodeName string) (*defines.NodeState, bool) // 单个

	GetNodeStateList(ctx context.Context, UUID string) (map[string]*defines.NodeState, error) // 所有节点状态

	// 设置图状态
	SetGraphStateStart(ctx context.Context, UUID, graphName string) error

	SetGraphStateSuccess(ctx context.Context, UUID, graphName string) error
	SetGraphStateFailed(ctx context.Context, UUID, graphName, errmsg string) error

	// 获取图状态
	GetGraphState(ctx context.Context, UUID, graphName string) (*defines.GraphState, bool)

	// 获取已经运行和全部完成的节点
	GetStartedAndCompletedNode(ctx context.Context, UUID string) ([]string, []string, error)

	// 异步节点失败记录
	SetAsyncNodeStop(ctx context.Context, UUID, graphName, nodeName, errmsg string) error

	// CloneStates复制状态，用于loop模式
	CloneStates(ctx context.Context, graphName, oriUUID string, newUUIDs []string) error
}
