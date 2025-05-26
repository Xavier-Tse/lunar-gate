package ws_api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WsApi struct {
}

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Request struct {
}

type Response struct {
	Type    string `json:"type"` // logs
	Content string `json:"content"`
}

var WsMap = map[string]*websocket.Conn{}

func SendLogMsg(content string) {
	for _, conn := range WsMap {
		conn.WriteJSON(Response{
			Type:    "logs",
			Content: content,
		})
	}
}

func (WsApi) Ws(c *gin.Context) {
	conn, err := UP.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Errorf("ws升级失败 %s", err)
		return
	}

	addr := conn.RemoteAddr().String()
	WsMap[addr] = conn
	logrus.Infof("ws %s 进入", addr)

	for {
		// 消息类型，消息，错误
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}

	}
	defer conn.Close()
	delete(WsMap, addr)
	logrus.Infof("ws %s 退出", addr)
}
