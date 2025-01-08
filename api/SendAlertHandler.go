package api

import (
	"Robot_Alertmanager_Fs/internal/app/controllers"
	"Robot_Alertmanager_Fs/pkg/tools"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Alert 表示报警的结构
type Alert struct {
	Annotations map[string]string `json:"annotations"`
	StartsAt    string            `json:"startsAt"`
	EndsAt      string            `json:"endsAt"`
	Lables      map[string]string `json:"labels"`
}

// AlertGroup 表示报警组的结构
type AlertGroup struct {
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	Alerts            []Alert           `json:"alerts"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
}

// sendAlertHandler 处理发送报警消息的 HTTP POST 请求
func SendAlertHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	// 检查请求方法是否为 POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从 URL 查询参数中提取 chatId
	chatId := r.URL.Query().Get("chatId")
	if chatId == "" {
		http.Error(w, "chatId is required", http.StatusBadRequest)
		return
	}

	// 读取请求体并打印
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close() // 关闭请求体

	//开始处理格式化alermanager告警信息
	var alertGroup AlertGroup
	if err := json.Unmarshal(bodyBytes, &alertGroup); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	/*for _, alert := range alertGroup.Alerts {
		fmt.Println("Request Body:", alert)
	}*/
	//  创建数据
	var Fields []controllers.Field
	for _, alert := range alertGroup.Alerts {
		for key, value := range alert.Annotations {
			Fields = append(Fields, controllers.Field{
				CustomField1: key,
				CustomField2: value,
			})
		}
		// 添加 StartsAt 作为单独的字段
		Fields = append(Fields, controllers.Field{
			CustomField1: "StartsAt",
			CustomField2: tools.ConvertUTCToChineseFormat(alert.StartsAt),
		})
	}

	//  调用生成 JSON 的方法
	result, err := controllers.GenerateJSON(Fields)
	if err != nil {
		fmt.Println("生成 JSON 时出错:", err)
		return
	}
	// 发送
	controllers.SendAlertMessage(chatId, result)
}
