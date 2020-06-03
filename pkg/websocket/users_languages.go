package websocket

//UsersGetLanguages get all languages
func UsersGetLanguages(s *Socket, a *Action) {
	if s.user.ID != 0 {
		ids, err := s.io.service.GetAllUsersLanguages(s.user.ID)
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
	if ok && s.user.ID != 0 {
		id, idOk := data["id"].(float64)
		if idOk {
			err := s.io.service.AddUsersLanguages(s.user.ID, int(id))
			if err != nil {
				s.EmitServerError("Error al registrar el idioma", err)
			} else {
				s.Emit("ADD_USERS_LANGUAGE", [1]int{int(id)})
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
	if ok && s.user.ID != 0 {
		id, idOk := data["id"].(float64)
		if idOk {
			err := s.io.service.DeleteUsersLanguages(s.user.ID, int(id))
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
