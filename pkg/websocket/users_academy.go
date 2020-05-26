package websocket

// UsersGetAcademy get users all academy
func UsersGetAcademy(s *Socket, a *Action) {
	if s.userID != 0 {
		academies, err := s.io.service.GetUserAcademiesByUserID(s.userID)
		if err != nil {
			s.EmitServerError("Users get academy", err)
		} else {
			s.Emit("ADD_USERS_ACADEMY", academies)
		}
	} else {
		s.EmitError("Error al obtener los academies")
	}
}

// UsersAddAcademy add users academy
func UsersAddAcademy(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.userID != 0 {
		degree, degreeOk := data["degree"].(float64)
		name, nameOk := data["name"].(string)
		institution, institutionOk := data["institution"].(string)
		year, yearOk := data["year"].(float64)
		if degreeOk && nameOk && institutionOk && yearOk {
			academy, err := s.io.service.AddUsersAcademy(s.userID, int(degree), name, institution, int(year))
			if err != nil {
				s.EmitServerError("Error al guardar los datos", err)
			} else {
				s.Emit("ADD_USERS_ACADEMY", academy)
			}
		} else {
			s.EmitError("Error en los datos")
		}
	} else {
		s.EmitError("Error al guardar la información")
	}
}

// UsersUpdateAcademy update users academy by id
func UsersUpdateAcademy(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.userID != 0 {
		degree, degreeOk := data["degree"].(float64)
		name, nameOk := data["name"].(string)
		institution, institutionOk := data["institution"].(string)
		year, yearOk := data["year"].(float64)
		academyID, academyIDOk := data["id"].(float64)
		if degreeOk && nameOk && institutionOk && yearOk && academyIDOk {
			academy, err := s.io.service.UpdateUsersAcademy(int(academyID), s.userID, int(degree), name, institution, int(year))
			if err != nil {
				s.EmitServerError("Error al actualizar tus datos", err)
			} else {
				s.Emit("ADD_USERS_ACADEMY", academy)
			}
		} else {
			s.EmitError("Error con los datos")
		}
	} else {
		s.EmitError("Error al actualizar tus datos")
	}
}

// UsersDeleteAcademy delete users academy by id
func UsersDeleteAcademy(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.userID != 0 {
		academyID, academyIDOk := data["id"].(float64)
		if academyIDOk {
			err := s.io.service.DeleteUsersAcademy(int(academyID), s.userID)
			if err != nil {
				s.EmitServerError("Error al eliminar tus datos", err)
			} else {
				s.Emit("DELETE_USERS_ACADEMY", academyID)
			}
		} else {
			s.EmitError("Error al eliminar tus datos")
		}
	} else {
		s.EmitError("Error al eliminar tus datos")
	}
}
