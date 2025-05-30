package handlers

import (
	"fmt"
	"net/http"
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type connectionPool struct {
	Conns map[string]*websocket.Conn
}

var lock = &sync.Mutex{}
var pool *connectionPool

func addConnection(conn *websocket.Conn) string {
	id := uuid.New().String()
	if pool == nil {
		lock.Lock()
		defer lock.Unlock()
		if pool == nil {
			pool = &connectionPool{
				Conns: make(map[string]*websocket.Conn),
			}
		}
	}
	fmt.Println(fmt.Sprintf("[WS] Connection Added: %s", id))
	pool.Conns[id] = conn
	return id
}

func removeConnection(id string) {
	delete(pool.Conns, id)
	fmt.Println(fmt.Sprintf("[WS] Connection Removed: %s", id))
}

func sendPoolJson(jsonable interface{}, excludes ...string) error {
	for id, conn := range pool.Conns {
		if slices.Index(excludes, id) != -1 {
			continue
		}
		err := conn.WriteJSON(jsonable)
		if err != nil {
			return err
		}
	}
	return nil
}

type WsMessage struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func UpdateCounter(m map[string]interface{}, clientId string) error {
	msg := WsMessage{
		Message: "client-update-counter",
		Data:    m,
	}
	err := sendPoolJson(msg, clientId)
	if err != nil {
		return err
	}
    return nil
}

func WsCounter(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	WsRoutes := map[string]func(map[string]interface{}, string) error{
		"update-counter": UpdateCounter,
	}

	go func(conn *websocket.Conn) {
		id := addConnection(conn)
		defer removeConnection(id)
		defer conn.Close()

		for {
			msg := &WsMessage{}
			err := conn.ReadJSON(msg)
			if err != nil {
				fmt.Println(err)
				break
			}

			handler, ok := WsRoutes[msg.Message]
			if !ok {
				fmt.Println(fmt.Sprintf("[WS] Recieved invalid message: %s", msg.Message))
				continue
			}
			err = handler(msg.Data, id)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}(conn)
}
