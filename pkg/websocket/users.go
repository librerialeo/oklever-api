package websocket

import "github.com/savsgio/atreugo"

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

// GetUserBiography send user biography
func GetUserBiography(s *Socket, a *Action) {
	biography, err := s.io.service.GetUserBiography(s.userID)
	if err != nil {
		s.EmitServerError("GetUserBiography", err)
	} else {
		s.Emit(a.Type, atreugo.JSON{
			"biography": biography,
		})
	}
}

// SetUserBiography send user biography
func SetUserBiography(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		biography, ok := data["biography"]
		if ok {
			DBbiography, err := s.io.service.SetUserBiography(s.userID, biography.(string))
			if err != nil {
				s.EmitServerError("GetUserBiography", err)
			} else {
				s.Emit(a.Type, atreugo.JSON{
					"biography": DBbiography,
				})
			}
		}
	}
}
