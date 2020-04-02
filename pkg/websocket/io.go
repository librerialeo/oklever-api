package websocket

import (
	"fmt"

	"github.com/librerialeo/oklever-api/pkg/service"
)

// NewIO creates a new IO instance
func NewIO(s *service.Service) *IO {
	return &IO{
		sockets: make(map[*Socket]bool),
		rooms:   make(map[string]*Room),
		service: s,
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
func (io *IO) Emit(message []byte) {
	fmt.Println(message, io.sockets)
	for s := range io.sockets {
		s.Emit(message)
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
	io.rooms[name] = &Room{sockets: make(map[*Socket]bool)}
	return io.rooms[name]
}

// Room is a socket.io like structure
type Room struct {
	sockets map[*Socket]bool
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
func (r *Room) Emit(message []byte) {
	for s := range r.sockets {
		s.Emit(message)
	}
}
