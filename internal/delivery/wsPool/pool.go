package wsPool

import (
	"depeche/pkg/apperror"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type ConnectionPool struct {
	conns    map[string][]*websocket.Conn
	mx       *sync.RWMutex
	upgrader websocket.Upgrader
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		conns: make(map[string][]*websocket.Conn),
		mx:    &sync.RWMutex{},
		// TODO configure upgrader
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (cp *ConnectionPool) Connect(ctx *gin.Context) {

	emailRaw, ok := ctx.Get("email")
	email, ok := emailRaw.(string)
	if !ok {
		_ = ctx.Error(apperror.NoAuth)
		return
	}

	conn, err := cp.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	cp.NewConnection(email, conn)

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		err = cp.RemoveConn(email, conn)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	}(conn)

	for {
		fmt.Println("CONNECTED")
		// TODO add read, set offline etc handling
		_, _, err = conn.ReadMessage()
		if err != nil {
			err, ok := err.(*websocket.CloseError)
			if !ok {
				_ = ctx.Error(err)
				return
			}
			if err.Code != websocket.CloseNormalClosure {
				_ = ctx.Error(err)
				return
			}
			break
		}
	}

}

func (cp *ConnectionPool) NewConnection(email string, conn *websocket.Conn) {
	cp.mx.Lock()
	cp.conns[email] = append(cp.conns[email], conn)
	cp.mx.Unlock()
}

func (cp *ConnectionPool) SendMsg(email string, msg []byte) error {
	cp.mx.RLock()
	defer cp.mx.RUnlock()
	for _, conn := range cp.conns[email] {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cp *ConnectionPool) RemoveConn(email string, conn *websocket.Conn) error {
	cp.mx.Lock()
	defer cp.mx.Unlock()
	for i, stored := range cp.conns[email] {
		if conn == stored {
			cp.conns[email][i] = cp.conns[email][len(cp.conns[email])-1]
			cp.conns[email] = cp.conns[email][:len(cp.conns[email])-1]
			return nil
		}
	}
	return errors.New("ws: remove connection error: connection not found")
}
