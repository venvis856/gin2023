package gorillaWebsocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type GorillaWebsocketCtroller struct{}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (a *GorillaWebsocketCtroller) Websocket(c *gin.Context) {
	//升级get请求为webSocket协议
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	fmt.Println("有用户链接上来了======")

	// 创建读写channel
	readCh := make(chan []byte)
	writeCh := make(chan []byte)
	closeCh := make(chan bool)

	// 启动读写协程
	go read(conn, readCh, closeCh)
	go write(conn, writeCh, closeCh)

	for {
		// 监听读写channel
		select {
		case <-closeCh:
			return
		case msg := <-readCh:
			fmt.Println("msg received: ", string(msg))
			writeCh <- []byte("消息收到了")
		}
	}

}

func read(conn *websocket.Conn, readCh chan []byte, closeCh chan bool) {
	defer closeConn(conn, closeCh)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// 控制读取消息的速率
			time.Sleep(500 * time.Millisecond)
		default:
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("read err==", err)
				return
			}
			readCh <- msg
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func write(conn *websocket.Conn, writeCh chan []byte, closeCh chan bool) {
	defer closeConn(conn, closeCh)
	for {
		select {
		case msg := <-writeCh:
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Println("sent err===", err)
				return
			}
			time.Sleep(100 * time.Millisecond)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func closeConn(conn *websocket.Conn, closeCh chan bool) {
	err := conn.Close()
	if err != nil {
		fmt.Println("关闭连接失败", err)
		closeCh <- true
		return
	}
	fmt.Println("关闭了连接")
	closeCh <- true
}
