package websocket

import (
	"fmt"
	"log"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/fasthttp/websocket"
	"github.com/savsgio/atreugo"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Action is a redux like action
type Action struct {
	Type  string      `json:"type"`
	Token string      `json:"token"`
	Data  interface{} `json:"data"`
}

// Socket is a client websocket connection instance
type Socket struct {
	ID    uint16
	io    *IO
	conn  *websocket.Conn
	rooms []string
	data  map[string]interface{}
	send  chan []byte
}

func (s *Socket) unregister() {
	s.io.removeSocket(s)
	s.conn.Close()
}

func (s *Socket) register() {
	s.io.addSocket(s)
}

// Emit send message to socket
func (s *Socket) Emit(actionType string, data interface{}) {
	message, err := jsoniter.Marshal(Action{
		Type: actionType,
		Data: data,
	})
	if err != nil {
		log.Printf("error: %v", err)
	} else {
		s.send <- message
	}
}

// Broadcast send message to all io sockets but socket
func (s *Socket) Broadcast(actionType string, data interface{}) {
	for socket := range s.io.sockets {
		if socket.ID != s.ID {
			socket.Emit(actionType, data)
		}
	}
}

// JoinRoom join socket to room
func (s *Socket) JoinRoom(room string) {
	s.io.Room(room).addSocket(s)
}

// LeaveRoom leave socket from room
func (s *Socket) LeaveRoom(room string) {
	s.io.Room(room).removeSocket(s)
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (s *Socket) readPump() {
	defer s.unregister()
	s.conn.SetReadLimit(maxMessageSize)
	s.conn.SetReadDeadline(time.Now().Add(pongWait))
	s.conn.SetPongHandler(func(string) error { s.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := s.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v\n", err)
			}
			break
		}
		var action Action
		err = jsoniter.Unmarshal(message, &action)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("socketID: %v, action: %s\n", s.ID, action.Type)

		// actions

		s.io.Emit(action.Type, action.Data)
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (s *Socket) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		s.conn.Close()
	}()
	for {
		select {
		case message, ok := <-s.send:
			s.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				s.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := s.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(s.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-s.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			s.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := s.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// SocketInit handles websocket requests from the peer.
func SocketInit(ctx *atreugo.RequestCtx, io *IO) {
	err := upgrader.Upgrade(ctx.RequestCtx, func(conn *websocket.Conn) {
		var id uint16
		id = 0
		for c := range io.sockets {
			if c.ID >= id {
				id = c.ID + 1
			}
		}
		socket := &Socket{ID: id, io: io, conn: conn, send: make(chan []byte, 256), data: make(map[string]interface{})}
		socket.register()

		go socket.writePump()
		go socket.readPump()
	})

	if err != nil {
		log.Println(err)
	}
}
