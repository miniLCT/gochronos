package datahub

import (
	"context"
)

type RestAPI interface {
	// 发布消息，即写入消息
	PublishMessage(ctx context.Context, message string) (*PublishResp, error)

	// 拉取单条，即读取消息，消息不会被消费，使用此方法读取消息时需要业务记录偏移位
	// Fetch(ctx context.Context) (any, error)

	// 消费单条，以竞争模式读取数据，消息消费后需要使用MSGID来调用ACK来标记确认，
	Receive(ctx context.Context) (*ReceiveResp, error)

	// Queue的消息，被标记确认后的消息，不会再通过Receive返回
	ACK(ctx context.Context, msgID string) error
}
