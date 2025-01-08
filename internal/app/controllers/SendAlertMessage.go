package controllers

import (
	robot_Clicent "Robot_Alertmanager_Fs/config"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// SendAlertMessage 发送报警消息
func SendAlertMessage(chatId string, content string) error {

	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType("chat_id").
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(chatId).
			MsgType("interactive").
			Content(content).
			Build()).
		Build()

	resp, _ := robot_Clicent.Client.Im.Message.Create(context.Background(), req)
	fmt.Println(resp)
	if !resp.Success() {
		return resp.CodeError
	}
	resp_s := larkcore.Prettify(resp)
	// 使用正则表达式提取 message_id

	re := regexp.MustCompile(`MessageId:\s*"(.*?)"`)
	matches := re.FindStringSubmatch(resp_s)
	messageID := matches[1]

	// 构建文件路径
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("获取当前工作目录失败: %v", err)
	}
	filePath := filepath.Join(wd+"/../../resources/messages", messageID)
	// 将字符串内容写入文件
	err1 := os.WriteFile(filePath, []byte(content), 0644)
	if err1 != nil {
		return fmt.Errorf("写入文件失败: %v", err1)
	}
	return nil
}
