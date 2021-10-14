package mymq

/**
实验打点数据上传至数据中心
一般使用mq
此处为mq mock
*/

import (
	"context"
	"fmt"
)

type message struct {
	key   string
	value string
}

var mq = make(chan *message)

func SendMessage(key, value string) {
	go func() {
		mq <- &message{
			key:   key,
			value: value,
		}
	}()
}

func Produce(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-mq:
				fmt.Println("send a msg:", string(msg.key), string(msg.value))
			}
		}
	}()
}
