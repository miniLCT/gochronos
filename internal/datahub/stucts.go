package datahub

// BpConf bp配置
type BpConf struct {
	PipeName    string `toml:"pipe_name"`
	ServiceName string `toml:"service_name"`
	UserName    string `toml:"user_name"`
	UserPass    string `toml:"user_pass"`
	QueueName   string `toml:"queue_name"`
	QueueToken  string `toml:"queue_token"`
	PipeletID   string `toml:"pipelet_id"`
}

// PipePublishResp 消息发布返回
type PublishResp struct {
	Status  int32  `json:"status"`
	ErrMsg  string `json:"errmsg"`
	Pipelet int32  `json:"pipelet"`
	MsgID   int64  `json:"topic_msgid"`
}

// ReceiveResp 消费消息时的返回
type ReceiveResp struct {
	Status   int32                `json:"status"`
	Messages []ReceiveRespMessage `json:"messages"`
}

// ReceiveRespMessage 消费消息时的消息结构
type ReceiveRespMessage struct {
	Data  string `json:"data"` // 若设置了encoding时 返回的为Base64后的字符串
	Msgid string `json:"msgid"`
}

// ACKResp ack的返回结构
type ACKResp struct {
	Status int `json:"status"`
}
