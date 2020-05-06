package websocket

import "fmt"

// UserRegister register new user
func UserRegister(socket *Socket, action *Action) {
}

// Logout desubscribe socket from rooms and send logout action
func Logout(s *Socket, a *Action) {
	for c, d := range s.rooms {
		fmt.Println(c, d)
	}
	s.SetToken("")
	s.Emit("LOGOUT", "")
}
