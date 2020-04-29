package websocket

import (
	"github.com/librerialeo/oklever-api/pkg/service"
)

// NewIO creates a new IO instance
func NewIO(s *service.Service) *IO {
	return &IO{
		sockets: make(map[*Socket]bool),
		rooms:   make(map[string]*Room),
		service: s,
		actions: make(map[string][]ActionHandler),
	}
}

// IOSender wraps sender interface
// type IOSender interface {
// 	Emit([]byte)
// }

// IO is a socket.io like structure
type IO struct {
	sockets map[*Socket]bool
	rooms   map[string]*Room
	service *service.Service
	actions map[string][]ActionHandler
}

// ActionHandler wraps action function handler
type ActionHandler struct {
	Handler     func(*Socket, *Action)
	Credentials []string
}

// ActionsHandlers is and array

// AddActionHandler adds a new action handler to map
func (io *IO) AddActionHandler(action string, handler func(*Socket, *Action), credentials []string) {
	io.actions[action] = []ActionHandler{{Handler: handler, Credentials: credentials}}
}

func (io *IO) addSocket(s *Socket) {
	io.sockets[s] = true
}

func (io *IO) removeSocket(s *Socket) {
	if _, ok := io.sockets[s]; ok {
		delete(io.sockets, s)
		close(s.send)
	}
}

// Emit send message to all sockets
func (io *IO) Emit(actionType string, data interface{}) {
	for s := range io.sockets {
		s.Emit(actionType, data)
	}
}

// Room create/get room name and return it
func (io *IO) Room(name string) *Room {
	if _, ok := io.rooms[name]; ok {
		return io.rooms[name]
	}
	return io.NewRoom(name)
}

// NewRoom create a new room for io
func (io *IO) NewRoom(name string) *Room {
	io.rooms[name] = &Room{
		sockets: make(map[*Socket]bool),
		io:      io,
	}
	return io.rooms[name]
}

// Room is a socket.io like structure
type Room struct {
	sockets map[*Socket]bool
	io      *IO
}

func (r *Room) addSocket(s *Socket) {
	r.sockets[s] = true
}

func (r *Room) removeSocket(s *Socket) {
	if _, ok := r.sockets[s]; ok {
		delete(r.sockets, s)
	}
}

// Emit send message to room's sockets
func (r *Room) Emit(actionType string, data interface{}) {
	for s := range r.sockets {
		s.Emit(actionType, data)
	}
}

// Broadcast send message to all sockets but room's sockets
func (r *Room) Broadcast(actionType string, data interface{}) {
	for s := range r.io.sockets {
		if _, ok := r.sockets[s]; !ok {
			s.Emit(actionType, data)
		}
	}
}
