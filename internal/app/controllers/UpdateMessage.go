package controllers

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func UpdateMessage(MessageId string, title string) (string, error) {
	// 读取根据 MessageId 构造的文件路径
	filePath := fmt.Sprintf("message/%s", MessageId)

	// 读取文件内容
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	card := string(bs)
	card = strings.Replace(card, "系统触发告警", title, -1)
	card = strings.Replace(card, "red", "grey", -1)
	card = strings.Replace(card, "blue", "grey", -1)
	card = strings.Replace(card, "turquoise", "grey", -1)

	return card, nil
}
