package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	robot_Clicent "Robot_Alertmanager_Fs/config"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// 恢复json，需要传入消息id和倒计时时间
func BackMessage(MessageId string, time_back int) (string, error) {
	time.Sleep(time.Duration(time_back) * time.Minute)
	// 根据 MessageId 构造文件路径
	filePath := fmt.Sprintf("message/%s", MessageId)

	// 读取文件内容
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	card := string(bs)
	// 构造请求
	req := larkim.NewPatchMessageReqBuilder().
		MessageId(MessageId).
		Body(larkim.NewPatchMessageReqBodyBuilder().
			Content(card). // 使用读取的内容
			Build()).
		Build()

	// 发起请求
	resp, err := robot_Clicent.Client.Im.Message.Patch(context.Background(), req)
	if err != nil {
		return "", err
	}
	if !resp.Success() {
		return "", resp.CodeError
	}

	return "ok", nil

}
