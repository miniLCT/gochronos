package defines

import "encoding/json"

// 每个算子节点的运行状态
const (
	NodeStatePending = "PENDING"  // 任务被调度器调度，等待worker接受
	NodeStateReceive = "RECEIVED" // 任务已经被worker接受到，等待执行
	NodeStateStart   = "STARTED"  // 任务已经在wokrer节点上开始执行
	NodeStateWait    = "WAIT"     // 任务已经在wokrer节点上执行完毕，等待异步回调
	NodeStateRetry   = "RETRY"    // 任务执行失败，需要重试, 后续可以加入延迟队列相关的功能，重试任务
	NodeStateSkipped = "SKIPPED"  // 任务被跳过，不需要执行
	NodeStateSuccess = "SUCCESS"  // 任务执行成功
	NodeStateFailure = "FAILURE"  // 任务执行失败
)

const (
	GraphStateStart   = "START"   // 图开始执行
	GraphStateSuccess = "SUCCESS" // 图执行成功
	GraphStateError   = "ERROR"   // 图执行失败,但是可以继续执行其他节点
	GraphStateFailure = "FAILURE" // 图执行失败,并且停止执行其他节点
)

// NodeState 节点的状态
type NodeState struct {
	ID         uint        `json:"_id,omitempty" ddb:"id"`
	GraphName  string      `json:"graph_name,omitempty" ddb:"graph_name"`
	NodeName   string      `json:"node_name,omitempty" ddb:"node_name"`
	UUID       string      `json:"uuid,omitempty" ddb:"uuid"`
	State      string      `json:"state" ddb:"state"`
	Error      string      `json:"error,omitempty" ddb:"error"`
	CreateAt   string      `json:"create_at,omitempty" ddb:"create_at"`
	CreateAtMs int64       `json:"create_at_ms" ddb:"-"`
	CostMs     int64       `json:"cost_ms" ddb:"-"`
	StateHis   []NodeState `json:"state_history,omitempty" ddb:"-"`
}

func (ns *NodeState) String() string {
	bs, _ := json.Marshal(ns)
	return string(bs)
}

// GraphState 整个图的状态
type GraphState struct {
	ID        uint         `json:"id" ddb:"id"`
	GraphName string       `json:"graph_name" ddb:"graph_name"`
	UUID      string       `json:"uuid" ddb:"uuid"`
	State     string       `json:"state" ddb:"state"`
	Error     string       `json:"error" ddb:"error"`
	CreateAt  string       `json:"create_at" ddb:"create_at"`
	StateHis  []GraphState `json:"state_his,omitempty" ddb:"-"`
}

func (gs *GraphState) String() string {
	bs, _ := json.Marshal(gs)
	return string(bs)
}

type AllStates struct {
	GraphState *GraphState           `json:"graph_state"`
	NodeStates map[string]*NodeState `json:"node_states"`
	NodeDatas  map[string]any        `json:"node_datas"`
}

func (as *AllStates) String() string {
	bs, _ := json.Marshal(as)
	return string(bs)
}

func (gs *GraphState) IsSuccess() bool {
	return gs.State == GraphStateSuccess
}

func (gs *GraphState) IsFailure() bool {
	return gs.State == GraphStateFailure
}

func (ns *NodeState) IsCompleted() bool {
	return ns.State == NodeStateSuccess ||
		ns.State == NodeStateFailure ||
		ns.State == NodeStateSkipped
}

func (ns *NodeState) IsSuccess() bool {
	return ns.State == NodeStateSuccess
}

func (ns *NodeState) IsFailure() bool {
	return ns.State == NodeStateFailure
}
