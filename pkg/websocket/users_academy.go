package websocket

// UsersGetAcademy get users all academy
func UsersGetAcademy(s *Socket, a *Action) {
	if s.userID != 0 {
		academies, err := s.io.service.GetAllUsersAcademy(s.userID)
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
		degree, degreeOk := data["degree"]
		name, nameOk := data["name"].(string)
		institution, institutionOk := data["institution"].(string)
		year, yearOk := data["year"]
		if degreeOk && nameOk && institutionOk && yearOk {
			academy, err := s.io.service.AddUsersAcademy(s.userID, int(degree.(float64)), name, institution, int(year.(float64)))
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
