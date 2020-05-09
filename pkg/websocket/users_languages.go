package websocket

//UsersGetLanguages get all languages
func UsersGetLanguages(s *Socket, a *Action) {
	if s.userID != 0 {
		ids, err := s.io.service.GetAllUsersLanguages(s.userID)
		if err != nil {
			s.EmitServerError("Users get languages", err)
		} else {
			s.Emit("ADD_USERS_LANGUAGE", ids)
		}
	} else {
		s.EmitError("Error al obtener los idiomas")
	}
}

// UsersAddLanguage add language
func UsersAddLanguage(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.userID != 0 {
		id, idOk := data["id"]
		if idOk {
			err := s.io.service.AddUsersLanguages(s.userID, int(id.(float64)))
			if err != nil {
				s.EmitServerError("Error al registrar el idioma", err)
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
			err := s.io.service.DeleteUsersLanguages(s.userID, int(id.(float64)))
			if err != nil {
				s.EmitServerError("No se puedo eliminar el idioma", err)
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
