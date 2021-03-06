package services

import (
	"encoding/json"
	"net/http"

	"github.com/SunspotsInys/thedoor/logs"
	"github.com/SunspotsInys/thedoor/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetSysInfo(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logs.Error("Upgrade HTTP request failed")
		return
	}
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			logs.Error("Failed to read WS message")
			break
		}
		var u uint64
		err = json.Unmarshal(message, &u)
		if err != nil {
			logs.Errorf("Failed to Unmarshal data, err = %v", err)
			break
		}
		switch u {
		case 1:
			message, err = json.Marshal(utils.GetSysInfos())
			if err != nil {
				logs.Errorf("Failed to marshal data, err = %v", err)
				continue
			}
		case 2:
			message, err = json.Marshal(utils.GetNewestSysInfo())
			if err != nil {
				logs.Errorf("Failed to marshal data, err = %v", err)
				continue
			}
		default:
			logs.Infof("Unknown command, %d", u)
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			logs.Errorf("Failed to write data, err = %+v", err)
			continue
		}
	}
	ws.Close()
}
