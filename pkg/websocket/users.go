package websocket

// UserRegister register new user
func UserRegister(socket *Socket, action *Action) {
}

// Logout desubscribe socket from rooms and send logout action
func Logout(s *Socket, a *Action) {
	for _, r := range s.rooms {
		s.LeaveRoom(r)
	}
	s.SetToken("")
	s.Emit("LOGOUT")
}
