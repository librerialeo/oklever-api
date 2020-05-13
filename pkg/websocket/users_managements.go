package websocket

// GetUserManagement Get user management
func GetUserManagement(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok {
		managementID, ok := data["id"]
		if ok {
			management, err := s.io.service.GetUserManagement(int32(managementID.(float64)))
			if err != nil {
				s.EmitServerError("GetUserManagement", err)
			} else {
				s.Emit(a.Type, management)
			}
		}
	}
}

// GetUserManagements Get user management
func GetUserManagements(s *Socket, a *Action) {
	if s.userID != 0 {
		managements, err := s.io.service.GetUserManagements(s.userID)
		if err != nil {
			s.EmitServerError("GetUserManagements", err)
		} else {
			s.Emit(a.Type, *managements)
		}
	}
}

// AddUserManagement Add user management
func AddUserManagement(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if ok && s.userID != 0 {
		job, jOk := data["job"]
		institution, iOk := data["institution"]
		months, mOk := data["months"]
		if jOk && iOk && mOk {
			management, err := s.io.service.AddUserManagement(s.userID, job.(string), institution.(string), int32(months.(float64)))
			if err != nil {
				s.EmitServerError("AddUserManagement", err)
			} else {
				s.Emit(a.Type, management)
			}
		}
	}
}

// UpdateUserManagement Update user management
func UpdateUserManagement(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if !ok || s.userID == 0 {
		return
	}
	managementID, ok := data["id"]
	job, jOk := data["job"]
	institution, rOk := data["institution"]
	months, mOk := data["months"]
	if !ok || !rOk || !mOk || jOk {
		return
	}
	management, err := s.io.service.GetUserManagement(int32(managementID.(float64)))
	if err != nil || management == nil {
		s.EmitServerError("UpdateUserManagement: management not found", err)
		return
	}
	if management.UserID.Int != s.userID {
		return
	}
	management, err = s.io.service.UpdateUserManagement(int32(managementID.(float64)), job.(string), institution.(string), int32(months.(float64)))
	if err != nil {
		s.EmitServerError("UpdateUserManagement: update user error", err)
		return
	}
	s.Emit(a.Type, management)
}

// DeleteUserManagement Delete user management
func DeleteUserManagement(s *Socket, a *Action) {
	data, ok := a.Data.(map[string]interface{})
	if !ok {
		return
	}
	managementID, ok := data["id"]
	if !ok {
		return
	}
	management, err := s.io.service.GetUserManagement(int32(managementID.(float64)))
	if err != nil || management == nil {
		s.EmitServerError("DeleteUserManagement: management not found", err)
		return
	}
	if management.UserID.Int != s.userID {
		return
	}
	err = s.io.service.DeleteUserManagement(int32(managementID.(float64)))
	if err != nil {
		s.EmitServerError("DeleteUserManagement: delete error", err)
		return
	}
	s.Emit(a.Type, managementID)
}
