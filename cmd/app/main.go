package main

import (
	"Robot_Alertmanager_Fs/api"
	robot_Clicent "Robot_Alertmanager_Fs/config"
	"net/http"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/core/httpserverext"
	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
)

func main() {
	// 注册卡片回调
	cardHandler := larkcard.NewCardActionHandler(robot_Clicent.ConfigInstance.VerificationToken, robot_Clicent.ConfigInstance.EncryptKey, api.DoInteractiveCard)

	http.HandleFunc("/send", api.SendAlertHandler)
	http.HandleFunc("/card", httpserverext.NewCardActionHandlerFunc(cardHandler,
		larkevent.WithLogLevel(larkcore.LogLevelDebug)))

	err := http.ListenAndServe(":7778", nil)
	if err != nil {
		panic(err)
	}
}
