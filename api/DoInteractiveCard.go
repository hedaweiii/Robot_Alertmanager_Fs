package api

import (
	"Robot_Alertmanager_Fs/internal/app/controllers"
	"context"
	"fmt"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
)

// DoInteractiveCard 处理卡片回调
func DoInteractiveCard(ctx context.Context, data *larkcard.CardAction) (interface{}, error) {
	// 如果回调数据中包含字段

	if data.Action.Value["key"] == "value" {
		option := data.Action.Option
		var time_back int
		var title_back string

		// 根据选项设置静默时间
		switch option {
		case "1":
			time_back = 10
			title_back = "屏蔽10分钟"
		case "2":
			time_back = 30
			title_back = "屏蔽30分钟"
		case "3":
			time_back = 60
			title_back = "屏蔽1小时"
		case "4":
			time_back = 1440
			title_back = "屏蔽24小时"
		default:
			return nil, fmt.Errorf("未知的选项: %s", option)
		}

		// 启动一个后台任务，用来复原卡片
		go controllers.BackMessage(data.OpenMessageID, time_back)

		// 更新卡片
		return controllers.UpdateMessage(data.OpenMessageID, title_back)
	}

	// 返回默认空值
	return nil, nil
}
