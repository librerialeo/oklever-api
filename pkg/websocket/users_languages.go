package websocket

// UsersAddLanguage add language
func UsersAddLanguage(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.userID != 0 {
		id, idOk := data["id"]
		if idOk {
			err := s.io.service.AddUsersLanguages(int(s.userID), int(id.(float64)))
			if err != nil {
				s.EmitError("Error al registrar el idioma")
			} else {
				s.Emit("ADD_USERS_LANGUAGE", [1]int{int(id.(float64))})
			}
		} else {
			s.EmitError("Error al registrar el idioma")
		}
	} else {
		s.EmitError("Error al registrar el idioma")
	}
}

// UsersDeleteLanguage delete teacher language
func UsersDeleteLanguage(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.userID != 0 {
		id, idOk := data["id"]
		if idOk {
			err := s.io.service.DeleteUsersLanguages(int(s.userID), int(id.(float64)))
			if err != nil {
				s.EmitError("No se puedo eliminar el idioma")
			} else {
				s.Emit("DELETE_USERS_LANGUAGE", id)
			}
		} else {
			s.EmitError("No se puedo eliminar el idioma")
		}
	} else {
		s.EmitError("No se puedo eliminar el idioma")
	}
}
